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

package fields

import (
	"strings"
	"unicode"

	"github.com/rabbitstack/fibratus/pkg/kevent/kparams"
)

// Field represents the type alias for the field
type Field string

const (
	// PsPid represents the process id field
	PsPid Field = "ps.pid"
	// PsPpid represents the parent process id field
	PsPpid Field = "ps.ppid"
	// PsName represents the process name field
	PsName Field = "ps.name"
	// PsComm represents the process command line field. Deprecated.
	PsComm Field = "ps.comm"
	// PsCmdline represents the process command line field
	PsCmdline Field = "ps.cmdline"
	// PsExe represents the process image path field
	PsExe Field = "ps.exe"
	// PsArgs represents the process command line arguments
	PsArgs Field = "ps.args"
	// PsCwd represents the process current working directory
	PsCwd Field = "ps.cwd"
	// PsSID represents the process security identifier
	PsSID Field = "ps.sid"
	// PsDomain represents the process domain field
	PsDomain Field = "ps.domain"
	// PsUsername represents the process username field
	PsUsername Field = "ps.username"
	// PsSessionID represents the session id bound to the process
	PsSessionID Field = "ps.sessionid"
	// PsEnvs represents the process environment variables
	PsEnvs Field = "ps.envs"
	// PsHandleNames represents the process handles
	PsHandleNames Field = "ps.handles"
	// PsHandleTypes represents the process handle types
	PsHandleTypes Field = "ps.handle.types"
	// PsDTB represents the process directory table base address
	PsDTB Field = "ps.dtb"
	// PsModuleNames represents the process module file names
	PsModuleNames Field = "ps.modules"
	// PsParentPid represents the parent process identifier field
	PsParentPid Field = "ps.parent.pid"
	// PsParentName represents the parent process name field
	PsParentName Field = "ps.parent.name"
	// PsParentComm represents the parent process command line field. Deprecated
	PsParentComm Field = "ps.parent.comm"
	// PsParentCmdline represents the parent process command line field
	PsParentCmdline Field = "ps.parent.cmdline"
	// PsParentExe represents the parent process image path field
	PsParentExe Field = "ps.parent.exe"
	// PsParentArgs represents the parent process command line arguments field
	PsParentArgs Field = "ps.parent.args"
	// PsParentCwd represents the parent process current working directory field
	PsParentCwd Field = "ps.parent.cwd"
	// PsParentSID represents the parent process security identifier field
	PsParentSID Field = "ps.parent.sid"
	// PsParentUsername represents the parent process username field
	PsParentUsername Field = "ps.parent.username"
	// PsParentDomain represents the parent process domain field
	PsParentDomain Field = "ps.parent.domain"
	// PsParentSessionID represents the session id field bound to the parent process
	PsParentSessionID Field = "ps.parent.sessionid"
	// PsParentEnvs represents the parent process environment variables field
	PsParentEnvs Field = "ps.parent.envs"
	// PsParentHandles represents the parent process handles field
	PsParentHandles Field = "ps.parent.handles"
	// PsParentHandleTypes represents the parent process handle types field
	PsParentHandleTypes Field = "ps.parent.handle.types"
	// PsParentDTB represents the parent process directory table base address field
	PsParentDTB Field = "ps.parent.dtb"
	// PsAncestor represents the process ancestor field
	PsAncestor Field = "ps.ancestor"
	// PsAccessMask represents the process access rights field
	PsAccessMask Field = "ps.access.mask"
	// PsAccessMaskNames represents the process access rights list field
	PsAccessMaskNames Field = "ps.access.mask.names"
	// PsAccessStatus represents the process access status field
	PsAccessStatus Field = "ps.access.status"
	// PsIsWOW64Field represents the field that indicates if the 32-bit process is created in 64-bit Windows system
	PsIsWOW64Field Field = "ps.is_wow64"
	// PsIsPackagedField represents the field that indicates if a process is packaged with the MSIX technology
	PsIsPackagedField Field = "ps.is_packaged"
	// PsIsProtectedField represents the field that indicates if the process is to be run as a protected process
	PsIsProtectedField Field = "ps.is_protected"
	// PsParentIsWOW64Field represents the field that indicates if the 32-bit process is created in 64-bit Windows system
	PsParentIsWOW64Field Field = "ps.parent.is_wow64"
	// PsParentIsPackagedField represents the field that indicates if a process is packaged with the MSIX technology
	PsParentIsPackagedField Field = "ps.parent.is_packaged"
	// PsParentIsProtectedField represents the field that indicates if the process is to be run as a protected process
	PsParentIsProtectedField Field = "ps.parent.is_protected"

	// PsSiblingPid represents the sibling process identifier field. Deprecated
	PsSiblingPid Field = "ps.sibling.pid"
	// PsSiblingName represents the sibling process name field. Deprecated
	PsSiblingName Field = "ps.sibling.name"
	// PsSiblingComm represents the sibling process command line field. Deprecated
	PsSiblingComm Field = "ps.sibling.comm"
	// PsSiblingExe represents the sibling process complete executable path field. Deprecated
	PsSiblingExe Field = "ps.sibling.exe"
	// PsSiblingArgs represents the sibling process command line arguments path field. Deprecated
	PsSiblingArgs Field = "ps.sibling.args"
	// PsSiblingSID represents the sibling process security identifier field. Deprecated
	PsSiblingSID Field = "ps.sibling.sid"
	// PsSiblingSessionID represents the sibling process session id field. Deprecated
	PsSiblingSessionID Field = "ps.sibling.sessionid"
	// PsSiblingDomain represents the sibling process domain field. Deprecated
	PsSiblingDomain Field = "ps.sibling.domain"
	// PsSiblingUsername represents the sibling process username field. Deprecated
	PsSiblingUsername Field = "ps.sibling.username"
	// PsUUID represents the unique process identifier
	PsUUID Field = "ps.uuid"
	// PsParentUUID represents the unique parent process identifier
	PsParentUUID Field = "ps.parent.uuid"
	// PsChildUUID represents the unique child process identifier
	PsChildUUID Field = "ps.child.uuid"

	// PsChildPid represents the child process identifier field
	PsChildPid Field = "ps.child.pid"
	// PsChildName represents the child process name field
	PsChildName Field = "ps.child.name"
	// PsChildCmdline represents the child process command line field
	PsChildCmdline Field = "ps.child.cmdline"
	// PsChildExe represents the child process complete executable path field
	PsChildExe Field = "ps.child.exe"
	// PsChildArgs represents the child process command line arguments path field
	PsChildArgs Field = "ps.child.args"
	// PsChildSID represents the child process security identifier field
	PsChildSID Field = "ps.child.sid"
	// PsChildSessionID represents the child process session id field
	PsChildSessionID Field = "ps.child.sessionid"
	// PsChildDomain represents the child process domain field
	PsChildDomain Field = "ps.child.domain"
	// PsChildUsername represents the child process username field
	PsChildUsername Field = "ps.child.username"
	// PsChildPeFilename represents the original file name of the child process executable provided at compile-time
	PsChildPeFilename Field = "ps.child.pe.file.name"
	// PsChildIsWOW64Field  represents the field that indicates if the 32-bit process is created in 64-bit Windows system
	PsChildIsWOW64Field Field = "ps.child.is_wow64"
	// PsChildIsPackagedField represents the field that indicates if a process is packaged with the MSIX technology
	PsChildIsPackagedField Field = "ps.child.is_packaged"
	// PsChildIsProtectedField represents the field that indicates if the process is to be run as a protected process
	PsChildIsProtectedField Field = "ps.child.is_protected"

	// ThreadBasePrio is the base thread priority
	ThreadBasePrio Field = "thread.prio"
	// ThreadIOPrio is the thread I/O priority
	ThreadIOPrio Field = "thread.io.prio"
	// ThreadPagePrio is the thread page priority
	ThreadPagePrio Field = "thread.page.prio"
	// ThreadKstackBase is the thread kernel stack start address
	ThreadKstackBase Field = "thread.kstack.base"
	// ThreadKstackLimit is the thread kernel stack end address
	ThreadKstackLimit Field = "thread.kstack.limit"
	// ThreadUstackBase is the thread user stack start address
	ThreadUstackBase Field = "thread.ustack.base"
	// ThreadUstackLimit is the thead user stack end address
	ThreadUstackLimit Field = "thread.ustack.limit"
	// ThreadEntrypoint is the thread entrypoint address
	ThreadEntrypoint Field = "thread.entrypoint"
	// ThreadStartAddress is the thread start address
	ThreadStartAddress Field = "thread.start_address"
	// ThreadPID is the process identifier where the thread is created
	ThreadPID Field = "thread.pid"
	// ThreadTEB is the thread environment block base address
	ThreadTEB Field = "thread.teb_address"
	// ThreadAccessMask represents the thread access rights field
	ThreadAccessMask Field = "thread.access.mask"
	// ThreadAccessMaskNames represents the thread access rights list field
	ThreadAccessMaskNames Field = "thread.access.mask.names"
	// ThreadAccessStatus represents the thread access status field
	ThreadAccessStatus Field = "thread.access.status"
	// ThreadCallstackSummary represents the thread callstack summary field
	ThreadCallstackSummary Field = "thread.callstack.summary"
	// ThreadCallstackDetail represents the thread callstack detail field
	ThreadCallstackDetail Field = "thread.callstack.detail"
	// ThreadCallstackModules represents the callstack modules field
	ThreadCallstackModules Field = "thread.callstack.modules"
	// ThreadCallstackSymbols represents the callstack symbols field
	ThreadCallstackSymbols Field = "thread.callstack.symbols"
	// ThreadCallstackProtections represents the callstack region protections field
	ThreadCallstackProtections Field = "thread.callstack.protections"
	// ThreadCallstackAllocationSizes represents the private region page sizes field
	ThreadCallstackAllocationSizes Field = "thread.callstack.allocation_sizes"
	// ThreadCallstackCallsiteLeadingAssembly represents the callsite prelude opcodes field
	ThreadCallstackCallsiteLeadingAssembly Field = "thread.callstack.callsite_leading_assembly"
	// ThreadCallstackCallsiteTrailingAssembly represents the callsite postlude opcodes field
	ThreadCallstackCallsiteTrailingAssembly Field = "thread.callstack.callsite_trailing_assembly"
	// ThreadCallstackIsUnbacked represents the field that indicates if there is an unbacked stack frame
	ThreadCallstackIsUnbacked Field = "thread.callstack.is_unbacked"
	// ThreadStartAddressSymbol represents the symbol corresponding to the thread start address
	ThreadStartAddressSymbol Field = "thread.start_address.symbol"
	// ThreadStartAddressModule represents the module corresponding to the thread start address
	ThreadStartAddressModule Field = "thread.start_address.module"

	// PeNumSections represents the number of sections
	PeNumSections Field = "pe.nsections"
	// PeNumSymbols represents the number of exported symbols
	PeNumSymbols Field = "pe.nsymbols"
	// PeSymbols represents imported symbols
	PeSymbols Field = "pe.symbols"
	// PeImports represents imported libraries (e.g. kernel32.dll)
	PeImports Field = "pe.imports"
	// PeTimestamp is the PE build timestamp
	PeTimestamp Field = "pe.timestamp"
	// PeBaseAddress represents the base address when the binary is loaded
	PeBaseAddress Field = "pe.address.base"
	// PeEntrypoint is the address of the entrypoint function
	PeEntrypoint Field = "pe.address.entrypoint"
	// PeResources represents PE resources
	PeResources Field = "pe.resources"
	// PeCompany represents the company name resource
	PeCompany Field = "pe.company"
	// PeDescription represents the internal description of the file
	PeDescription Field = "pe.description"
	// PeFileVersion represents the internal file version
	PeFileVersion Field = "pe.file.version"
	// PeFileName represents the original file name provided at compile-time.
	PeFileName Field = "pe.file.name"
	// PeCopyright represents the copyright notice emitted at compile-time
	PeCopyright Field = "pe.copyright"
	// PeProduct represents the product name provided at compile-time
	PeProduct Field = "pe.product"
	// PeProductVersion represents the internal product version provided at compile-time
	PeProductVersion Field = "pe.product.version"
	// PeIsDLL indicates if the file is a DLL
	PeIsDLL Field = "pe.is_dll"
	// PeIsDriver indicates if the file is a driver
	PeIsDriver Field = "pe.is_driver"
	// PeIsExecutable indicates if the file is an executable
	PeIsExecutable Field = "pe.is_exec"
	// PeAnomalies represents the field that contains PE anomalies detected during parsing
	PeAnomalies Field = "pe.anomalies"
	// PeImphash is the field that yields the PE import hash
	PeImphash Field = "pe.imphash"
	// PeIsDotnet is the field which indicates if the binary contains the .NET assembly
	PeIsDotnet Field = "pe.is_dotnet"
	// PeIsSigned is the field which indicates if the binary is signed, either by embedded or catalog signature
	PeIsSigned Field = "pe.is_signed"
	// PeIsTrusted is the field which indicates if the binary signature is trusted
	PeIsTrusted Field = "pe.is_trusted"
	// PeCertIssuer is the field which indicates the certificate issuer
	PeCertIssuer Field = "pe.cert.issuer"
	// PeCertSubject is the field which indicates the certificate subject
	PeCertSubject Field = "pe.cert.subject"
	// PeCertSerial is the field which indicates the certificate serial
	PeCertSerial Field = "pe.cert.serial"
	// PeCertAfter is the field which indicates the timestamp after certificate is no longer valid
	PeCertAfter Field = "pe.cert.after"
	// PeCertBefore is the field which indicates the timestamp of the certificate enrollment date
	PeCertBefore Field = "pe.cert.before"
	// PeIsModified is the field that indicates whether disk and in-memory PE headers differ
	PeIsModified Field = "pe.is_modified"
	// PePsChildFileName represents the original file name of the child process executable provided at compile-time
	PePsChildFileName Field = "pe.ps.child.file.name"

	// KevtSeq is the event sequence number
	KevtSeq Field = "kevt.seq"
	// KevtPID is the process identifier that generated the event
	KevtPID Field = "kevt.pid"
	// KevtTID is the thread identifier that generated the event
	KevtTID Field = "kevt.tid"
	// KevtCPU is the CPU core where the event was generated
	KevtCPU Field = "kevt.cpu"
	// KevtDesc represents the event description
	KevtDesc Field = "kevt.desc"
	// KevtHost represents the host where the event was produced
	KevtHost Field = "kevt.host"
	// KevtTime is the event time
	KevtTime Field = "kevt.time"
	// KevtTimeHour is the hour part of the event time
	KevtTimeHour Field = "kevt.time.h"
	// KevtTimeMin is the minute part of the event time
	KevtTimeMin Field = "kevt.time.m"
	// KevtTimeSec is the second part of the event time
	KevtTimeSec Field = "kevt.time.s"
	// KevtTimeNs is the nanosecond part of the event time
	KevtTimeNs Field = "kevt.time.ns"
	// KevtDate is the event date
	KevtDate Field = "kevt.date"
	// KevtDateDay is the day of event date
	KevtDateDay Field = "kevt.date.d"
	// KevtDateMonth is the month of event date
	KevtDateMonth Field = "kevt.date.m"
	// KevtDateYear is the year of event date
	KevtDateYear Field = "kevt.date.y"
	// KevtDateTz is the time zone of event timestamp
	KevtDateTz Field = "kevt.date.tz"
	// KevtDateWeek is the event week number
	KevtDateWeek Field = "kevt.date.week"
	// KevtDateWeekday is the event week day
	KevtDateWeekday Field = "kevt.date.weekday"
	// KevtName is the event name
	KevtName Field = "kevt.name"
	// KevtCategory is the event category
	KevtCategory Field = "kevt.category"
	// KevtMeta is the event metadata
	KevtMeta Field = "kevt.meta"
	// KevtNparams is the number of event parameters
	KevtNparams Field = "kevt.nparams"
	// KevtArg represents the field sequence for generic argument access
	KevtArg Field = "kevt.arg"

	// HandleID represents the handle identifier within the process address space
	HandleID Field = "handle.id"
	// HandleObject represents the handle object address
	HandleObject Field = "handle.object"
	// HandleName represents the handle name
	HandleName Field = "handle.name"
	// HandleType represents the handle type (e.g. file)
	HandleType Field = "handle.type"

	// NetDIP represents network destination IP address
	NetDIP Field = "net.dip"
	// NetSIP represents the source IP address
	NetSIP Field = "net.sip"
	// NetDport represents the destination port
	NetDport Field = "net.dport"
	// NetSport represents the source port
	NetSport Field = "net.sport"
	// NetDportName represents the destination port IANA name
	NetDportName Field = "net.dport.name"
	// NetSportName represents the source port IANA name
	NetSportName Field = "net.sport.name"
	// NetL4Proto represents the Layer4 protocol name (e.g. TCP)
	NetL4Proto Field = "net.l4.proto"
	// NetPacketSize represents the packet size
	NetPacketSize Field = "net.size"
	// NetSIPNames represents the source IP names
	NetSIPNames Field = "net.sip.names"
	// NetDIPNames represents the destination IP names
	NetDIPNames Field = "net.dip.names"

	// FileObject represents the address of the file object
	FileObject Field = "file.object"
	// FileName represents the file base name (e.g. cmd.exe)
	FileName Field = "file.name"
	// FilePath represents the file full path (e.g. C:\Windows\System32\cmd.exe)
	FilePath Field = "file.path"
	// FileExtension represents the file extension (e.g. .exe or .dll)
	FileExtension Field = "file.extension"
	// FileOperation represents the file operation (e.g. create)
	FileOperation Field = "file.operation"
	// FileShareMask represents the file share mask
	FileShareMask Field = "file.share.mask"
	// FileIOSize represents the number of read/written bytes
	FileIOSize Field = "file.io.size"
	// FileOffset represents the read/write offset
	FileOffset Field = "file.offset"
	// FileType represents the file type
	FileType Field = "file.type"
	// FileAttributes represents a slice of file attributes
	FileAttributes Field = "file.attributes"
	// FileStatus represents the status message of the file operation
	FileStatus Field = "file.status"
	// FileViewBase represents the base address of the mapped view
	FileViewBase Field = "file.view.base"
	// FileViewSize represents the size of the mapped view
	FileViewSize Field = "file.view.size"
	// FileViewType represents the type of the mapped view section
	FileViewType Field = "file.view.type"
	// FileViewProtection represents the protection attributes of the section view
	FileViewProtection Field = "file.view.protection"
	// FileIsDriverVulnerable represents the field that denotes whether the created file is a vulnerable driver
	FileIsDriverVulnerable Field = "file.is_driver_vulnerable"
	// FileIsDriverMalicious represents the field that denotes whether the created file is a malicious driver
	FileIsDriverMalicious Field = "file.is_driver_malicious"
	// FileIsDLL indicates if the created file is a DLL
	FileIsDLL Field = "file.is_dll"
	// FileIsDriver indicates if the created file is a driver
	FileIsDriver Field = "file.is_driver"
	// FileIsExecutable indicates if the created file is an executable
	FileIsExecutable Field = "file.is_exec"
	// FilePID represents the field that denotes the process id performing file operations
	FilePID Field = "file.pid"
	// FileKey represents the field that uniquely identifies the file object.
	FileKey Field = "file.key"
	// FileInfoClass represents the field that identifies the file information class
	FileInfoClass Field = "file.info_class"
	// FileInfoAllocationSize represents the field that contains the file allocation size
	FileInfoAllocationSize Field = "file.info.allocation_size"
	// FileInfoEOFSize represents the field that contains the EOF size
	FileInfoEOFSize Field = "file.info.eof_size"
	// FileInfoIsDispositionDeleteFile represents the field that indicates if the file is deleted when its handle is closed
	FileInfoIsDispositionDeleteFile Field = "file.info.is_disposition_delete_file"

	// RegistryPath represents the full registry path
	RegistryPath Field = "registry.path"
	// RegistryKeyName represents the registry key name
	RegistryKeyName Field = "registry.key.name"
	// RegistryKeyHandle represents the registry KCB address
	RegistryKeyHandle Field = "registry.key.handle"
	// RegistryValue represents the registry value
	RegistryValue Field = "registry.value"
	// RegistryValueType represents the registry value type
	RegistryValueType Field = "registry.value.type"
	// RegistryStatus represent the registry operation status
	RegistryStatus Field = "registry.status"

	// ImageBase is the module base address
	ImageBase Field = "image.base.address"
	// ImageSize is the module size
	ImageSize Field = "image.size"
	// ImageChecksum represents the module checksum hash
	ImageChecksum Field = "image.checksum"
	// ImageDefaultAddress represents the module address
	ImageDefaultAddress Field = "image.default.address"
	// ImagePath is the module full path
	ImagePath Field = "image.path"
	// ImageName is the module name
	ImageName Field = "image.name"
	// ImagePID is the pid of the process where the image was loaded
	ImagePID Field = "image.pid"
	// ImageSignatureType represents the image signature type
	ImageSignatureType Field = "image.signature.type"
	// ImageSignatureLevel represents the image signature level
	ImageSignatureLevel Field = "image.signature.level"
	// ImageCertSubject is the field that indicates the subject of the certificate is the entity its public key is associated with.
	ImageCertSubject = "image.cert.subject"
	// ImageCertIssuer is the field that represents the certificate authority (CA).
	ImageCertIssuer = "image.cert.issuer"
	// ImageCertSerial is the field that represents the serial number MUST be a positive integer assigned
	// by the CA to each certificate.
	ImageCertSerial = "image.cert.serial"
	// ImageCertBefore is the field that specifies the certificate won't be valid before this timestamp.
	ImageCertBefore = "image.cert.before"
	// ImageCertAfter is the field that specifies the certificate won't be valid after this timestamp.
	ImageCertAfter = "image.cert.after"
	// ImageIsDriverVulnerable represents the field that denotes whether loaded driver is vulnerable
	ImageIsDriverVulnerable Field = "image.is_driver_vulnerable"
	// ImageIsDriverMalicious represents the field that denotes whether the loaded driver is malicious
	ImageIsDriverMalicious Field = "image.is_driver_malicious"
	// ImageIsDLL indicates if the loaded image is a DLL
	ImageIsDLL Field = "image.is_dll"
	// ImageIsDriver indicates if the loaded image is a driver
	ImageIsDriver Field = "image.is_driver"
	// ImageIsExecutable indicates if the loaded image is an executable
	ImageIsExecutable Field = "image.is_exec"
	// ImageIsDotnet indicates if the loaded image is a .NET assembly
	ImageIsDotnet Field = "image.is_dotnet"

	// MemBaseAddress identifies the field that denotes the allocation base address
	MemBaseAddress Field = "mem.address"
	// MemRegionSize Field identifies the field that represents the allocated region size
	MemRegionSize Field = "mem.size"
	// MemAllocType identifies the field that represents region allocation type
	MemAllocType Field = "mem.alloc"
	// MemPageType identifies the parameter that represents the allocated region type
	MemPageType Field = "mem.type"
	// MemProtection identifies the field that represents the memory protection for the range of pages
	MemProtection Field = "mem.protection"
	// MemProtectionMask identifies the field that represents the memory protection in mask notation
	MemProtectionMask Field = "mem.protection.mask"

	// DNSName identifies the field that represents the DNS name
	DNSName Field = "dns.name"
	// DNSRR identifies the field that represents the DNS record type
	DNSRR Field = "dns.rr"
	// DNSOptions identifies the field that represents the DNS options
	DNSOptions Field = "dns.options"
	// DNSAnswers identifies the field that represents the DNS answers
	DNSAnswers Field = "dns.answers"
	// DNSRcode identifies the field that represents the DNS response code
	DNSRcode Field = "dns.rcode"
)

