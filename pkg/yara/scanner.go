//go:build yara
// +build yara

/*
 * Copyright 2019-2020 by Nedim Sabic Sabic
 * https://www.fibratus.io
 * All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package yara

import (
	"encoding/json"
	"expvar"
	"fmt"
	"github.com/google/uuid"
	"github.com/rabbitstack/fibratus/pkg/alertsender"
	"github.com/rabbitstack/fibratus/pkg/event/params"
	libntfs "github.com/rabbitstack/fibratus/pkg/fs/ntfs"
	"github.com/rabbitstack/fibratus/pkg/sys"
	"github.com/rabbitstack/fibratus/pkg/util/signature"
	"github.com/rabbitstack/fibratus/pkg/util/va"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
	"os"
	"path/filepath"
	"strings"

	"github.com/hillu/go-yara/v4"
	"github.com/rabbitstack/fibratus/pkg/event"
	"github.com/rabbitstack/fibratus/pkg/ps"
	"github.com/rabbitstack/fibratus/pkg/util/multierror"
	"github.com/rabbitstack/fibratus/pkg/yara/config"
	ytypes "github.com/rabbitstack/fibratus/pkg/yara/types"
	log "github.com/sirupsen/logrus"
)

var (
	// ruleMatches computes all the rule matches
	ruleMatches = expvar.NewInt("yara.rule.matches")
	// rulesInCompiler keeps the counter of the number of rules in the compiler
	rulesInCompiler = expvar.NewInt("yara.rules.in.compiler")
	// totalScans computes the number of process/file scans
	totalScans = expvar.NewInt("yara.total.scans")

	procScans     = expvar.NewInt("yara.proc.spawned.scans")
	moduleScans   = expvar.NewInt("yara.module.loaded.scans")
	fileScans     = expvar.NewInt("yara.file.created.scans")
	streamScans   = expvar.NewInt("yara.ads.created.scans")
	allocScans    = expvar.NewInt("yara.alloc.scans")
	mmapScans     = expvar.NewInt("yara.mmap.scans")
	registryScans = expvar.NewInt("yara.registry.value.set.scans")
)

type scanner struct {
	c      *yara.Compiler
	rules  *yara.Rules
	config config.Config

	psnap ps.Snapshotter

	rwxs  map[uint32]va.Address // contains scanned and matched RWX process allocations
	mmaps map[uint32]va.Address // contains scanned and matched suspicious memory mappings
}

// NewScanner creates a new YARA scanner.
func NewScanner(psnap ps.Snapshotter, config config.Config) (Scanner, error) {
	c, err := yara.NewCompiler()
	if err != nil {
		return nil, fmt.Errorf("unable to create yara compiler: %v", err)
	}
	// add yara rules from file system paths by walking the dirs recursively
	for _, dir := range config.Rule.Paths {
		f, err := os.Stat(dir.Path)
		if err != nil {
			log.Warnf("cannot access %q rule path: %v", dir.Path, err)
			continue
		}
		if !f.IsDir() {
			continue
		}
		err = filepath.Walk(dir.Path, func(path string, fi os.FileInfo, err error) error {
			if filepath.Ext(path) != ".yar" {
				return nil
			}
			f, err := os.Open(path)
			if err != nil {
				log.Warnf("cannot open the rule %q: %v", path, err)
				return nil
			}
			err = c.AddFile(f, dir.Namespace)
			_ = f.Close()
			if err != nil {
				log.Warnf("couldn't add %s rule: %v", fi.Name(), err)
				return nil
			}
			rulesInCompiler.Add(1)
			log.Infof("loading yara rule(s) from %s", path)

			return nil
		})
		if err != nil {
			log.Warnf("couldn't walk %s path: %v", dir.Path, err)
		}
	}

	// add yara rules from config strings
	for _, s := range config.Rule.Strings {
		err := c.AddString(s.String, s.Namespace)
		if err != nil {
			log.Warnf("couldn't add %s rule string: %v", s.String, err)
			continue
		}
		rulesInCompiler.Add(1)
	}

	if len(c.Errors) > 0 {
		return nil, parseCompilerErrors(c.Errors)
	}

	rules, err := c.GetRules()
	if err != nil {
		return nil, fmt.Errorf("couldn't compile yara rules: %v", err)
	}

	return &scanner{
		c:      c,
		rules:  rules,
		config: config,
		psnap:  psnap,
		rwxs:   make(map[uint32]va.Address),
		mmaps:  make(map[uint32]va.Address),
	}, nil
}

// newInternalScanner creates a new instance of the go-yara scanner.
func (s scanner) newInternalScanner() (*yara.Scanner, error) {
	sn, err := yara.NewScanner(s.rules)
	if err != nil {
		return nil, fmt.Errorf("fail to create yara scanner: %v", err)
	}
	// set scan flags
	var flags yara.ScanFlags
	if s.config.FastScanMode {
		flags |= yara.ScanFlagsFastMode
	}
	sn.SetFlags(flags)
	sn.SetTimeout(s.config.ScanTimeout)
	return sn, nil
}

func parseCompilerErrors(errors []yara.CompilerMessage) error {
	errs := make([]error, len(errors))
	for i, err := range errors {
		errs[i] = fmt.Errorf("%s, filename: %s line: %d", err.Text, err.Filename, err.Line)
	}
	return multierror.Wrap(errs...)
}

func (s scanner) CanEnqueue() bool { return false }

func (s *scanner) ProcessEvent(evt *event.Event) (bool, error) {
	if evt.IsTerminateProcess() {
		// cleanup
		pid := evt.Params.MustGetPid()
		delete(s.rwxs, pid)
		delete(s.mmaps, pid)
	}
	return s.Scan(evt)
}

func (s scanner) Scan(e *event.Event) (bool, error) {
	var matches yara.MatchRules
	var isScanned bool
	var err error

	switch e.Type {
	case event.CreateProcess:
		// scan the created child process
		pid := e.Params.MustGetPid()
		log.Debugf("scanning child process. pid: %d, exe: %s", pid, e.GetParamAsString(params.Exe))
		matches, err = s.scan(pid)
		procScans.Add(1)
		isScanned = true
	case event.LoadImage:
		// scan the process loading unsigned/untrusted module
		// or loading the module from unbacked memory region
		pid := e.PID
		addr := e.Params.MustGetUint64(params.ImageBase)
		typ := e.Params.MustGetUint32(params.ImageSignatureType)
		if typ != signature.None {
			return false, nil
		}
		filename := e.GetParamAsString(params.ImagePath)
		if s.config.ShouldSkipFile(filename) {
			return false, nil
		}

		// get module signature
		sign := signature.GetSignatures().GetSignature(addr)
		if sign == nil {
			sign = &signature.Signature{Filename: filename}
			sign.Type, sign.Level, err = sign.Check()
			if sign.IsSigned() {
				sign.Verify()
			}
		}

		if !sign.IsSigned() || !sign.IsTrusted() || (!e.Callstack.IsEmpty() && e.Callstack.ContainsUnbacked()) {
			log.Debugf("scanning suspicious module loading. pid: %d, module: %s", pid, filename)
			matches, err = s.scan(pid)
			moduleScans.Add(1)
			isScanned = true
		}
	case event.CreateFile:
		if s.config.SkipFiles {
			return false, nil
		}
		if e.IsOpenDisposition() {
			return false, nil
		}

		filename := e.GetParamAsString(params.FilePath)
		if s.config.ShouldSkipFile(filename) || (e.PS != nil && s.config.ShouldSkipProcess(e.PS.Exe)) {
			return false, nil
		}

		// scan dropped PE files
		isDLL := strings.ToLower(filepath.Ext(filename)) == ".dll" || e.Params.TryGetBool(params.FileIsDLL)
		isDriver := strings.ToLower(filepath.Ext(filename)) == ".sys" || e.Params.TryGetBool(params.FileIsDriver)
		isExe := strings.ToLower(filepath.Ext(filename)) == ".exe" || e.Params.TryGetBool(params.FileIsExecutable)

		if isExe || isDLL || isDriver {
			log.Debugf("scanning PE file %s. pid: %d", filename, e.PID)
			matches, err = s.scan(filename)
			fileScans.Add(1)
			isScanned = true
			break
		}

		// scan dropped ADS (Alternate Data Stream)
		n := strings.LastIndex(filename, ":")
		ads := n > 2 && n+1 <= len(filename)
		if !ads {
			return false, nil
		}
		// read ADS data
		ntfs := libntfs.NewFS()
		data, n, err := ntfs.Read(filename, 0, 1024*1024)
		defer ntfs.Close()
		if err != nil {
			return false, nil
		}
		if n > 0 {
			data = data[:n]
			log.Debugf("scanning ADS %s. pid: %d", filename, e.PID)
			matches, err = s.scan(data)
			streamScans.Add(1)
			isScanned = true
		}
	case event.VirtualAlloc:
		if s.config.SkipAllocs {
			return false, nil
		}
		// scan process allocating RWX memory region
		pid := e.Params.MustGetPid()
		addr := e.Params.TryGetAddress(params.MemBaseAddress)
		if e.PID != 4 && e.Params.TryGetUint32(params.MemProtect) == windows.PAGE_EXECUTE_READWRITE && !s.isRwxMatched(pid) {
			log.Debugf("scanning RWX allocation. pid: %d, exe: %s, addr: %s", pid, e.GetParamAsString(params.Exe),
				e.GetParamAsString(params.MemBaseAddress))
			matches, err = s.scan(pid)
			allocScans.Add(1)
			isScanned = true
			if len(matches) > 0 {
				s.rwxs[pid] = addr
			}
		}
	case event.MapViewFile:
		if s.config.SkipMmaps {
			return false, nil
		}
		// scan process mapping a suspicious RX/RWX section view
		pid := e.Params.MustGetPid()
		prot := e.Params.MustGetUint32(params.MemProtect)
		size := e.Params.MustGetUint64(params.FileViewSize)
		if e.PID != 4 && size >= 4096 && ((prot&sys.SectionRX) != 0 && (prot&sys.SectionRWX) != 0) && !s.isMmapMatched(pid) {
			filename := e.GetParamAsString(params.FilePath)
			// skip mappings of signed images
			addr := e.Params.MustGetUint64(params.FileViewBase)
			sign := signature.GetSignatures().GetSignature(addr)
			if sign != nil && sign.IsSigned() && sign.IsTrusted() {
				return false, nil
			}
			// data/image file was mapped?
			if filename != "" {
				if s.config.ShouldSkipFile(filename) {
					return false, nil
				}
				log.Debugf("scanning %s section view mapping. filename: %s pid: %d, addr: %s", e.GetParamAsString(params.MemProtect),
					filename, pid, e.GetParamAsString(params.FileViewBase))
				matches, err = s.scan(filename)
			} else {
				// otherwise, scan the process
				log.Debugf("scanning %s section view mapping. pid: %d, addr: %s", e.GetParamAsString(params.MemProtect), pid,
					e.GetParamAsString(params.FileViewBase))
				matches, err = s.scan(pid)
			}
			if len(matches) > 0 {
				s.mmaps[pid] = va.Address(addr)
			}
			mmapScans.Add(1)
			isScanned = true
		}
	case event.RegSetValue:
		if s.config.SkipRegistry {
			return false, nil
		}
		if e.PS != nil && s.config.ShouldSkipProcess(e.PS.Exe) {
			return false, nil
		}
		// scan registry binary values
		if typ := e.Params.TryGetUint32(params.RegValueType); typ != registry.BINARY {
			return false, nil
		}
		v, err := e.Params.Get(params.RegData)
		if err != nil {
			// value not attached to the event
			return false, nil
		}
		if b, ok := v.Value.([]byte); ok && len(b) > 0 {
			log.Debugf("scanning registry binary value %s. pid: %d", e.GetParamAsString(params.RegPath), e.PID)
			matches, err = s.scan(b)
			registryScans.Add(1)
			isScanned = true
		}
	}

	if err != nil {
		return false, err
	}
	if len(matches) == 0 || !isScanned {
		return false, nil
	}

	totalScans.Add(1)
	ruleMatches.Add(int64(len(matches)))

	return len(matches) > 0, s.emit(matches, e)
}

func (s scanner) scan(target any) (yara.MatchRules, error) {
	var matches yara.MatchRules
	sn, err := s.newInternalScanner()
	if err != nil {
		return nil, err
	}

	switch n := target.(type) {
	case uint32: // pid
		// skip the scan for this process?
		ok, proc := s.psnap.Find(n)
		if ok && s.config.ShouldSkipProcess(proc.Exe) {
			return matches, nil
		}
		err = sn.SetCallback(&matches).ScanProc(int(n))
	case string: // file
		err = sn.SetCallback(&matches).ScanFile(n)
	case []byte: // mem
		err = sn.SetCallback(&matches).ScanMem(n)
	}

	if err != nil {
		return nil, err
	}

	return matches, nil
}

func (s scanner) emit(matches yara.MatchRules, e *event.Event) error {
	senders := alertsender.FindAll()
	if len(senders) == 0 {
		return fmt.Errorf("no alertsenders registered. Alert won't be sent")
	}

	ymatches := make([]ytypes.MatchRule, 0)

	for _, match := range matches {
		// encode rule matches as JSON and append to event metadata
		m := ytypes.MatchRule{
			Rule:      match.Rule,
			Namespace: match.Namespace,
			Tags:      match.Tags,
			Metas:     make([]ytypes.Meta, 0),
			Strings:   make([]ytypes.MatchString, 0),
		}
		for _, meta := range match.Metas {
			m.Metas = append(m.Metas, ytypes.Meta{Value: meta.Value, Identifier: meta.Identifier})
		}
		for _, s := range match.Strings {
			m.Strings = append(m.Strings, ytypes.MatchString{Name: s.Name, Base: s.Base, Data: s.Data, Offset: s.Offset})
		}
		ymatches = append(ymatches, m)

		b, err := json.Marshal(ymatches)
		if err != nil {
			return err
		}
		e.AddMeta(event.YaraMatchesKey, string(b))

		// render alert title and text
		title := s.config.AlertTitle(e)
		text, err := s.config.AlertText(e, m)
		if err != nil {
			return err
		}

		// send alert via all registered alert senders
		for _, sender := range senders {
			log.Infof("sending alert: [%s]. Text: %s Event: %s", title, text, e.String())

			alert := alertsender.NewAlert(
				title,
				text,
				m.Tags,
				m.SeverityFromScore(),
			)

			id := m.ID()
			// generate id if it doesn't exist in meta fields
			if id == "" {
				id = uuid.New().String()
			}
			alert.ID = id
			alert.Events = []*event.Event{e}
			alert.Labels = m.Labels()
			alert.Description = m.Description()

			err := sender.Send(alert)
			if err != nil {
				return fmt.Errorf("unable to emit YARA alert via [%s] sender: %v", sender.Type(), err)
			}
		}
	}

	return nil
}

func (s scanner) Close() {
	if s.c != nil {
		s.c.Destroy()
	}
}

// isRwxMatched returns true if the process already triggered RWX allocation rule match.
func (s *scanner) isRwxMatched(pid uint32) (ok bool) {
	_, ok = s.rwxs[pid]
	return ok
}

// isMmapMatched returns true if the process already triggered suspicious mmap rule match.
func (s *scanner) isMmapMatched(pid uint32) (ok bool) {
	_, ok = s.mmaps[pid]
	return ok
}
