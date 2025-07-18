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

package processors

import (
	"github.com/rabbitstack/fibratus/pkg/event"
	"github.com/rabbitstack/fibratus/pkg/event/params"
	"github.com/rabbitstack/fibratus/pkg/ps"
	"github.com/rabbitstack/fibratus/pkg/util/cmdline"
	"github.com/rabbitstack/fibratus/pkg/util/multierror"
	"github.com/rabbitstack/fibratus/pkg/util/va"
	"golang.org/x/sys/windows"
	"time"
)

type psProcessor struct {
	psnap        ps.Snapshotter
	regionProber *va.RegionProber
}

// newPsProcessor creates a new event processor for process/thread events.
func newPsProcessor(psnap ps.Snapshotter, regionProber *va.RegionProber) Processor {
	return &psProcessor{psnap: psnap, regionProber: regionProber}
}

func (p psProcessor) ProcessEvent(e *event.Event) (*event.Event, bool, error) {
	switch e.Type {
	case event.CreateProcess, event.CreateProcessInternal, event.TerminateProcess, event.ProcessRundown, event.ProcessRundownInternal:
		evt, err := p.processEvent(e)
		if evt.IsTerminateProcess() {
			p.regionProber.Remove(evt.Params.MustGetPid())
			return evt, false, multierror.Wrap(err, p.psnap.Remove(evt))
		}

		return evt, false, multierror.Wrap(err, p.psnap.Write(evt))
	case event.CreateThread, event.TerminateThread, event.ThreadRundown:
		pid, err := e.Params.GetPid()
		if err != nil {
			return e, false, err
		}
		proc := p.psnap.FindAndPut(pid)
		if proc != nil {
			e.AppendParam(params.Exe, params.UnicodeString, proc.Exe)
		}
		if !e.IsTerminateThread() {
			return e, false, p.psnap.AddThread(e)
		}
		tid, err := e.Params.GetTid()
		if err != nil {
			return e, false, err
		}

		return e, false, p.psnap.RemoveThread(pid, tid)
	case event.OpenProcess, event.OpenThread:
		pid, err := e.Params.GetPid()
		if err != nil {
			return e, false, err
		}
		proc := p.psnap.FindAndPut(pid)
		if proc != nil {
			e.AppendParam(params.Exe, params.Path, proc.Exe)
			e.AppendParam(params.ProcessName, params.AnsiString, proc.Name)
		}

		return e, false, nil
	}

	return e, true, nil
}

//nolint:unparam
func (p psProcessor) processEvent(e *event.Event) (*event.Event, error) {
	if e.IsCreateProcessInternal() || e.IsProcessRundownInternal() {
		return e, nil
	}

	cmndline := cmdline.New(e.GetParamAsString(params.Cmdline)).
		// get rid of leading/trailing quotes in the executable path
		CleanExe().
		// expand all variations of the SystemRoot environment variable
		ExpandSystemRoot().
		// some system processes are reported without the path in the command line,
		// but we can expand the path from the SystemRoot environment variable
		CompleteSysProc(e.GetParamAsString(params.ProcessName))

	// append executable path parameter
	exe := cmndline.Exeline()
	if exe == "" {
		exe = e.GetParamAsString(params.ProcessName)
	}
	e.AppendParam(params.Exe, params.Path, exe)

	if e.IsTerminateProcess() {
		return e, nil
	}

	// query process start time
	pid := e.Params.MustGetPid()
	started, err := getStartTime(pid, e)
	if err != nil {
		started = e.Timestamp
	}
	e.AppendParam(params.StartTime, params.Time, started)

	return e, nil
}

func (psProcessor) Name() ProcessorType { return Ps }
func (p psProcessor) Close()            {}

func getStartTime(pid uint32, e *event.Event) (time.Time, error) {
	proc, err := windows.OpenProcess(windows.PROCESS_QUERY_LIMITED_INFORMATION, false, pid)
	if err != nil {
		return e.Timestamp, err
	}
	//nolint:errcheck
	defer windows.CloseHandle(proc)
	var (
		ct windows.Filetime
		xt windows.Filetime
		kt windows.Filetime
		ut windows.Filetime
	)
	err = windows.GetProcessTimes(proc, &ct, &xt, &kt, &ut)
	if err != nil {
		return e.Timestamp, err
	}
	return time.Unix(0, ct.Nanoseconds()), nil
}