// String casts the field type to string.
func (f Field) String() string { return string(f) }

func (f Field) IsPsField() bool       { return strings.HasPrefix(string(f), "ps.") }
func (f Field) IsKevtField() bool     { return strings.HasPrefix(string(f), "kevt.") }
func (f Field) IsThreadField() bool   { return strings.HasPrefix(string(f), "thread.") }
func (f Field) IsImageField() bool    { return strings.HasPrefix(string(f), "image.") }
func (f Field) IsFileField() bool     { return strings.HasPrefix(string(f), "file.") }
func (f Field) IsRegistryField() bool { return strings.HasPrefix(string(f), "registry.") }
func (f Field) IsNetworkField() bool  { return strings.HasPrefix(string(f), "net.") }
func (f Field) IsHandleField() bool   { return strings.HasPrefix(string(f), "handle.") }
func (f Field) IsPeField() bool       { return strings.HasPrefix(string(f), "pe.") || f == PsChildPeFilename }
func (f Field) IsMemField() bool      { return strings.HasPrefix(string(f), "mem.") }
func (f Field) IsDNSField() bool      { return strings.HasPrefix(string(f), "dns.") }

func (f Field) IsPeSection() bool { return f == PeNumSections }
func (f Field) IsPeSymbol() bool  { return f == PeSymbols || f == PeNumSymbols || f == PeImports }
func (f Field) IsPeVersionResource() bool {
	return f == PeCompany || f == PeCopyright || f == PeDescription || f == PeFileName || f == PeFileVersion || f == PeProduct || f == PeProductVersion || f == PePsChildFileName || f == PsChildPeFilename
}
func (f Field) IsPeVersionResources() bool { return f == PeResources }
func (f Field) IsPeImphash() bool          { return f == PeImphash }
func (f Field) IsPeDotnet() bool           { return f == PeIsDotnet }
func (f Field) IsPeAnomalies() bool        { return f == PeAnomalies }
func (f Field) IsPeSignature() bool {
	return f == PeIsTrusted || f == PeIsSigned || f == PeCertIssuer || f == PeCertSerial || f == PeCertSubject || f == PeCertBefore || f == PeCertAfter
}
func (f Field) IsPeIsTrusted() bool { return f == PeIsTrusted }
func (f Field) IsPeIsSigned() bool  { return f == PeIsSigned }

