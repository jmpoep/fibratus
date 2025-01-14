/*
 * Copyright 2021-2022 by Nedim Sabic Sabic
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

package va

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sys/windows"
	"os"
	"testing"
	"time"
	"unsafe"
)

func TestRegionProber(t *testing.T) {
	base, err := windows.VirtualAlloc(0, 1024, windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_READONLY)
	require.NoError(t, err)
	defer func() {
		_ = windows.VirtualFree(base, 1024, windows.MEM_DECOMMIT)
	}()
	pid := uint32(os.Getpid())

	regionProber := NewRegionProber()
	region := regionProber.Query(pid, uint64(base))

	require.NotNil(t, region)
	regionProber.mu.Lock()
	require.True(t, regionProber.procs[pid] != windows.InvalidHandle)
	regionProber.mu.Unlock()

	assert.Equal(t, uint32(windows.PAGE_READONLY), region.Protect)
	assert.Equal(t, "R", region.ProtectMask())
	assert.False(t, region.IsMapped())
	assert.Equal(t, MemPrivate, region.Type)

	// test limiter
	for i := 0; i < 499; i++ {
		assert.NotNil(t, regionProber.Query(pid, uint64(base)))
	}
	// all buckets consumed
	assert.Nil(t, regionProber.Query(pid, uint64(base)))

	// wait for the buckets to get refilled
	time.Sleep(time.Second)
	for i := 0; i < 300; i++ {
		assert.NotNil(t, regionProber.Query(pid, uint64(base)))
	}

	require.True(t, regionProber.Remove(pid))
}

func TestReadRegion(t *testing.T) {
	addr, err := getModuleBaseAddress(uint32(os.Getpid()))
	require.NoError(t, err)
	rgn, err := NewRegion(windows.CurrentProcess(), addr)
	require.NoError(t, err)
	require.True(t, rgn.Size(addr) > 0)

	size, b := rgn.Read(addr, uint(os.Getpagesize()), 0x100, false)
	require.True(t, size > 0)
	require.Len(t, b, os.Getpagesize())
	// verify it is the DOS header
	require.Equal(t, 'M', rune(b[0]))
	require.Equal(t, 'Z', rune(b[1]))

	var oldProtect uint32
	_ = windows.VirtualProtectEx(windows.CurrentProcess(), addr, uintptr(rgn.Size(addr)), windows.PAGE_NOACCESS, &oldProtect)

	size, b = rgn.Read(addr, uint(os.Getpagesize()), 0x100, false)
	// shouldn't be able to read the region
	require.True(t, size == 0)
	require.Len(t, b, 0)
	_ = windows.VirtualProtectEx(windows.CurrentProcess(), addr, uintptr(rgn.Size(addr)), oldProtect, &oldProtect)

	_ = windows.VirtualProtectEx(windows.CurrentProcess(), addr, 4096, windows.PAGE_NOACCESS, &oldProtect)
	defer func() {
		_ = windows.VirtualProtectEx(windows.CurrentProcess(), addr, 4096, oldProtect, &oldProtect)
	}()

	noAccessRgn, err := NewRegion(windows.CurrentProcess(), addr)
	require.NoError(t, err)

	size, b = noAccessRgn.Read(addr, uint(os.Getpagesize()), 0x100, true)
	// force protection changing, so should be able to read the region
	require.True(t, size > 0)
	require.Len(t, b, os.Getpagesize())
}

func TestReadArea(t *testing.T) {
	addr, err := getModuleBaseAddress(uint32(os.Getpid()))
	require.NoError(t, err)

	area := ReadArea(windows.CurrentProcess(), addr, uint(os.Getpagesize()), 0x100, false)
	require.Len(t, area, os.Getpagesize())
	require.False(t, Zeroed(area))

	// allocate region with no access protection
	base, err := windows.VirtualAlloc(0, 1024, windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_NOACCESS)
	require.NoError(t, err)
	defer func() {
		_ = windows.VirtualFree(base, 1024, windows.MEM_DECOMMIT)
	}()
	var oldProtect uint32
	_ = windows.VirtualProtectEx(windows.CurrentProcess(), base, 16, windows.PAGE_NOACCESS, &oldProtect)

	// it should read all bytes set to zero
	zeroArea := ReadArea(windows.CurrentProcess(), base, 4096, 0x100, false)
	require.Len(t, zeroArea, 4096)
	require.True(t, Zeroed(zeroArea))
}

func TestQueryWorkingSet(t *testing.T) {
	addr, err := getModuleBaseAddress(uint32(os.Getpid()))
	require.NoError(t, err)

	b := QueryWorkingSet(windows.CurrentProcess(), uint64(addr))
	require.NotNil(t, b)

	require.True(t, b.Valid())
	require.False(t, b.Bad())
	require.True(t, b.SharedOriginal())
	require.True(t, (b.Win32Protection()&windows.PAGE_READONLY) != 0)
}

func getModuleBaseAddress(pid uint32) (uintptr, error) {
	var moduleHandles [1024]windows.Handle
	var cbNeeded uint32
	proc, err := windows.OpenProcess(windows.PROCESS_QUERY_INFORMATION|windows.PROCESS_VM_READ, false, pid)
	if err != nil {
		return 0, err
	}
	if err := windows.EnumProcessModules(proc, &moduleHandles[0], 1024, &cbNeeded); err != nil {
		return 0, err
	}
	moduleHandle := moduleHandles[0]
	var moduleInfo windows.ModuleInfo
	if err := windows.GetModuleInformation(proc, moduleHandle, &moduleInfo, uint32(unsafe.Sizeof(moduleInfo))); err != nil {
		return 0, err
	}
	return moduleInfo.BaseOfDll, nil
}
