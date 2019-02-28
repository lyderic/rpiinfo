package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/lyderic/tools"
)

func init() {
	if os.Getenv("RPIINFO_DEBUG") == "on" {
		dbg = true
	}
	debug("*** RPIINFO DEBUG MODE ON ***\n")
}

func main() {
	for idx,_ := range commands {
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

func found(arg string) (command Command, ok bool) {
	idx := 0
	for idx, command = range commands {
		if command.Letter == arg || command.Name == arg {
			debug("Command %#v found at index %d\n", command, idx)
			ok = true
			return
		}
	}
	ok = false
	return
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

func getHostname() (hostname string) {
	hostname, err := os.Hostname()
	if err != nil {
		tools.PrintColorf(tools.RED, "Cannot get hostname! %s\n", err)
	}
	return
}

func getModel() (model string) {
	model, err := getFileString(MODEL_FILE)
	if err != nil {
		tools.PrintColorf(tools.RED, "%s\nCannot get model. Are you sure you are running this on a Raspberry Pi?\n", err)
	}
	return
}

func getCelsius() (celsius string) {
	return fmt.Sprintf("%.1f\u00b0C", getCelsiusTemperature())
}

func getFarenheit() (farenheit string) {
	return fmt.Sprintf("%.1f\u00b0F", (getCelsiusTemperature()*1.8)+32)
}

func getCelsiusTemperature() (celsius float64) {
	rawtemperature, err := getFileString(CPU_TEMP_FILE)
	if err != nil {
		tools.PrintColorf(tools.RED, "Cannot get CPU temperature! Are you sure you are running this on a Raspberry Pi? %s\n", err)
		return
	}
	kcelsius, err := strconv.ParseFloat(rawtemperature, 64)
	if err != nil {
		tools.PrintColorf(tools.RED, "Cannot parse raw temperature from %q: %s\n", CPU_TEMP_FILE, rawtemperature)
		return
	}
	celsius = kcelsius / 1000
	return
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
