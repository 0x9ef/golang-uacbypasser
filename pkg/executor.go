// Copyright (c) 2022 0x9ef. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.
package uacbypass

// OnceExecutor is suitable for all single-use options that clean up data immediately after their work
type OnceExecutor func(path string) error

// PersistExecutor same as OnceExecutor, but has Revert function
// that can be called manually and revert all changes which were applied.
type PersistExecutor interface {
	Exec(path string) error
	Revert() error
}
