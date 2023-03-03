package win32

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

func GetLastError() uint32 {
	rval, _, _ := procs["GetLastError"].Call()
	return uint32(rval)
}

func GetModuleHandleExW(flags uint32, moduleName string) (HInstance, Win32Error) {
	var rval HInstance
	var lpModuleName unsafe.Pointer = nil
	if moduleName != "" {
		lpModuleName = unsafe.Pointer(windows.StringToUTF16Ptr(moduleName))
	}

	success, _, err := procs["GetModuleHandleExW"].Call(
		uintptr(flags),
		uintptr(lpModuleName),
		uintptr(unsafe.Pointer(&rval)),
	)
	if success == 0 {
		errCode := GetLastError()
		if err != nil {
			fmt.Println("OOPS")
			fmt.Println(err)
		}
		return HInstance(0), Win32Error(errCode)
	} else {
		return rval, Win32Error(0)
	}
}