func (f Field) IsPeCert() bool    { return strings.HasPrefix(string(f), "pe.cert.") }
func (f Field) IsImageCert() bool { return strings.HasPrefix(string(f), "image.cert.") }

func (f Field) IsPeModified() bool { return f == PeIsModified }

// Segment represents the type alias for the segment. Segment
// denotes the property anchored to the bound field reference.
// Let's look through an example. $module.name is the literal
// composed of bound field ($module) and the segment (name).
// Segments are most commonly used in the context of bound
// variables in foreach function.
type Segment string

const (
	PathSegment     Segment = "path"
	NameSegment     Segment = "name"
	TypeSegment     Segment = "type"
	SizeSegment     Segment = "size"
	ChecksumSegment Segment = "checksum"
	AddressSegment  Segment = "address"
	OffsetSegment   Segment = "offset"
	EntropySegment  Segment = "entropy"
	MD5Segment      Segment = "md5"

	PIDSegment       Segment = "pid"
	CmdlineSegment   Segment = "cmdline"
	ExeSegment       Segment = "exe"
	ArgsSegment      Segment = "args"
	CwdSegment       Segment = "cwd"
	SIDSegment       Segment = "sid"
	SessionIDSegment Segment = "sessionid"
	UsernameSegment  Segment = "username"
	DomainSegment    Segment = "domain"

	TidSegment              Segment = "tid"
	StartAddressSegment     Segment = "start_address"
	UserStackBaseSegment    Segment = "user_stack_base"
	UserStackLimitSegment   Segment = "user_stack_limit"
	KernelStackBaseSegment  Segment = "kernel_stack_base"
	KernelStackLimitSegment Segment = "kernel_stack_limit"

	SymbolSegment                   Segment = "symbol"
	ModuleSegment                   Segment = "module"
	AllocationSizeSegment           Segment = "allocation_size"
	ProtectionSegment               Segment = "protection"
	IsUnbackedSegment               Segment = "is_unbacked"
	CallsiteLeadingAssemblySegment  Segment = "callsite_leading_assembly"
	CallsiteTrailingAssemblySegment Segment = "callsite_trailing_assembly"
)

