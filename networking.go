package main

import (
	"fmt"
	"strings"
)

type NIC struct {
	Name      string
	IpAddress string
	State     string
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
