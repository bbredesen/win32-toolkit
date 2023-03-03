package win32

import "golang.org/x/sys/windows"

type Handle windows.Handle

type HBrush Handle
type HCursor Handle
type HIcon Handle
type HInstance Handle
type HModule HInstance
type HWnd Handle

const NULL_HANDLE Handle = Handle(0)

type Rect struct {
	Left, Top, Right, Bottom int32
}

const (
	IDC_ARROW uintptr = 32512
)

const (
	CS_OWNDC   uint32 = 0x0020
	CS_VREDRAW uint32 = 0x0001
	CS_HREDRAW uint32 = 0x0002

	SW_SHOW        uint32 = 5
	SW_USE_DEFAULT uint32 = 0x80000000
)

const (
	WS_MAXIMIZE_BOX uint32 = 0x00010000
	WS_MINIMIZEBOX  uint32 = 0x00020000
	WS_THICKFRAME   uint32 = 0x00040000
	WS_SYSMENU      uint32 = 0x00080000
	WS_CAPTION      uint32 = 0x00C00000
	WS_VISIBLE      uint32 = 0x10000000

	WS_OVERLAPPEDWINDOW uint32 = 0x00CF0000
)

const (
	PM_NOREMOVE uint32 = 0x0000
	PM_REMOVE   uint32 = 0x0001
	PM_NOYIELD  uint32 = 0x0002
)

const (
	COLOR_WINDOW uintptr = 5
)

type WNDCLASSEXW struct {
	Size       uint32
	Style      uint32
	WndProc    uintptr
	ClsExtra   int32
	WndExtra   int32
	Instance   HInstance
	Icon       HIcon
	Cursor     HCursor
	Background HBrush
	MenuName   *uint16
	ClassName  *uint16
	IconSm     HIcon
}

type MSG struct {
	Hwnd           HWnd
	Message        uint32
	WParam, LParam uintptr
	Time           uint32
	Pt             POINT
	lPrivate       uint32
}

type POINT struct {
	x, y uint32
}
