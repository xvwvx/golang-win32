// +build windows

package win32

var (
	NULL  = uintptr(0)
	FALSE = BOOL(0)
	TRUE  = BOOL(1)
)

type BOOL int
type BOOLEAN BYTE
type BYTE byte
type CCHAR byte
type CHAR byte
type COLORREF DWORD
type DWORDLONG uint64
type DWORD uint32
type DWORD32 uint32
type DWORD64 uint64
type HANDLE PVOID
type HWND HANDLE
type WORD uint16
type SHORT int16
type USHORT uint16
type LONG int32
type ULONG uint32
type LONGLONG int64
type ULONGLONG uint64
type SIZE_T ULONG_PTR
type ULONG_PTR uintptr
type LONG_PTR int
type PVOID uintptr
type LPVOID uintptr
type LPCVOID uintptr
type LVOID uintptr
type NTSTATUS LONG
type KAFFINITY ULONG_PTR
type KPRIORITY LONG
type PPEB uintptr //not sure
type UINT uint
type WCHAR uint16
type PWSTR *WCHAR
type ACCESS_MASK ULONG

//
const (
	INVALID_HANDLE = LONG_PTR(-1)
)

const (
	MAX_PATH = 260
	MAX_MODULE_NAME32 = 255
)

// winbase.h
const (
	DEBUG_PROCESS                    = 0x1
	DEBUG_ONLY_THIS_PROCESS          = 0x2
	CREATE_SUSPENDED                 = 0x4
	DETACHED_PROCESS                 = 0x8
	CREATE_NEW_CONSOLE               = 0x10
	NORMAL_PRIORITY_CLASS            = 0x20
	IDLE_PRIORITY_CLASS              = 0x40
	HIGH_PRIORITY_CLASS              = 0x80
	REALTIME_PRIORITY_CLASS          = 0x100
	CREATE_NEW_PROCESS_GROUP         = 0x200
	CREATE_UNICODE_ENVIRONMENT       = 0x400
	CREATE_SEPARATE_WOW_VDM          = 0x800
	CREATE_SHARED_WOW_VDM            = 0x1000
	CREATE_FORCEDOS                  = 0x2000
	BELOW_NORMAL_PRIORITY_CLASS      = 0x4000
	ABOVE_NORMAL_PRIORITY_CLASS      = 0x8000
	INHERIT_PARENT_AFFINITY          = 0x10000
	INHERIT_CALLER_PRIORITY          = 0x20000
	CREATE_PROTECTED_PROCESS         = 0x40000
	EXTENDED_STARTUPINFO_PRESENT     = 0x80000
	PROCESS_MODE_BACKGROUND_BEGIN    = 0x100000
	PROCESS_MODE_BACKGROUND_END      = 0x200000
	CREATE_BREAKAWAY_FROM_JOB        = 0x1000000
	CREATE_PRESERVE_CODE_AUTHZ_LEVEL = 0x2000000
	CREATE_DEFAULT_ERROR_MODE        = 0x4000000
	CREATE_NO_WINDOW                 = 0x8000000
	PROFILE_USER                     = 0x10000000
	PROFILE_KERNEL                   = 0x20000000
	PROFILE_SERVER                   = 0x40000000
	CREATE_IGNORE_SYSTEM_DEFAULT     = 0x80000000

	// Thread Priorities
	THREAD_PRIORITY_LOWEST        = THREAD_BASE_PRIORITY_MIN
	THREAD_PRIORITY_BELOW_NORMAL  = (THREAD_PRIORITY_LOWEST + 1)
	THREAD_PRIORITY_NORMAL        = 0
	THREAD_PRIORITY_HIGHEST       = THREAD_BASE_PRIORITY_MAX
	THREAD_PRIORITY_ABOVE_NORMAL  = (THREAD_PRIORITY_HIGHEST - 1)
	THREAD_PRIORITY_ERROR_RETURN  = (MAXLONG)
	THREAD_PRIORITY_TIME_CRITICAL = THREAD_BASE_PRIORITY_LOWRT
	THREAD_PRIORITY_IDLE          = THREAD_BASE_PRIORITY_IDLE
)

