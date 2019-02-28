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

type NIC struct {
	Name      string
	IpAddress string
	State     string
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

func displayNetworking(networking []NIC) (output string) {
	var buffer strings.Builder
	buffer.WriteString(fmt.Sprintf(" %-8.8s %-8.8s %s\n", "Name", "State", "IP Address"))
	buffer.WriteString(fmt.Sprintf(" %-8.8s %-8.8s %s\n", "----", "-----", "----------"))
	for idx, nic := range networking {
		buffer.WriteString(nic.String())
		if idx != len(networking)-1 {
			buffer.WriteString("\n")
		}
	}
	return buffer.String()
}

func (nic NIC) String() string {
	return fmt.Sprintf(" %-8.8s %-8.8s %s", nic.Name, nic.State, nic.IpAddress)
}
