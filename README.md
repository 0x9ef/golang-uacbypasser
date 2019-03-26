# Golang UAC Bypasser (GUACBP)

![Image of Golang UAC Bypasser](http://s01.geekpic.net/di-Q8HD4W.jpeg)
Collection of bypass techiques written in Golang.

Techniques are found online, on different blogs and repos here on GitHub. I do not take credit for any of the findings, thanks to all the researchers. 

Rewrite of - https://github.com/rootm0s/WinPwnage to Golang. 

## Techniques implemented:
* UAC Bypass using _computerdefaults.exe_
* UAC Bypass using _eventvwr.exe_
* UAC Bypass using _fodhelper.exe_
* UAC Bypass using _HKCU Registry_
* UAC Bypass using _HKLM Registry_
* UAC Bypass using _IFEO_
* UAC Bypass using _schtasks.exe_
* UAC Bypass using _sdcltcontrol.exe_
* UAC Bypass using _silentcleanup.exe_
* UAC Bypass using _slui.exe_
* UAC Bypass using _userinit.exe_
* UAC Bypass using _wmic.exe_
 
## How to build: 
  1. `set CGO_ENABLED=0`
  2. `go build -v -a -ldflags="-w -s" -o guacbypasser.exe main.go`

## If you find error in the code or you want to support project please commit this changes. 
## **_Support project - BITCOIN: 18YsYvrQhyrtAqUcpTXpHFrQ6RHyd73dS6_**
