package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/lyderic/tools"
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

func getNetworking() (nics []NIC) {
	nics, err := getNICs()
	if err != nil {
		panic(err)
	}
	return nics
}

func getNICs() (nics []NIC, err error) {
	cmd := exec.Command("ip", "-brief", "address")
	output, err := cmd.CombinedOutput()
	if err != nil {
		tools.PrintColorf(tools.RED, "ip command failed! %v\nOutput: %s\n", cmd.Args, string(output))
		return
	}
	scanner := bufio.NewScanner(bytes.NewReader(output))
	for scanner.Scan() {
		var nic NIC
		fields := strings.Fields(scanner.Text())
		nic.Name = fields[0]
		nic.State = fields[1]
		if len(fields) > 2 {
			nic.IpAddress = fields[2]
		}
		if nic.Name != "lo" {
			nics = append(nics, nic)
		}
	}
	return
}
