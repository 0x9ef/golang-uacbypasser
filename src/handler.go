package guacbypasser

func Init(method string, target string) error {
	switch method {
	case "hkcu":
		if err := pHkcu(target); err != nil {
			return err
		}
	case "hklm":
		if err := pHklm(target); err != nil {
			return err
		}
	case "ifeo":
		if err := pIfeo(target); err != nil {
			return err
		}
	case "schtasks":
		if err := pSchtasks(target); err != nil {
			return err
		}
	case "userinit":
		if err := pUserinit(target); err != nil {
			return err
		}
	case "wmic":
		if err := pWmic(target); err != nil {
			return err
		}
	case "computerdefaults":
		if err := tComputerDefaults(target); err != nil {
			return err
		}
	case "eventvwr":
		if err := tEventvwr(target); err != nil {
			return err
		}
	case "fodhelper":
		if err := tFodhelper(target); err != nil {
			return err
		}
	case "sdcltcontrol":
		if err := tSdcltcontrol(target); err != nil {
			return err
		}
	case "silentcleanup":
		if err := tSilentCleanup(target); err != nil {
			return err
		}
	case "slui":
		if err := tSlui(target); err != nil {
			return err
		}
	}
	return nil
}