var segments = map[Segment]bool{
	NameSegment:                     true,
	PathSegment:                     true,
	TypeSegment:                     true,
	EntropySegment:                  true,
	SizeSegment:                     true,
	MD5Segment:                      true,
	AddressSegment:                  true,
	ChecksumSegment:                 true,
	PIDSegment:                      true,
	CmdlineSegment:                  true,
	ExeSegment:                      true,
	ArgsSegment:                     true,
	CwdSegment:                      true,
	SIDSegment:                      true,
	SessionIDSegment:                true,
	UsernameSegment:                 true,
	DomainSegment:                   true,
	TidSegment:                      true,
	StartAddressSegment:             true,
	UserStackBaseSegment:            true,
	UserStackLimitSegment:           true,
	KernelStackBaseSegment:          true,
	KernelStackLimitSegment:         true,
	OffsetSegment:                   true,
	SymbolSegment:                   true,
	ModuleSegment:                   true,
	AllocationSizeSegment:           true,
	ProtectionSegment:               true,
	IsUnbackedSegment:               true,
	CallsiteLeadingAssemblySegment:  true,
	CallsiteTrailingAssemblySegment: true,
}

var allowedSegments = map[Field][]Segment{
	PsAncestors:     {NameSegment, PIDSegment, CmdlineSegment, ExeSegment, ArgsSegment, CwdSegment, SIDSegment, SessionIDSegment, UsernameSegment, DomainSegment},
	PsThreads:       {TidSegment, StartAddressSegment, UserStackBaseSegment, UserStackLimitSegment, KernelStackBaseSegment, KernelStackLimitSegment},
	PsModules:       {PathSegment, NameSegment, AddressSegment, SizeSegment, ChecksumSegment},
	PsMmaps:         {AddressSegment, TypeSegment, AddressSegment, SizeSegment, ProtectionSegment, PathSegment},
	PeSections:      {NameSegment, SizeSegment, EntropySegment, MD5Segment},
	ThreadCallstack: {AddressSegment, OffsetSegment, SymbolSegment, ModuleSegment, AllocationSizeSegment, ProtectionSegment, IsUnbackedSegment, CallsiteLeadingAssemblySegment, CallsiteTrailingAssemblySegment},
}

func (s Segment) IsEntropy() bool { return s == EntropySegment }

// IsSegmentAllowed determines if the segment is valid for the pseudo field.
func IsSegmentAllowed(f Field, s Segment) bool {
	segs := allowedSegments[f]
	if len(segs) == 0 {
		return false
	}

	for _, seg := range segs {
		if seg == s {
			return true
		}
	}

	return false
}

// SegmentsHint returns the sequence of available segments for the pseudo field.
func SegmentsHint(f Field) string {
	segs := allowedSegments[f]
	if len(segs) == 0 {
		return ""
	}

	s := make([]string, len(segs))
	for i, seg := range segs {
		s[i] = string(seg)
	}

	return strings.Join(s, ", ")
}

// IsSegment indicates if the given string is recognized as a known segment.
func IsSegment(s string) bool {
	return segments[Segment(s)]
}

// Pseudo fields provide access to the process/event internal state. They
// are typically used in conjunction with the foreach function as its
// first argument.

var (
	PsModules       Field = "ps._modules"
	PsThreads       Field = "ps._threads"
	PsMmaps         Field = "ps._mmaps"
	PsAncestors     Field = "ps._ancestors"
	ThreadCallstack Field = "thread._callstack"
	PeSections      Field = "pe._sections"
)

func IsPseudoField(f Field) bool {
	return f == PsAncestors || f == PsModules || f == PsThreads || f == PsMmaps || f == ThreadCallstack || f == PeSections
}

func (f Field) IsPeSectionsPseudo() bool { return f == PeSections }

