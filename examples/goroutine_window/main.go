package main

import (
	"fmt"
	"os"
	"runtime"
	"unsafe"

	win32 "github.com/bbredesen/win32-toolkit"
	"golang.org/x/sys/windows"
)

// Demonstrates a very basic window and message loop using the toolkit.

func main() {
	quitChan := make(chan int)
	go createAndLoop(quitChan)

	rval := <-quitChan

	fmt.Println("Message loop ended, exiting cleanly")
	os.Exit(rval)
}

// The thread that creates the window also has to run the message loop. You
// can't createWindow and then go messageLoop, or the window simply freezes
func createAndLoop(quitChan chan int) {
	// This ensures that the spawned goroutine is 1-to-1 with the current
	// thread.
	runtime.LockOSThread()

	_, hWnd := createWindow()
	quitChan <- messageLoop(hWnd)
}

func createWindow() (win32.HInstance, win32.HWnd) {
	className := "testClass"

	hInstance, err := win32.GetModuleHandleExW(0, "")
	if err != 0 {
		panic(err)
	}

	cursor, err := win32.LoadCursor(0, win32.IDC_ARROW)

	fn := wndProc

	wndClass := win32.WNDCLASSEXW{
		Style:      (win32.CS_OWNDC | win32.CS_HREDRAW | win32.CS_VREDRAW),
		WndProc:    windows.NewCallback(fn),
		Instance:   hInstance,
		Cursor:     cursor,
		Background: win32.HBrush(win32.COLOR_WINDOW + 1),
		ClassName:  windows.StringToUTF16Ptr(className),
	}

	wndClass.Size = uint32(unsafe.Sizeof(wndClass))

	if _, err := win32.RegisterClassExW(&wndClass); err != 0 {
		panic(err)
	}

	hWnd, err := win32.CreateWindowExW(
		0,
		className,
		"Test Window Title",
		(win32.WS_VISIBLE | win32.WS_OVERLAPPEDWINDOW),
		win32.SW_USE_DEFAULT,
		win32.SW_USE_DEFAULT,
		800,
		600,
		0,
		0,
		hInstance,
		nil,
	)

	if err != 0 {
		panic(err)
	}

	fmt.Printf("Instance: 0x%.8x\n", hInstance)
	fmt.Printf("  Window: 0x%.8x\n", hWnd)
	return hInstance, hWnd
}

func messageLoop(hWnd win32.HWnd) int {
	var msg = win32.MSG{} //win32.HWnd(0)
	for {
		code, _ := win32.GetMessageW(&msg, win32.HWnd(0), 0, 0)
		if code == 0 {
			break
		}

		win32.TranslateMessage(&msg)
		win32.DispatchMessage(&msg)
	}
	return 0
}

func wndProc(hwnd win32.HWnd, msg win32.Msg, wParam, lParam uintptr) uintptr {
	switch msg {
	case win32.WM_CREATE:
		fmt.Printf("Message: CREATE\n")
		break

	case win32.WM_PAINT:
		fmt.Println("Message: PAINT")
		win32.ValidateRect(hwnd, nil)
		break

	case win32.WM_MOUSEMOVE:
		fmt.Printf("Message MOUSEMOVE at %d, %d\n", lParam&0xFFFF, lParam>>16)

	case win32.WM_CLOSE:
		fmt.Println("Message: CLOSE")
		win32.DestroyWindow(hwnd)
		break

	case win32.WM_DESTROY:
		fmt.Println("Message: DESTROY")
		win32.PostQuitMessage(0)
		break

	default:
		// fmt.Printf("%s\n", msg.String())
		return win32.DefWindowProcW(hwnd, msg, wParam, lParam)
	}

	return 0
}
