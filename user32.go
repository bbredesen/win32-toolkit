package win32

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

func CreateWindowExW(exStyle uint32,
	className, windowName string,
	style uint32,
	x, y, width, height uint32,
	parent, menu, instance HInstance,
	pParam unsafe.Pointer,
) (HWnd, Win32Error) {

	rval, _, err := procs["CreateWindowExW"].Call(
		uintptr(exStyle),
		uintptr(unsafe.Pointer(windows.StringToUTF16Ptr(className))),
		uintptr(unsafe.Pointer(windows.StringToUTF16Ptr(windowName))),
		uintptr(style),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(parent),
		uintptr(menu),
		uintptr(instance),
		uintptr(pParam),
	)

	if rval == 0 {
		errCode := GetLastError()
		if err != nil {
			fmt.Println("OOPS")
			fmt.Println(err)
		}
		return HWnd(0), Win32Error(errCode)
	} else {
		return HWnd(rval), Win32Error(0)
	}
}

func DefWindowProcW(hwnd HWnd, msg Msg, wParam, lParam uintptr) uintptr {
	rval, _, _ := procs["DefWindowProcW"].Call(
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam,
	)
	return rval
}

func DestroyWindow(hwnd HWnd) Win32Error {
	result, _, err := procs["DestroyWindow"].Call(
		uintptr(hwnd),
	)
	if int32(result) == 0 {
		errCode := GetLastError()
		if err != nil {
			fmt.Println("OOPS")
			fmt.Println(err)
		}
		return Win32Error(errCode)
	} else {
		return Win32Error(0)
	}

}

func DispatchMessage(pMsg *MSG) int32 {
	return DispatchMessageW(pMsg)
}

func DispatchMessageW(pMsg *MSG) int32 {
	result, _, _ := procs["DispatchMessageW"].Call(
		uintptr(unsafe.Pointer(pMsg)),
	)

	return int32(result)
}

func GetClientRect(hwnd HWnd, pRect *Rect) Win32Error {
	success, _, _ := procs["GetClientRect"].Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(pRect)),
	)
	if success == 0 {
		errCode := GetLastError()

		return Win32Error(errCode)
	} else {
		return Win32Error(0)
	}
}

func GetMessageW(pMsg *MSG, hwnd HWnd, msgFilterMin, msgFilterMax uint32) (int32, Win32Error) {
	// localMsg := MSG{}
	success, _, err := procs["GetMessageW"].Call(
		uintptr(unsafe.Pointer(pMsg)), //&localMsg)),
		uintptr(hwnd),
		uintptr(msgFilterMin),
		uintptr(msgFilterMax),
	)
	if int32(success) < 0 {
		errCode := GetLastError()
		if err != nil {
			fmt.Println("OOPS")
			fmt.Println(err)
		}
		return int32(success), Win32Error(errCode)
	} else {
		return int32(success), Win32Error(0)
	}
}

func LoadCursor(hInstance HInstance, cursorName uintptr) (HCursor, Win32Error) {
	rval, _, _ := procs["LoadCursorW"].Call(
		uintptr(hInstance),
		cursorName,
	)
	if rval == 0 {
		errCode := GetLastError()
		return HCursor(0), Win32Error(errCode)
	}
	return HCursor(rval), Win32Error(0)

}

func PeekMessageW(pMsg *MSG, hwnd HWnd, msgFilterMin, msgFilterMax, removeMessage uint32) (int32, Win32Error) {
	success, _, err := procs["PeekMessageW"].Call(
		uintptr(unsafe.Pointer(pMsg)),
		uintptr(hwnd),
		uintptr(msgFilterMin),
		uintptr(msgFilterMax),
		uintptr(removeMessage),
	)
	if int32(success) < 0 {
		errCode := GetLastError()
		if err != nil {
			fmt.Println("OOPS")
			fmt.Println(err)
		}
		return int32(success), Win32Error(errCode)
	} else {
		return int32(success), Win32Error(0)
	}
}

func PostQuitMessage(exitCode int32) {
	procs["PostQuitMessage"].Call(
		uintptr(exitCode),
	)
}

func RegisterClassExW(wc *WNDCLASSEXW) (uint16, Win32Error) {
	rval, _, _ := procs["RegisterClassExW"].Call(
		uintptr(unsafe.Pointer(wc)),
	)

	if rval == 0 {
		errCode := GetLastError()
		return uint16(0), Win32Error(errCode)
	}
	return uint16(rval), Win32Error(0)

}

func TranslateMessage(pMsg *MSG) int32 {
	result, _, _ := procs["TranslateMessage"].Call(
		uintptr(unsafe.Pointer(pMsg)),
	)
	return int32(result)
}

func ValidateRect(hwnd HWnd, pRect *Rect) Win32Error {
	success, _, err := procs["ValidateRect"].Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(pRect)),
	)
	if success == 0 {
		errCode := GetLastError()
		if err != nil {
			fmt.Println("OOPS")
			fmt.Println(err)
		}
		return Win32Error(errCode)
	} else {
		return Win32Error(0)
	}
}