var fields = map[Field]FieldInfo{
	KevtSeq:         {KevtSeq, "event sequence number", kparams.Uint64, []string{"kevt.seq > 666"}, nil, nil},
	KevtPID:         {KevtPID, "process identifier generating the kernel event", kparams.Uint32, []string{"kevt.pid = 6"}, nil, nil},
	KevtTID:         {KevtTID, "thread identifier generating the kernel event", kparams.Uint32, []string{"kevt.tid = 1024"}, nil, nil},
	KevtCPU:         {KevtCPU, "logical processor core where the event was generated", kparams.Uint8, []string{"kevt.cpu = 2"}, nil, nil},
	KevtName:        {KevtName, "symbolical kernel event name", kparams.AnsiString, []string{"kevt.name = 'CreateThread'"}, nil, nil},
	KevtCategory:    {KevtCategory, "event category", kparams.AnsiString, []string{"kevt.category = 'registry'"}, nil, nil},
	KevtDesc:        {KevtDesc, "event description", kparams.AnsiString, []string{"kevt.desc contains 'Creates a new process'"}, nil, nil},
	KevtHost:        {KevtHost, "host name on which the event was produced", kparams.UnicodeString, []string{"kevt.host contains 'kitty'"}, nil, nil},
	KevtTime:        {KevtTime, "event timestamp as a time string", kparams.Time, []string{"kevt.time = '17:05:32'"}, nil, nil},
	KevtTimeHour:    {KevtTimeHour, "hour within the day on which the event occurred", kparams.Time, []string{"kevt.time.h = 23"}, nil, nil},
	KevtTimeMin:     {KevtTimeMin, "minute offset within the hour on which the event occurred", kparams.Time, []string{"kevt.time.m = 54"}, nil, nil},
	KevtTimeSec:     {KevtTimeSec, "second offset within the minute  on which the event occurred", kparams.Time, []string{"kevt.time.s = 0"}, nil, nil},
	KevtTimeNs:      {KevtTimeNs, "nanoseconds specified by event timestamp", kparams.Int64, []string{"kevt.time.ns > 1591191629102337000"}, nil, nil},
	KevtDate:        {KevtDate, "event timestamp as a date string", kparams.Time, []string{"kevt.date = '2018-03-03'"}, nil, nil},
	KevtDateDay:     {KevtDateDay, "day of the month on which the event occurred", kparams.Time, []string{"kevt.date.d = 12"}, nil, nil},
	KevtDateMonth:   {KevtDateMonth, "month of the year on which the event occurred", kparams.Time, []string{"kevt.date.m = 11"}, nil, nil},
	KevtDateYear:    {KevtDateYear, "year on which the event occurred", kparams.Uint32, []string{"kevt.date.y = 2020"}, nil, nil},
	KevtDateTz:      {KevtDateTz, "time zone associated with the event timestamp", kparams.AnsiString, []string{"kevt.date.tz = 'UTC'"}, nil, nil},
	KevtDateWeek:    {KevtDateWeek, "week number within the year on which the event occurred", kparams.Uint8, []string{"kevt.date.week = 2"}, nil, nil},
	KevtDateWeekday: {KevtDateWeekday, "week day on which the event occurred", kparams.AnsiString, []string{"kevt.date.weekday = 'Monday'"}, nil, nil},
	KevtNparams:     {KevtNparams, "number of parameters", kparams.Int8, []string{"kevt.nparams > 2"}, nil, nil},
	KevtArg: {KevtArg, "event parameter", kparams.Object, []string{"kevt.arg[cmdline] istartswith 'C:\\Windows'"}, nil, &Argument{Optional: false, Pattern: "[a-z0-9_]+", ValidationFunc: func(s string) bool {
		for _, c := range s {
			switch {
			case unicode.IsLower(c):
			case unicode.IsNumber(c):
			case c == '_':
			default:
				return false
			}
		}
		return true
	}}},

	PsPid:                    {PsPid, "process identifier", kparams.PID, []string{"ps.pid = 1024"}, nil, nil},
	PsPpid:                   {PsPpid, "parent process identifier", kparams.PID, []string{"ps.ppid = 45"}, nil, nil},
	PsName:                   {PsName, "process image name including the file extension", kparams.UnicodeString, []string{"ps.name contains 'firefox'"}, nil, nil},
	PsComm:                   {PsComm, "process command line", kparams.UnicodeString, []string{"ps.comm contains 'java'"}, &Deprecation{Since: "1.10.0", Fields: []Field{PsCmdline}}, nil},
	PsCmdline:                {PsCmdline, "process command line", kparams.UnicodeString, []string{"ps.cmdline contains 'java'"}, nil, nil},
	PsExe:                    {PsExe, "full name of the process' executable", kparams.UnicodeString, []string{"ps.exe = 'C:\\Windows\\system32\\cmd.exe'"}, nil, nil},
	PsArgs:                   {PsArgs, "process command line arguments", kparams.Slice, []string{"ps.args in ('/cdir', '/-C')"}, nil, nil},
	PsCwd:                    {PsCwd, "process current working directory", kparams.UnicodeString, []string{"ps.cwd = 'C:\\Users\\Default'"}, nil, nil},
	PsSID:                    {PsSID, "security identifier under which this process is run", kparams.UnicodeString, []string{"ps.sid contains 'SYSTEM'"}, nil, nil},
	PsSessionID:              {PsSessionID, "unique identifier for the current session", kparams.Int16, []string{"ps.sessionid = 1"}, nil, nil},
	PsDomain:                 {PsDomain, "process domain", kparams.UnicodeString, []string{"ps.domain contains 'SERVICE'"}, nil, nil},
	PsUsername:               {PsUsername, "process username", kparams.UnicodeString, []string{"ps.username contains 'system'"}, nil, nil},
	PsEnvs:                   {PsEnvs, "process environment variables", kparams.Slice, []string{"ps.envs in ('SystemRoot:C:\\WINDOWS')", "ps.envs[windir] = 'C:\\WINDOWS'"}, nil, &Argument{Optional: true, ValidationFunc: func(arg string) bool { return true }}},
	PsHandleNames:            {PsHandleNames, "allocated process handle names", kparams.Slice, []string{"ps.handles in ('\\BaseNamedObjects\\__ComCatalogCache__')"}, nil, nil},
	PsHandleTypes:            {PsHandleTypes, "allocated process handle types", kparams.Slice, []string{"ps.handle.types in ('Key', 'Mutant', 'Section')"}, nil, nil},
	PsDTB:                    {PsDTB, "process directory table base address", kparams.Address, []string{"ps.dtb = '7ffe0000'"}, nil, nil},
	PsModuleNames:            {PsModuleNames, "modules loaded by the process", kparams.Slice, []string{"ps.modules in ('crypt32.dll', 'xul.dll')"}, nil, nil},
	PsParentName:             {PsParentName, "parent process image name including the file extension", kparams.UnicodeString, []string{"ps.parent.name contains 'cmd.exe'"}, nil, nil},
	PsParentPid:              {PsParentPid, "parent process id", kparams.Uint32, []string{"ps.parent.pid = 4"}, nil, nil},
	PsParentComm:             {PsParentComm, "parent process command line", kparams.UnicodeString, []string{"ps.parent.comm contains 'java'"}, &Deprecation{Since: "1.10.0", Fields: []Field{PsParentCmdline}}, nil},
	PsParentCmdline:          {PsParentCmdline, "parent process command line", kparams.UnicodeString, []string{"ps.parent.cmdline contains 'java'"}, nil, nil},
	PsParentExe:              {PsParentExe, "full name of the parent process' executable", kparams.UnicodeString, []string{"ps.parent.exe = 'C:\\Windows\\system32\\explorer.exe'"}, nil, nil},
	PsParentArgs:             {PsParentArgs, "parent process command line arguments", kparams.Slice, []string{"ps.parent.args in ('/cdir', '/-C')"}, nil, nil},
	PsParentCwd:              {PsParentCwd, "parent process current working directory", kparams.UnicodeString, []string{"ps.parent.cwd = 'C:\\Temp'"}, nil, nil},
	PsParentSID:              {PsParentSID, "security identifier under which the parent process is run", kparams.UnicodeString, []string{"ps.parent.sid contains 'SYSTEM'"}, nil, nil},
	PsParentDomain:           {PsParentDomain, "parent process domain", kparams.UnicodeString, []string{"ps.parent.domain contains 'SERVICE'"}, nil, nil},
	PsParentUsername:         {PsParentUsername, "parent process username", kparams.UnicodeString, []string{"ps.parent.username contains 'system'"}, nil, nil},
	PsParentSessionID:        {PsParentSessionID, "unique identifier for the current session of parent process", kparams.Int16, []string{"ps.parent.sessionid = 1"}, nil, nil},
	PsParentEnvs:             {PsParentEnvs, "parent process environment variables", kparams.Slice, []string{"ps.parent.envs in ('MOZ_CRASHREPORTER_DATA_DIRECTORY')"}, nil, nil},
	PsParentHandles:          {PsParentHandles, "allocated parent process handle names", kparams.Slice, []string{"ps.parent.handles in ('\\BaseNamedObjects\\__ComCatalogCache__')"}, nil, nil},
	PsParentHandleTypes:      {PsParentHandleTypes, "allocated parent process handle types", kparams.Slice, []string{"ps.parent.handle.types in ('File', 'SymbolicLink')"}, nil, nil},
	PsParentDTB:              {PsParentDTB, "parent process directory table base address", kparams.Address, []string{"ps.parent.dtb = '7ffe0000'"}, nil, nil},
	PsAccessMask:             {PsAccessMask, "process desired access rights", kparams.AnsiString, []string{"ps.access.mask = '0x1400'"}, nil, nil},
	PsAccessMaskNames:        {PsAccessMaskNames, "process desired access rights as a string list", kparams.Slice, []string{"ps.access.mask.names in ('SUSPEND_RESUME')"}, nil, nil},
	PsAccessStatus:           {PsAccessStatus, "process access status", kparams.UnicodeString, []string{"ps.access.status = 'access is denied.'"}, nil, nil},
	PsSiblingPid:             {PsSiblingPid, "created or terminated process identifier", kparams.PID, []string{"ps.sibling.pid = 320"}, &Deprecation{Since: "1.10.0", Fields: []Field{PsChildPid}}, nil},
	PsChildPid:               {PsChildPid, "created or terminated process identifier", kparams.PID, []string{"ps.child.pid = 320"}, nil, nil},
	PsSiblingName:            {PsSiblingName, "created or terminated process name", kparams.UnicodeString, []string{"ps.sibling.name = 'notepad.exe'"}, &Deprecation{Since: "1.10.0", Fields: []Field{PsChildName}}, nil},
	PsChildName:              {PsChildName, "created or terminated process name", kparams.UnicodeString, []string{"ps.child.name = 'notepad.exe'"}, nil, nil},
	PsSiblingComm:            {PsSiblingComm, "created or terminated process command line", kparams.UnicodeString, []string{"ps.sibling.comm contains '\\k \\v'"}, &Deprecation{Since: "1.10.0", Fields: []Field{PsChildCmdline}}, nil},
	PsChildCmdline:           {PsChildCmdline, "created or terminated process command line", kparams.UnicodeString, []string{"ps.child.cmdline contains '\\k \\v'"}, nil, nil},
	PsSiblingArgs:            {PsSiblingArgs, "created process command line arguments", kparams.Slice, []string{"ps.sibling.args in ('/cdir', '/-C')"}, &Deprecation{Since: "1.10.0", Fields: []Field{PsChildArgs}}, nil},
	PsChildArgs:              {PsChildArgs, "created process command line arguments", kparams.Slice, []string{"ps.child.args in ('/cdir', '/-C')"}, nil, nil},
	PsSiblingExe:             {PsSiblingExe, "created, terminated, or opened process id", kparams.UnicodeString, []string{"ps.sibling.exe contains '\\Windows\\cmd.exe'"}, &Deprecation{Since: "1.10.0", Fields: []Field{PsChildExe}}, nil},
	PsChildExe:               {PsChildExe, "created, terminated, or opened process id", kparams.UnicodeString, []string{"ps.child.exe contains '\\Windows\\cmd.exe'"}, nil, nil},
	PsSiblingSID:             {PsSiblingSID, "created or terminated process security identifier", kparams.UnicodeString, []string{"ps.sibling.sid contains 'SERVICE'"}, &Deprecation{Since: "1.10.0", Fields: []Field{PsChildSID}}, nil},
	PsChildSID:               {PsChildSID, "created or terminated process security identifier", kparams.UnicodeString, []string{"ps.child.sid contains 'SERVICE'"}, nil, nil},
	PsSiblingSessionID:       {PsSiblingSessionID, "created or terminated process session identifier", kparams.Int16, []string{"ps.sibling.sessionid == 1"}, &Deprecation{Since: "1.10.0", Fields: []Field{PsChildSessionID}}, nil},
	PsChildSessionID:         {PsChildSessionID, "created or terminated process session identifier", kparams.Int16, []string{"ps.child.sessionid == 1"}, nil, nil},
	PsSiblingDomain:          {PsSiblingDomain, "created or terminated process domain", kparams.UnicodeString, []string{"ps.sibling.domain contains 'SERVICE'"}, &Deprecation{Since: "1.10.0", Fields: []Field{PsChildDomain}}, nil},
	PsChildDomain:            {PsChildDomain, "created or terminated process domain", kparams.UnicodeString, []string{"ps.child.domain contains 'SERVICE'"}, nil, nil},
	PsSiblingUsername:        {PsSiblingUsername, "created or terminated process username", kparams.UnicodeString, []string{"ps.sibling.username contains 'system'"}, &Deprecation{Since: "1.10.0", Fields: []Field{PsChildUsername}}, nil},
	PsChildUsername:          {PsChildUsername, "created or terminated process username", kparams.UnicodeString, []string{"ps.child.username contains 'system'"}, nil, nil},
	PsUUID:                   {PsUUID, "unique process identifier", kparams.Uint64, []string{"ps.uuid > 6000054355"}, nil, nil},
	PsParentUUID:             {PsParentUUID, "unique parent process identifier", kparams.Uint64, []string{"ps.parent.uuid > 6000054355"}, nil, nil},
	PsChildUUID:              {PsChildUUID, "unique child process identifier", kparams.Uint64, []string{"ps.child.uuid > 6000054355"}, nil, nil},
	PsChildPeFilename:        {PsChildPeFilename, "original file name of the child process executable supplied at compile-time", kparams.UnicodeString, []string{"ps.child.pe.file.name = 'NOTEPAD.EXE'"}, nil, nil},
	PsChildIsWOW64Field:      {PsChildIsWOW64Field, "indicates if the 32-bit child process is created in 64-bit Windows system", kparams.Bool, []string{"ps.child.is_wow64"}, nil, nil},
	PsChildIsPackagedField:   {PsChildIsPackagedField, "indicates if the child process is packaged with the MSIX technology", kparams.Bool, []string{"ps.child.is_packaged"}, nil, nil},
	PsChildIsProtectedField:  {PsChildIsProtectedField, "indicates if the child process is a protected process", kparams.Bool, []string{"ps.child.is_protected"}, nil, nil},
	PsIsWOW64Field:           {PsIsWOW64Field, "indicates if the process generating the event is a 32-bit process created in 64-bit Windows system", kparams.Bool, []string{"ps.is_wow64"}, nil, nil},
	PsIsPackagedField:        {PsIsPackagedField, "indicates if the process generating the event is packaged with the MSIX technology", kparams.Bool, []string{"ps.is_packaged"}, nil, nil},
	PsIsProtectedField:       {PsIsProtectedField, "indicates if the process generating the event is a protected process", kparams.Bool, []string{"ps.is_protected"}, nil, nil},
	PsParentIsWOW64Field:     {PsParentIsWOW64Field, "indicates if the parent process generating the event is a 32-bit process created in 64-bit Windows system", kparams.Bool, []string{"ps.parent.is_wow64"}, nil, nil},
	PsParentIsPackagedField:  {PsParentIsPackagedField, "indicates if the parent process generating the event is packaged with the MSIX technology", kparams.Bool, []string{"ps.parent.is_packaged"}, nil, nil},
	PsParentIsProtectedField: {PsParentIsProtectedField, "indicates if the the parent process generating the event is a protected process", kparams.Bool, []string{"ps.parent.is_protected"}, nil, nil},
	PsAncestor: {PsAncestor, "the process ancestor name", kparams.UnicodeString, []string{"ps.ancestor[1] = 'svchost.exe'", "ps.ancestor in ('winword.exe')"}, nil, &Argument{Optional: true, Pattern: "[0-9]+", ValidationFunc: func(s string) bool {
		for _, c := range s {
			if !unicode.IsNumber(c) {
				return false
			}
		}
		return true
	}}},

	ThreadBasePrio:                          {ThreadBasePrio, "scheduler priority of the thread", kparams.Int8, []string{"thread.prio = 5"}, nil, nil},
	ThreadIOPrio:                            {ThreadIOPrio, "I/O priority hint for scheduling I/O operations", kparams.Int8, []string{"thread.io.prio = 4"}, nil, nil},
	ThreadPagePrio:                          {ThreadPagePrio, "memory page priority hint for memory pages accessed by the thread", kparams.Int8, []string{"thread.page.prio = 12"}, nil, nil},
	ThreadKstackBase:                        {ThreadKstackBase, "base address of the thread's kernel space stack", kparams.Address, []string{"thread.kstack.base = 'a65d800000'"}, nil, nil},
	ThreadKstackLimit:                       {ThreadKstackLimit, "limit of the thread's kernel space stack", kparams.Address, []string{"thread.kstack.limit = 'a85d800000'"}, nil, nil},
	ThreadUstackBase:                        {ThreadUstackBase, "base address of the thread's user space stack", kparams.Address, []string{"thread.ustack.base = '7ffe0000'"}, nil, nil},
	ThreadUstackLimit:                       {ThreadUstackLimit, "limit of the thread's user space stack", kparams.Address, []string{"thread.ustack.limit = '8ffe0000'"}, nil, nil},
	ThreadEntrypoint:                        {ThreadEntrypoint, "starting address of the function to be executed by the thread", kparams.Address, []string{"thread.entrypoint = '7efe0000'"}, &Deprecation{Since: "2.3.0", Fields: []Field{ThreadStartAddress}}, nil},
	ThreadStartAddress:                      {ThreadStartAddress, "thread start address", kparams.Address, []string{"thread.start_address = '7efe0000'"}, nil, nil},
	ThreadPID:                               {ThreadPID, "the process identifier where the thread is created", kparams.Uint32, []string{"kevt.pid != thread.pid"}, nil, nil},
	ThreadTEB:                               {ThreadTEB, "the base address of the thread environment block", kparams.Address, []string{"thread.teb_address = '8f30893000'"}, nil, nil},
	ThreadAccessMask:                        {ThreadAccessMask, "thread desired access rights", kparams.AnsiString, []string{"thread.access.mask = '0x1fffff'"}, nil, nil},
	ThreadAccessMaskNames:                   {ThreadAccessMaskNames, "thread desired access rights as a string list", kparams.Slice, []string{"thread.access.mask.names in ('IMPERSONATE')"}, nil, nil},
	ThreadAccessStatus:                      {ThreadAccessStatus, "thread access status", kparams.UnicodeString, []string{"thread.access.status = 'success'"}, nil, nil},
	ThreadCallstackSummary:                  {ThreadCallstackSummary, "callstack summary", kparams.UnicodeString, []string{"thread.callstack.summary contains 'ntdll.dll|KERNELBASE.dll'"}, nil, nil},
	ThreadCallstackDetail:                   {ThreadCallstackDetail, "detailed information of each stack frame", kparams.UnicodeString, []string{"thread.callstack.detail contains 'KERNELBASE.dll!CreateProcessW'"}, nil, nil},
	ThreadCallstackModules:                  {ThreadCallstackModules, "list of modules comprising the callstack", kparams.Slice, []string{"thread.callstack.modules in ('C:\\WINDOWS\\System32\\KERNELBASE.dll')"}, nil, nil},
	ThreadCallstackSymbols:                  {ThreadCallstackSymbols, "list of symbols comprising the callstack", kparams.Slice, []string{"thread.callstack.symbols in ('ntdll.dll!NtCreateProcess')"}, nil, nil},
	ThreadCallstackAllocationSizes:          {ThreadCallstackAllocationSizes, "allocation sizes of private pages", kparams.Slice, []string{"thread.callstack.allocation_sizes > 10000"}, nil, nil},
	ThreadCallstackProtections:              {ThreadCallstackProtections, "page protections masks of each frame", kparams.Slice, []string{"thread.callstack.protections in ('RWX', 'WX')"}, nil, nil},
	ThreadCallstackCallsiteLeadingAssembly:  {ThreadCallstackCallsiteLeadingAssembly, "callsite leading assembly instructions", kparams.Slice, []string{"thread.callstack.callsite_leading_assembly in ('mov r10,rcx', 'syscall')"}, nil, nil},
	ThreadCallstackCallsiteTrailingAssembly: {ThreadCallstackCallsiteTrailingAssembly, "callsite trailing assembly instructions", kparams.Slice, []string{"thread.callstack.callsite_trailing_assembly in ('add esp, 0xab')"}, nil, nil},
	ThreadCallstackIsUnbacked:               {ThreadCallstackIsUnbacked, "indicates if the callstack contains unbacked regions", kparams.Bool, []string{"thread.callstack.is_unbacked"}, nil, nil},
	ThreadStartAddressSymbol:                {ThreadStartAddressSymbol, "thread start address symbol", kparams.UnicodeString, []string{"thread.start_address.symbol = 'LoadImage'"}, nil, nil},
	ThreadStartAddressModule:                {ThreadStartAddressModule, "thread start address module", kparams.UnicodeString, []string{"thread.start_address.module endswith 'kernel32.dll'"}, nil, nil},

	ImagePath:               {ImagePath, "full image path", kparams.UnicodeString, []string{"image.patj = 'C:\\Windows\\System32\\advapi32.dll'"}, nil, nil},
	ImageName:               {ImageName, "image name", kparams.UnicodeString, []string{"image.name = 'advapi32.dll'"}, nil, nil},
	ImageBase:               {ImageBase, "the base address of process in which the image is loaded", kparams.Address, []string{"image.base.address = 'a65d800000'"}, nil, nil},
	ImageChecksum:           {ImageChecksum, "image checksum", kparams.Uint32, []string{"image.checksum = 746424"}, nil, nil},
	ImageSize:               {ImageSize, "image size", kparams.Uint32, []string{"image.size > 1024"}, nil, nil},
	ImageDefaultAddress:     {ImageDefaultAddress, "default image address", kparams.Address, []string{"image.default.address = '7efe0000'"}, nil, nil},
	ImagePID:                {ImagePID, "target process identifier", kparams.Uint32, []string{"image.pid = 80"}, nil, nil},
	ImageSignatureType:      {ImageSignatureType, "image signature type", kparams.AnsiString, []string{"image.signature.type != 'NONE'"}, nil, nil},
	ImageSignatureLevel:     {ImageSignatureLevel, "image signature level", kparams.AnsiString, []string{"image.signature.level = 'AUTHENTICODE'"}, nil, nil},
	ImageCertSerial:         {ImageCertSerial, "image certificate serial number", kparams.UnicodeString, []string{"image.cert.serial = '330000023241fb59996dcc4dff000000000232'"}, nil, nil},
	ImageCertSubject:        {ImageCertSubject, "image certificate subject", kparams.UnicodeString, []string{"image.cert.subject contains 'Washington, Redmond, Microsoft Corporation'"}, nil, nil},
	ImageCertIssuer:         {ImageCertIssuer, "image certificate CA", kparams.UnicodeString, []string{"image.cert.issuer contains 'Washington, Redmond, Microsoft Corporation'"}, nil, nil},
	ImageCertAfter:          {ImageCertAfter, "image certificate expiration date", kparams.Time, []string{"image.cert.after contains '2024-02-01 00:05:42 +0000 UTC'"}, nil, nil},
	ImageCertBefore:         {ImageCertBefore, "image certificate enrollment date", kparams.Time, []string{"image.cert.before contains '2024-02-01 00:05:42 +0000 UTC'"}, nil, nil},
	ImageIsDriverMalicious:  {ImageIsDriverMalicious, "indicates if the loaded driver is malicious", kparams.Bool, []string{"image.is_driver_malicious"}, nil, nil},
	ImageIsDriverVulnerable: {ImageIsDriverVulnerable, "indicates if the loaded driver is vulnerable", kparams.Bool, []string{"image.is_driver_vulnerable"}, nil, nil},
	ImageIsDLL:              {ImageIsDLL, "indicates if the loaded image is a DLL", kparams.Bool, []string{"image.is_dll'"}, nil, nil},
	ImageIsDriver:           {ImageIsDriver, "indicates if the loaded image is a driver", kparams.Bool, []string{"image.is_driver'"}, nil, nil},
	ImageIsExecutable:       {ImageIsExecutable, "indicates if the loaded image is an executable", kparams.Bool, []string{"image.is_exec'"}, nil, nil},
	ImageIsDotnet:           {ImageIsDotnet, "indicates if the loaded image is a .NET assembly", kparams.Bool, []string{"image.is_dotnet'"}, nil, nil},

	FileObject:                      {FileObject, "file object address", kparams.Uint64, []string{"file.object = 18446738026482168384"}, nil, nil},
	FilePath:                        {FilePath, "full file path", kparams.UnicodeString, []string{"file.path = 'C:\\Windows\\System32'"}, nil, nil},
	FileName:                        {FileName, "full file name", kparams.UnicodeString, []string{"file.name contains 'mimikatz'"}, nil, nil},
	FileOperation:                   {FileOperation, "file operation", kparams.AnsiString, []string{"file.operation = 'open'"}, nil, nil},
	FileShareMask:                   {FileShareMask, "file share mask", kparams.AnsiString, []string{"file.share.mask = 'rw-'"}, nil, nil},
	FileIOSize:                      {FileIOSize, "file I/O size", kparams.Uint32, []string{"file.io.size > 512"}, nil, nil},
	FileOffset:                      {FileOffset, "file offset", kparams.Uint64, []string{"file.offset = 1024"}, nil, nil},
	FileType:                        {FileType, "file type", kparams.AnsiString, []string{"file.type = 'directory'"}, nil, nil},
	FileExtension:                   {FileExtension, "file extension", kparams.AnsiString, []string{"file.extension = '.dll'"}, nil, nil},
	FileAttributes:                  {FileAttributes, "file attributes", kparams.Slice, []string{"file.attributes in ('archive', 'hidden')"}, nil, nil},
	FileStatus:                      {FileStatus, "file operation status message", kparams.UnicodeString, []string{"file.status != 'success'"}, nil, nil},
	FileViewBase:                    {FileViewBase, "view base address", kparams.Address, []string{"file.view.base = '25d42170000'"}, nil, nil},
	FileViewSize:                    {FileViewSize, "size of the mapped view", kparams.Uint64, []string{"file.view.size > 1024"}, nil, nil},
	FileViewType:                    {FileViewType, "type of the mapped view section", kparams.Enum, []string{"file.view.type = 'IMAGE'"}, nil, nil},
	FileViewProtection:              {FileViewProtection, "protection rights of the section view", kparams.AnsiString, []string{"file.view.protection = 'READONLY'"}, nil, nil},
	FileIsDriverMalicious:           {FileIsDriverMalicious, "indicates if the dropped driver is malicious", kparams.Bool, []string{"file.is_driver_malicious"}, nil, nil},
	FileIsDriverVulnerable:          {FileIsDriverVulnerable, "indicates if the dropped driver is vulnerable", kparams.Bool, []string{"file.is_driver_vulnerable"}, nil, nil},
	FileIsDLL:                       {FileIsDLL, "indicates if the created file is a DLL", kparams.Bool, []string{"file.is_dll'"}, nil, nil},
	FileIsDriver:                    {FileIsDriver, "indicates if the created file is a driver", kparams.Bool, []string{"file.is_driver'"}, nil, nil},
	FileIsExecutable:                {FileIsExecutable, "indicates if the created file is an executable", kparams.Bool, []string{"file.is_exec'"}, nil, nil},
	FilePID:                         {FilePID, "denotes the process id performing file operation", kparams.PID, []string{"file.pid = 4"}, nil, nil},
	FileKey:                         {FileKey, "uniquely identifies the file object", kparams.Uint64, []string{"file.key = 12446738026482168384"}, nil, nil},
	FileInfoClass:                   {FileInfoClass, "identifies the file information class", kparams.Enum, []string{"file.info_class = 'Allocation'"}, nil, nil},
	FileInfoAllocationSize:          {FileInfoAllocationSize, "file allocation size", kparams.Uint64, []string{"file.info.allocation_size > 645400"}, nil, nil},
	FileInfoEOFSize:                 {FileInfoEOFSize, "file EOF size", kparams.Uint64, []string{"file.info.eof_size > 1000"}, nil, nil},
	FileInfoIsDispositionDeleteFile: {FileInfoIsDispositionDeleteFile, "indicates if the file is deleted when its handle is closed", kparams.Bool, []string{"file.info.is_disposition_file_delete = true"}, nil, nil},

	RegistryPath:      {RegistryPath, "fully qualified registry path", kparams.UnicodeString, []string{"registry.path = 'HKEY_LOCAL_MACHINE\\SYSTEM'"}, nil, nil},
	RegistryKeyName:   {RegistryKeyName, "registry key name", kparams.UnicodeString, []string{"registry.key.name = 'CurrentControlSet'"}, nil, nil},
	RegistryKeyHandle: {RegistryKeyHandle, "registry key object address", kparams.Address, []string{"registry.key.handle = 'FFFFB905D60C2268'"}, nil, nil},
	RegistryValue:     {RegistryValue, "registry value content", kparams.UnicodeString, []string{"registry.value = '%SystemRoot%\\system32'"}, nil, nil},
	RegistryValueType: {RegistryValueType, "type of registry value", kparams.UnicodeString, []string{"registry.value.type = 'REG_SZ'"}, nil, nil},
	RegistryStatus:    {RegistryStatus, "status of registry operation", kparams.UnicodeString, []string{"registry.status != 'success'"}, nil, nil},

	NetDIP:        {NetDIP, "destination IP address", kparams.IP, []string{"net.dip = 172.17.0.3"}, nil, nil},
	NetSIP:        {NetSIP, "source IP address", kparams.IP, []string{"net.sip = 127.0.0.1"}, nil, nil},
	NetDport:      {NetDport, "destination port", kparams.Uint16, []string{"net.dport in (80, 443, 8080)"}, nil, nil},
	NetSport:      {NetSport, "source port", kparams.Uint16, []string{"net.sport != 3306"}, nil, nil},
	NetDportName:  {NetDportName, "destination port name", kparams.AnsiString, []string{"net.dport.name = 'dns'"}, nil, nil},
	NetSportName:  {NetSportName, "source port name", kparams.AnsiString, []string{"net.sport.name = 'http'"}, nil, nil},
	NetL4Proto:    {NetL4Proto, "layer 4 protocol name", kparams.AnsiString, []string{"net.l4.proto = 'TCP"}, nil, nil},
	NetPacketSize: {NetPacketSize, "packet size", kparams.Uint32, []string{"net.size > 512"}, nil, nil},
	NetSIPNames:   {NetSIPNames, "source IP names", kparams.Slice, []string{"net.sip.names in ('github.com.')"}, nil, nil},
	NetDIPNames:   {NetDIPNames, "destination IP names", kparams.Slice, []string{"net.dip.names in ('github.com.')"}, nil, nil},

	HandleID:     {HandleID, "handle identifier", kparams.Uint16, []string{"handle.id = 24"}, nil, nil},
	HandleObject: {HandleObject, "handle object address", kparams.Address, []string{"handle.object = 'FFFFB905DBF61988'"}, nil, nil},
	HandleName:   {HandleName, "handle name", kparams.UnicodeString, []string{"handle.name = '\\Device\\NamedPipe\\chrome.12644.28.105826381'"}, nil, nil},
	HandleType:   {HandleType, "handle type", kparams.AnsiString, []string{"handle.type = 'Mutant'"}, nil, nil},

	PeNumSections: {PeNumSections, "number of sections", kparams.Uint16, []string{"pe.nsections < 5"}, nil, nil},
	PeNumSymbols:  {PeNumSymbols, "number of entries in the symbol table", kparams.Uint32, []string{"pe.nsymbols > 230"}, nil, nil},
	PeBaseAddress: {PeBaseAddress, "image base address", kparams.Address, []string{"pe.address.base = '140000000'"}, nil, nil},
	PeEntrypoint:  {PeEntrypoint, "address of the entrypoint function", kparams.Address, []string{"pe.address.entrypoint = '20110'"}, nil, nil},
	PeSymbols:     {PeSymbols, "imported symbols", kparams.Slice, []string{"pe.symbols in ('GetTextFaceW', 'GetProcessHeap')"}, nil, nil},
	PeImports:     {PeImports, "imported dynamic linked libraries", kparams.Slice, []string{"pe.imports in ('msvcrt.dll', 'GDI32.dll'"}, nil, nil},

	PeResources: {PeResources, "version resources", kparams.Map, []string{"pe.resources[FileDescription] = 'Notepad'"}, nil, &Argument{Optional: true, Pattern: "[a-zA-Z0-9_]+", ValidationFunc: func(s string) bool {
		for _, c := range s {
			switch {
			case unicode.IsLower(c):
			case unicode.IsUpper(c):
			case unicode.IsNumber(c):
			case c == '_':
			default:
				return false
			}
		}
		return true
	}}},

	PeCompany:         {PeCompany, "internal company name of the file provided at compile-time", kparams.UnicodeString, []string{"pe.company = 'Microsoft Corporation'"}, nil, nil},
	PeCopyright:       {PeCopyright, "copyright notice for the file emitted at compile-time", kparams.UnicodeString, []string{"pe.copyright = '© Microsoft Corporation'"}, nil, nil},
	PeDescription:     {PeDescription, "internal description of the file provided at compile-time", kparams.UnicodeString, []string{"pe.description = 'Notepad'"}, nil, nil},
	PeFileName:        {PeFileName, "original file name supplied at compile-time", kparams.UnicodeString, []string{"pe.file.name = 'NOTEPAD.EXE'"}, nil, nil},
	PeFileVersion:     {PeFileVersion, "file version supplied at compile-time", kparams.UnicodeString, []string{"pe.file.version = '10.0.18362.693 (WinBuild.160101.0800)'"}, nil, nil},
	PeProduct:         {PeProduct, "internal product name of the file provided at compile-time", kparams.UnicodeString, []string{"pe.product = 'Microsoft® Windows® Operating System'"}, nil, nil},
	PeProductVersion:  {PeProductVersion, "internal product version of the file provided at compile-time", kparams.UnicodeString, []string{"pe.product.version = '10.0.18362.693'"}, nil, nil},
	PeIsDLL:           {PeIsDLL, "indicates if the loaded image or created file is a DLL", kparams.Bool, []string{"pe.is_dll'"}, &Deprecation{Since: "2.0.0", Fields: []Field{FileIsDLL, ImageIsDLL}}, nil},
	PeIsDriver:        {PeIsDriver, "indicates if the loaded image or created file is a driver", kparams.Bool, []string{"pe.is_driver'"}, &Deprecation{Since: "2.0.0", Fields: []Field{FileIsDriver, ImageIsDriver}}, nil},
	PeIsExecutable:    {PeIsExecutable, "indicates if the loaded image or created file is an executable", kparams.Bool, []string{"pe.is_exec'"}, &Deprecation{Since: "2.0.0", Fields: []Field{FileIsExecutable, ImageIsExecutable}}, nil},
	PeImphash:         {PeImphash, "import hash", kparams.AnsiString, []string{"pe.impash = '5d3861c5c547f8a34e471ba273a732b2'"}, nil, nil},
	PeIsDotnet:        {PeIsDotnet, "indicates if PE contains CLR data", kparams.Bool, []string{"pe.is_dotnet"}, nil, nil},
	PeAnomalies:       {PeAnomalies, "contains PE anomalies detected during parsing", kparams.Slice, []string{"pe.anomalies in ('number of sections is 0')"}, nil, nil},
	PeIsSigned:        {PeIsSigned, "indicates if the PE has embedded or catalog signature", kparams.Bool, []string{"pe.is_signed"}, nil, nil},
	PeIsTrusted:       {PeIsTrusted, "indicates if the PE certificate chain is trusted", kparams.Bool, []string{"pe.is_trusted"}, nil, nil},
	PeCertSerial:      {PeCertSerial, "PE certificate serial number", kparams.UnicodeString, []string{"pe.cert.serial = '330000023241fb59996dcc4dff000000000232'"}, nil, nil},
	PeCertSubject:     {PeCertSubject, "PE certificate subject", kparams.UnicodeString, []string{"pe.cert.subject contains 'Washington, Redmond, Microsoft Corporation'"}, nil, nil},
	PeCertIssuer:      {PeCertIssuer, "PE certificate CA", kparams.UnicodeString, []string{"pe.cert.issuer contains 'Washington, Redmond, Microsoft Corporation'"}, nil, nil},
	PeCertAfter:       {PeCertAfter, "PE certificate expiration date", kparams.Time, []string{"pe.cert.after contains '2024-02-01 00:05:42 +0000 UTC'"}, nil, nil},
	PeCertBefore:      {PeCertBefore, "PE certificate enrollment date", kparams.Time, []string{"pe.cert.before contains '2024-02-01 00:05:42 +0000 UTC'"}, nil, nil},
	PeIsModified:      {PeIsModified, "indicates if disk and in-memory PE headers differ", kparams.Bool, []string{"pe.is_modified"}, nil, nil},
	PePsChildFileName: {PePsChildFileName, "original file name of the child process executable supplied at compile-time", kparams.UnicodeString, []string{"pe.ps.child.file.name = 'NOTEPAD.EXE'"}, &Deprecation{Since: "2.3.0", Fields: []Field{PsChildPeFilename}}, nil},

	MemBaseAddress:    {MemBaseAddress, "region base address", kparams.Address, []string{"mem.address = '211d13f2000'"}, nil, nil},
	MemRegionSize:     {MemRegionSize, "region size", kparams.Uint64, []string{"mem.size > 438272"}, nil, nil},
	MemAllocType:      {MemAllocType, "region allocation or release type", kparams.Flags, []string{"mem.alloc = 'COMMIT'"}, nil, nil},
	MemPageType:       {MemPageType, "page type of the allocated region", kparams.Enum, []string{"mem.type = 'PRIVATE'"}, nil, nil},
	MemProtection:     {MemProtection, "allocated region protection type", kparams.Enum, []string{"mem.protection = 'READWRITE'"}, nil, nil},
	MemProtectionMask: {MemProtectionMask, "allocated region protection in mask notation", kparams.Enum, []string{"mem.protection.mask = 'RWX'"}, nil, nil},

	DNSName:    {DNSName, "dns query name", kparams.UnicodeString, []string{"dns.name = 'example.org'"}, nil, nil},
	DNSRR:      {DNSRR, "dns resource record type", kparams.AnsiString, []string{"dns.rr = 'AA'"}, nil, nil},
	DNSOptions: {DNSOptions, "dns query options", kparams.Flags64, []string{"dns.options in ('ADDRCONFIG', 'DUAL_ADDR')"}, nil, nil},
	DNSRcode:   {DNSRR, "dns response status", kparams.AnsiString, []string{"dns.rcode = 'NXDOMAIN'"}, nil, nil},
	DNSAnswers: {DNSAnswers, "dns response answers", kparams.Slice, []string{"dns.answers in ('o.lencr.edgesuite.net', 'a1887.dscq.akamai.net')"}, nil, nil},
}

// ArgumentOf returns argument data for the specified field.
func ArgumentOf(name string) *Argument {
	f, ok := fields[Field(name)]
	if !ok {
		// this can happen for pseudo fields
		return nil
	}
	return f.Argument
}

// IsField returns true if the provided string is a
// recognized field or pseudo field. Otherwise, it
// returns false.
func IsField(name string) bool {
	if _, ok := fields[Field(name)]; ok || IsPseudoField(Field(name)) {
		return true
	}
	return false
}
