//go:build windows
// +build windows

package kernel32

import (
	"fmt"
	"math"
	"unsafe"

	"github.com/0xrawsec/golang-win32/win32"
)

const (
	STANDARD_RIGHTS_REQUIRED = 0x000F0000
	DELETE                   = 0x00010000
	READ_CONTROL             = 0x00020000
	WRITE_DAC                = 0x00040000
	WRITE_OWNER              = 0x00080000
	SYNCHRONIZE              = 0x00100000

	STANDARD_RIGHTS_READ    = READ_CONTROL
	STANDARD_RIGHTS_WRITE   = READ_CONTROL
	STANDARD_RIGHTS_EXECUTE = READ_CONTROL

	STANDARD_RIGHTS_ALL = 0x001F0000

	SPECIFIC_RIGHTS_ALL = 0x0000FFFF

	ACCESS_SYSTEM_SECURITY = 0x01000000
	MAXIMUM_ALLOWED        = 0x02000000

	GENERIC_READ    = 0x80000000
	GENERIC_WRITE   = 0x40000000
	GENERIC_EXECUTE = 0x20000000
	GENERIC_ALL     = 0x10000000

	PROCESS_ALL_ACCESS                = STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | 0xffff
	PROCESS_TERMINATE                 = 0x0001
	PROCESS_CREATE_THREAD             = 0x0002
	PROCESS_SET_SESSIONID             = 0x0004
	PROCESS_VM_OPERATION              = 0x0008
	PROCESS_VM_READ                   = 0x0010
	PROCESS_VM_WRITE                  = 0x0020
	PROCESS_DUP_HANDLE                = 0x0040
	PROCESS_CREATE_PROCESS            = 0x0080
	PROCESS_SET_QUOTA                 = 0x0100
	PROCESS_SET_INFORMATION           = 0x0200
	PROCESS_QUERY_INFORMATION         = 0x0400
	PROCESS_SUSPEND_RESUME            = 0x0800
	PROCESS_QUERY_LIMITED_INFORMATION = 0x1000
)

// threads
const (
	THREAD_TERMINATE                 = 0x0001
	THREAD_SUSPEND_RESUME            = 0x0002
	THREAD_GET_CONTEXT               = 0x0008
	THREAD_SET_CONTEXT               = 0x0010
	THREAD_SET_INFORMATION           = 0x0020
	THREAD_QUERY_INFORMATION         = 0x0040
	THREAD_SET_THREAD_TOKEN          = 0x0080
	THREAD_IMPERSONATE               = 0x0100
	THREAD_DIRECT_IMPERSONATION      = 0x0200
	THREAD_SET_LIMITED_INFORMATION   = 0x0400
	THREAD_QUERY_LIMITED_INFORMATION = 0x0800
	THREAD_ALL_ACCESS                = STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | 0xffff
)

// thelp32.h

type LPTHREADENTRY32 *THREADENTRY32
type LPCTHREADENTRY32 *THREADENTRY32

const (
	TH32CS_SNAPHEAPLIST = 0x00000001
	TH32CS_SNAPPROCESS  = 0x00000002
	TH32CS_SNAPTHREAD   = 0x00000004
	TH32CS_SNAPMODULE   = 0x00000008
	TH32CS_SNAPMODULE32 = 0x00000010
	TH32CS_SNAPALL      = TH32CS_SNAPHEAPLIST | TH32CS_SNAPPROCESS | TH32CS_SNAPTHREAD | TH32CS_SNAPMODULE
	TH32CS_INHERIT      = 0x80000000
)

type THREADENTRY32 struct {
	DwSize             win32.DWORD
	CntUsage           win32.DWORD
	Th32ThreadID       win32.DWORD
	Th32OwnerProcessID win32.DWORD
	TpBasePri          win32.LONG
	TpDeltaPri         win32.LONG
	DwFlags            win32.DWORD
}

func NewThreadEntry32() THREADENTRY32 {
	te := THREADENTRY32{}
	te.DwSize = win32.DWORD(unsafe.Sizeof(te))
	return te
}

