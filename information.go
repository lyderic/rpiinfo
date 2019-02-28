package main

import (
	"fmt"
	"strings"
)

type Information struct {
	Hostname   string
	Model      string
	Celsius    string
	Farenheit  string
	Networking []NIC
}

func (information Information) String() string {
	var buffer strings.Builder
	buffer.WriteString(fmt.Sprintf("%-12.12s: %s\n", "Hostname", information.Hostname))
	buffer.WriteString(fmt.Sprintf("%-12.12s: %s\n", "Model", information.Model))
	buffer.WriteString(fmt.Sprintf("%-12.12s: %s %s\n", "Temperature", information.Celsius, information.Farenheit))
	buffer.WriteString(fmt.Sprintf("%-12.12s:\n", "Networking"))
	buffer.WriteString(displayNetworking(information.Networking))
	return buffer.String()
}