//Process Access Rights
//https://msdn.microsoft.com/en-us/library/windows/desktop/ms684880(v=vs.85).aspx
const (
	PROCESS_CREATE_PROCESS            = 0x0080  //Required to create a process.
	PROCESS_CREATE_THREAD             = 0x0002  //Required to create a thread.
	PROCESS_DUP_HANDLE                = 0x0040  //Required to duplicate a handle using DuplicateHandle.
	PROCESS_QUERY_INFORMATION         = 0x0400  //Required to retrieve certain information about a process, such as its token, exit code, and priority class (see OpenProcessToken).
	PROCESS_QUERY_LIMITED_INFORMATION = 0x1000  //Required to retrieve certain information about a process (see GetExitCodeProcess, GetPriorityClass, IsProcessInJob, QueryFullProcessImageName). A handle that has the PROCESS_QUERY_INFORMATION access right is automatically granted
	PROCESS_SET_INFORMATION           = 0x0200  //Required to set certain information about a process, such as its priority class (see SetPriorityClass).
	PROCESS_SET_QUOTA                 = 0x0100  //Required to set memory limits using SetProcessWorkingSetSize.
	PROCESS_SUSPEND_RESUME            = 0x0800  //Required to suspend or resume a process.
	PROCESS_TERMINATE                 = 0x0001  //Required to terminate a process using TerminateProcess.
	PROCESS_VM_OPERATION              = 0x0008  //Required to perform an operation on the address space of a process (see VirtualProtectEx and WriteProcessMemory).
	PROCESS_VM_READ                   = 0x0010  //Required to read memory in a process using ReadProcessMemory.
	PROCESS_VM_WRITE                  = 0x0020  //Required to write to memory in a process using WriteProcessMemory.
	PROCESS_ALL_ACCESS                = 2035711 //This is not recommended.
	//SYNCHRONIZE                       = 0x00100000
)

// Memory Allocation Types
const (
	MEM_COMMIT      = 0x1000
	MEM_RESERVE     = 0x2000
	MEM_DECOMMIT    = 0x4000
	MEM_RELEASE     = 0x8000
	MEM_FREE        = 0x10000
	MEM_PRIVATE     = 0x20000
	MEM_MAPPED      = 0x40000
	MEM_RESET       = 0x80000
	MEM_TOP_DOWN    = 0x100000
	MEM_WRITE_WATCH = 0x200000
	MEM_PHYSICAL    = 0x400000
	MEM_ROTATE      = 0x800000
	MEM_LARGE_PAGES = 0x20000000
	MEM_4MB_PAGES   = 0x80000000
)

// Memory Protections
const (
	PAGE_NOACCESS          = 0x01
	PAGE_READONLY          = 0x02
	PAGE_READWRITE         = 0x04
	PAGE_WRITECOPY         = 0x08
	PAGE_EXECUTE           = 0x10
	PAGE_EXECUTE_READ      = 0x20
	PAGE_EXECUTE_READWRITE = 0x40
	PAGE_EXECUTE_WRITECOPY = 0x80
	PAGE_GUARD             = 0x100
	PAGE_NOCACHE           = 0x200
	PAGE_WRITECOMBINE      = 0x400
)

const (
	IGNORE   = 0
	INFINITE = 0xffffffff
)

// Winerror.h
const (
	//ERROR_NO_MORE_ITEMS     = 259
	//ERROR_INVALID_OPERATION = 4317

	WAIT_ABANDONED       = 0x80
	WAIT_OBJECT_0        = 0x0
	WAIT_TIMEOUT         = 0x102
	WAIT_FAILED          = 0xFFFFFFFF
	MAXIMUM_WAIT_OBJECTS = 64
)

// ntstatus.h
const (
	STATUS_SUCCESS = 0x0
	STATUS_PENDING = 0x00000103
)

