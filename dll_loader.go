package win32

import (
	"golang.org/x/sys/windows"
)

var procs = make(map[string]*windows.LazyProc)

func fn(dll *windows.LazyDLL, name string) {
	procs[name] = dll.NewProc(name)
}

func init() {
	loadKernel32()
	loadUser32()
	loadGdi32()
}

func loadKernel32() {
	dll := windows.NewLazySystemDLL("kernel32.dll")
	fn(dll, "GetLastError")
	fn(dll, "GetModuleHandleExW")
}

func loadUser32() {
	dll := windows.NewLazySystemDLL("user32.dll")
	fn(dll, "CreateWindowExW")
	fn(dll, "DefWindowProcW")
	fn(dll, "DestroyWindow")
	fn(dll, "DispatchMessageW")
	fn(dll, "GetClientRect")
	fn(dll, "GetMessageW")
	fn(dll, "PeekMessageW")
	fn(dll, "LoadCursorW")
	fn(dll, "PostQuitMessage")
	fn(dll, "RegisterClassExW")
	fn(dll, "TranslateMessage")
	fn(dll, "ValidateRect")
	fn(dll, "GetClientRect")
}

func loadGdi32() {
	dll := windows.NewLazySystemDLL("gdi32.dll")
	fn(dll, "GetStockObject")
}
