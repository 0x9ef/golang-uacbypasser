package uacbypass

import (
	"strconv"

	"golang.org/x/sys/windows/registry"
)

func GetBuildNumber() int {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE,
		"Software\\Microsoft\\Windows NT\\CurrentVersion", registry.QUERY_VALUE)
	if err != nil {
		return 0
	}
	defer k.Close()
	bn, _, err := k.GetStringValue("CurrentBuildNumber")
	n, err := strconv.Atoi(bn)
	if err != nil {
		return 0
	}
	return n
}

func GetUACLevel() int {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE,
		"Software\\Microsoft\\Windows\\CurrentVersion\\Policies\\System", registry.QUERY_VALUE)
	if err != nil {
		return 0
	}
	defer k.Close()
	cpba, _, err := k.GetIntegerValue("ConsentPromptBehaviorAdmin")
	if err != nil {
		return 0
	}
	cpbu, _, err := k.GetIntegerValue("ConsentPromptBehaviorUser")
	if err != nil {
		return 0
	}
	posd, _, err := k.GetIntegerValue("PromptOnSecureDesktop")
	if err != nil {
		return 0
	}

	const higherLevel = 4
	const mediumLevel = 3
	const lowLevel = 2
	const noLevel = 1
	var level int
	switch {
	case cpba == 0x02 && cpbu == 0x3 && posd == 0x1:
		level = higherLevel
	case cpba == 0x05 && cpbu == 0x3 && posd == 0x1:
		level = mediumLevel
	case cpba == 0x05 && cpbu == 0x3 && posd == 0x1:
		level = lowLevel
	default:
		level = noLevel
	}
	return level
}
