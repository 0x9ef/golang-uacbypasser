package guacbypasser

import (
	"os/exec"
)

const (
	cmdPrefix         string = "cmd"
	cmdArgumentPrefix string = "/C"
)

type Cmd struct {
	cmdN string
	argN string

	command string
	exec    *exec.Cmd
}

func newCmd(command string) Cmd {
	return Cmd{
		cmdN: cmdPrefix,
		argN: cmdArgumentPrefix,

		command: command,
		exec:    exec.Command(cmdPrefix, cmdArgumentPrefix, command),
	}
}

type UAC struct {
	// Type of UAC - User Account COntrol
	uac string

	// Level of UAC
	// 1 - UAC Disable
	// 2 - UAC Enable with low priority settings
	// 3 - UAC Enable with high priority settings
	// 4 - UAC Enable with high priority settings
	level int
}
