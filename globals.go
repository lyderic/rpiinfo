package main

const (
	VERSION       = "0.0.8"
	MODEL_FILE    = "/sys/firmware/devicetree/base/model"
	CPU_TEMP_FILE = "/sys/class/thermal/thermal_zone0/temp"
)

var (
	dbg         bool
	information Information
)
