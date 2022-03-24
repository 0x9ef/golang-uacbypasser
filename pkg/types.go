// Copyright (c) 2022 0x9ef. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.
package uacbypass

type Info struct {
	Id          uint8
	Type        string
	Name        string
	Description string
	Subinfo     struct {
		Fixed       bool
		FixedIn     string
		OnlyAdmin   bool
		OnlyPayload bool
	}
}

var InfoOnceCmstp = Info{
	Id:          1,
	Type:        "once",
	Name:        "cmstp",
	Description: "Using cmstp.exe and .ini file manipulations",
	// Fixed: false
	// FixedIn: <nil>
	// Admin: false
	// Payload: true
}

var InfoOnceComputerdefaults = Info{
	Id:          2,
	Type:        "once",
	Name:        "computerdefaults",
	Description: "Using computerdefaults.exe and registry keys manipulations",
	// Fixed: false
	// FixedIn: <nil>
	// Admin: false
	// Payload: true
}

var InfoOnceEventvwr = Info{
	Id:          3,
	Type:        "once",
	Name:        "eventvwr",
	Description: "Using eventvwr.exe and registry keys manipulations",
	// Fixed: false
	// FixedIn: 15031
	// Admin: false
	// Payload: true
}

var InfoOnceFodhelper = Info{
	Id:          4,
	Type:        "once",
	Name:        "fodhelper",
	Description: "Using fodhelper.exe and registry keys manipulations",
	// Fixed: false
	// FixedIn: <nil>
	// Admin: false
	// Payload: true
}

var InfoOnceSdcltcontrol = Info{
	Id:          5,
	Type:        "once",
	Name:        "sdcltcontrol",
	Description: "Using sdclt.exe folder and registry keys manipulations",
	// Fixed: true
	// FixedIn: <nil>
	// Admin: false
	// Payload: true
}

var InfoOnceSilentcleanup = Info{
	Id:          6,
	Type:        "once",
	Name:        "silentcleanup",
	Description: "Using silentcleanup.exe and registry keys manipulations",
	// Fixed: false
	// FixedIn: <nil>
	// Admin: false
	// Payload: true
}

var InfoOnceSlui = Info{
	Id:          7,
	Name:        "slui",
	Type:        "once",
	Description: "Using slui.exe and registry keys manipulations",
	// Fixed: true
	// FixedIn: <nil>
	// Admin: false
	// Payload: true
}

var InfoOnceWsreset = Info{
	Id:          8,
	Type:        "once",
	Name:        "wsreset",
	Description: "Using wsreset.exe and registry keys manipulations",
	// Fixed: true
	// FixedIn: <nil>
	// Admin: false
	// Payload: true
}

var InfoPersistCortana = Info{
	Id:          9,
	Type:        "persist",
	Name:        "cortana",
	Description: "Using registry key class manipulation",
	// Fixed: true
	// FixedIn: <nil>
	// Admin: false
	// Payload: true
}

var InfoPersistHkcu = Info{
	Id:          10,
	Type:        "persist",
	Name:        "Using registry key (HKEY_CURRENT_USER) manipulation",
	Description: "",
	// Fixed: true
	// FixedIn: <nil>
	// Admin: false
	// Payload: true
}

var InfoPersistHklm = Info{
	Id:          11,
	Type:        "persist",
	Name:        "hklm",
	Description: "Using registry key (HKEY_LOCAL_MACHINE) manipulation",
	// Fixed: true
	// FixedIn: <nil>
	// Admin: false
	// Payload: true
}

var InfoPersistMagnifier = Info{
	Id:          12,
	Type:        "persist",
	Name:        "magnifier",
	Description: "Using magnifier.exe, Image File Execution Options debugger and accessibility application",
	// Fixed: true
	// FixedIn: <nil>
	// Admin: false
	// Payload: true
}

var InfoPersistPeople = Info{
	Id:          13,
	Type:        "persist",
	Name:        "people",
	Description: "Using registry key class manipulation",
	// Fixed: true
	// FixedIn: <nil>
	// Admin: false
	// Payload: true
}

var InfoPersistStartup = Info{
	Id:          14,
	Type:        "persist",
	Name:        "startup",
	Description: "Using malicious lnk file in startup directory",
	// Fixed: true
	// FixedIn: <nil>
	// Admin: false
	// Payload: true
}

var InfoPersistUserinit = Info{
	Id:          15,
	Type:        "persist",
	Name:        "userinit",
	Description: "Using userinit registry key manipulations",
	// Fixed: true
	// FixedIn: <nil>
	// Admin: false
	// Payload: true
}
