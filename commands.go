package main

type Command struct {
	Letter      string
	Name        string
	Header      string
	Description string
	Execute     bool
	Action      func()
}

var commands = []Command{
	{"h", "hostname", "Hostname", "show hostname", false, hostname},
	{"m", "model", "Model", "show model", false, model},
	{"t", "temperature", "T\u00b0", "show temperature", false, temperature},
	{"c", "celsius", "T\u00b0C", "show temperature (celsius only)", false, celsius},
	{"f", "farenheit", "T\u00b0F", "show temperature (farenheit only)", false, farenheit},
	{"n", "networking", "Networking", "show networking", false, networking},
	{"j", "json", "", "dump all information as json", false, dump},
}
