package main

import (
	"flag"
	"fmt"
	"os"
)

func init() {
	if os.Getenv("RPIINFO_DEBUG") == "on" {
		dbg = true
	}
	debug("*** RPIINFO DEBUG MODE ON ***\n")
}

func main() {
	for idx, _ := range commands {
		debug("adding command: %#v ...\n", commands[idx])
		flag.BoolVar(&commands[idx].Execute, commands[idx].Letter, false, commands[idx].Description)
	}
	flag.Usage = usage
	flag.Parse()
	getInformation()
	debug("GATHERED: %#v\n", information)
	if len(os.Args) == 1 {
		debug("no os flag found\n")
		fmt.Println(information)
		return
	}
	for _, command := range commands {
		debug("%v\n", command)
		if command.Execute {
			command.Action()
		}
	}
}

func getInformation() {
	debug("Gathering information from system ...\n")
	information.Hostname = getHostname()
	information.Model = getModel()
	information.Celsius = getCelsius()
	information.Farenheit = getFarenheit()
	information.Networking = getNetworking()
	return
}