type PROCESSENTRY32W struct {
	DwSize              win32.DWORD
	CntUsage            win32.DWORD
	Th32ProcessID       win32.DWORD
	Th32DefaultHeapID   win32.ULONG_PTR
	Th32ModuleID        win32.DWORD
	CntThreads          win32.DWORD
	Th32ParentProcessID win32.DWORD
	PcPriClassBase      win32.LONG
	DwFlags             win32.DWORD
	SzExeFile           [win32.MAX_PATH]uint16
}

type LPPROCESSENTRY32W *PROCESSENTRY32W
type LPCPROCESSENTRY32W *PROCESSENTRY32W

func NewProcessEntry32W() PROCESSENTRY32W {
	return PROCESSENTRY32W{DwSize: win32.DWORD(unsafe.Sizeof(PROCESSENTRY32W{}))}
}

type MODULEENTRY32 struct {
	DwSize        win32.DWORD
	Th32ModuleID  uint32
	Th32ProcessID uint32
	GlblcntUsage  uint32
	ProccntUsage  uint32
	ModBaseAddr   uintptr
	ModBaseSize   uint32
	Module        win32.HANDLE
	SzModule      [win32.MAX_MODULE_NAME32 + 1]uint16
	SzExePath     [win32.MAX_PATH]uint16
}

type LPMODULEENTRY32 *MODULEENTRY32
type LPCMODULEENTRY32 *MODULEENTRY32

func NewModuleEntry32() MODULEENTRY32 {
	entry := MODULEENTRY32{}
	entry.DwSize = win32.DWORD(unsafe.Sizeof(entry))
	return entry
}

type MODULEINFO struct {
	LpBaseOfDll win32.LPVOID
	// Size of the image mapped in memory
	// To compute it from the image file we need to add all section sizes
	// rounded up to the dwPageSize (minimum alloc size) + 1 page for the PE header
	SizeOfImage win32.DWORD
	EntryPoint  win32.LPVOID
}

func (mi MODULEINFO) String() string {
	return fmt.Sprintf("LpBaseOfDll: 0x%016x SizeOfImage: %d Entrypoint: 0x%016x Entrypoint (relative to base): 0x%08x", mi.LpBaseOfDll, mi.SizeOfImage, mi.EntryPoint,
		mi.EntryPoint-mi.LpBaseOfDll)
}

type ProcessInformationClass uint32

const (
	ProcessMemoryPriorityClass = ProcessInformationClass(iota)
	ProcessMemoryExhaustionInfoClass
	ProcessAppMemoryInfoClass
	ProcessInPrivateInfoClass
	ProcessPowerThrottlingClass
	ProcessReservedValue1Class
	ProcessTelemetryCoverageInfoClass
	ProcessProtectionLevelInfoClass
	ProcessLeapSecondInfoClass
	ProcessMachineTypeInfoClass
	ProcessInformationClassMaxClass
)

type ProcessProtectionLevelInformation uint32

const (
	// values assumed from https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/ns-processthreadsapi-process_protection_level_information
	// partially confirmed from: https://github.com/processhacker/processhacker/blob/master/phnt/include/ntpsapi.h
	ProtectionLevelWintcbLight = ProcessProtectionLevelInformation(iota)
	ProtectionLevelWindows
	ProtectionLevelWindowsLight
	ProtectionLevelAntimalwareLight
	ProtectionLevelLsaLight
	ProtectionLevelWintcb
	ProtectionLevelCodegenLight
	ProtectionLevelAuthenticode
	ProtectionLevelPplApp

	ProtectionLevelNone = math.MaxUint32 - 1
)

type MemoryPriorityInformation uint32

const (
	MemoryPriorityVeryLow = MemoryPriorityInformation(iota + 1)
	MemoryPriorityLow
	MemoryPriorityMedium
	MemoryPriorityBelowNormal
	MemoryPriorityNormal
)

type ProcessMemoryExhaustionType uint32

const (
	PMETypeFailFastOnCommitFailure = ProcessMemoryExhaustionType(iota)
	PMETypeMax
)

type ProcessMemoryExhaustionInfo struct {
	Version  uint16
	Reserved uint16
	Type     ProcessMemoryExhaustionType
	Value    uintptr
}

type AppMemoryInformation struct {
	AvailableCommit        uint64
	PrivateCommitUsage     uint64
	PeakPrivateCommitUsage uint64
	TotalCommitUsage       uint64
}
