package main

import "fmt"

type Command struct {
	Letter      string
	Name        string
	Header      string
	Description string
	Action      func()
}

var commands = []Command{
	{"h", "hostname", "Hostname", "show hostname", hostname},
	{"m", "model", "Model", "show model", model},
	{"t", "temperature", "T\u00b0", "show temperature", temperature},
	{"n", "networking", "Networking", "show networking", networking},
	{"j", "json", "", "dump all information as json", dump},
}

func (command Command) String() string {
	return fmt.Sprintf("  %s  %-12.12s %s", command.Letter, command.Name, command.Description)
}
