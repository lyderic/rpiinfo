package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lyderic/tools"
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
	// There is an extra byte (00), that can be show like this, at the end of MODEL_FILE:
	// $ hexdump -C MODEL_FILE
	// We need to get rid of it
	if model[len(model)-1] == 0 {
		model = model[:len(model)-1]
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
		tools.PrintColorf(tools.RED, "%s\nCannot get CPU temperature! Are you sure you are running this on a Raspberry Pi?\n", err)
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