// minwinbase.h
const (
	STILL_ACTIVE = STATUS_PENDING
)

const (
	SERVICE_KERNEL_DRIVER      = 0x00000001
	SERVICE_FILE_SYSTEM_DRIVER = 0x00000002
	SERVICE_ADAPTER            = 0x00000004
	SERVICE_RECOGNIZER_DRIVER  = 0x00000008

	SERVICE_DRIVER = (SERVICE_KERNEL_DRIVER |
		SERVICE_FILE_SYSTEM_DRIVER |
		SERVICE_RECOGNIZER_DRIVER)

	SERVICE_WIN32_OWN_PROCESS   = 0x00000010
	SERVICE_WIN32_SHARE_PROCESS = 0x00000020
	SERVICE_WIN32               = (SERVICE_WIN32_OWN_PROCESS |
		SERVICE_WIN32_SHARE_PROCESS)

	SERVICE_INTERACTIVE_PROCESS = 0x00000100

	SERVICE_TYPE_ALL = (SERVICE_WIN32 |
		SERVICE_ADAPTER |
		SERVICE_DRIVER |
		SERVICE_INTERACTIVE_PROCESS)
)

const (
	DELETE                   = 0x00010000
	READ_CONTROL             = 0x00020000
	WRITE_DAC                = 0x00040000
	WRITE_OWNER              = 0x00080000
	SYNCHRONIZE              = 0x00100000
	STANDARD_RIGHTS_REQUIRED = 0x000F0000
	STANDARD_RIGHTS_READ     = READ_CONTROL
	STANDARD_RIGHTS_WRITE    = READ_CONTROL
	STANDARD_RIGHTS_EXECUTE  = READ_CONTROL
	STANDARD_RIGHTS_ALL      = 0x001F0000
	SPECIFIC_RIGHTS_ALL      = 0x0000FFFF
	ACCESS_SYSTEM_SECURITY   = 0x01000000
	MAXIMUM_ALLOWED          = 0x02000000
	GENERIC_READ             = 0x80000000
	GENERIC_WRITE            = 0x40000000
	GENERIC_EXECUTE          = 0x20000000
	GENERIC_ALL              = 0x10000000
)

const (
	KEY_QUERY_VALUE        = 0x0001
	KEY_SET_VALUE          = 0x0002
	KEY_CREATE_SUB_KEY     = 0x0004
	KEY_ENUMERATE_SUB_KEYS = 0x0008
	KEY_NOTIFY             = 0x0010
	KEY_CREATE_LINK        = 0x0020
	KEY_WOW64_64KEY        = 0x0100
	KEY_WOW64_32KEY        = 0x0200
	KEY_WOW64_RES          = 0x0300
	KEY_READ               = (STANDARD_RIGHTS_READ | KEY_QUERY_VALUE | KEY_ENUMERATE_SUB_KEYS | KEY_NOTIFY&(^SYNCHRONIZE))
	KEY_WRITE              = (STANDARD_RIGHTS_WRITE | KEY_SET_VALUE | KEY_CREATE_SUB_KEY&(^SYNCHRONIZE))
	KEY_EXECUTE            = (KEY_READ & (^SYNCHRONIZE))
	KEY_ALL_ACCESS         = (STANDARD_RIGHTS_ALL | KEY_QUERY_VALUE | KEY_SET_VALUE | KEY_CREATE_SUB_KEY | KEY_ENUMERATE_SUB_KEYS | KEY_NOTIFY | KEY_CREATE_LINK&(^SYNCHRONIZE))
)

// Winnt.h
const (
	// Thread base priorities
	THREAD_BASE_PRIORITY_LOWRT = 15
	THREAD_BASE_PRIORITY_MAX   = 2
	THREAD_BASE_PRIORITY_MIN   = -2
	THREAD_BASE_PRIORITY_IDLE  = -15
)

// ntdef.h
const (
	MAXLONG = 0x7fffffff
)
