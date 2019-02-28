package main

import "fmt"

func hostname() {
	fmt.Println(information.Hostname)
}

func model() {
	fmt.Println(information.Model)
}

func temperature() {
	fmt.Println(information.Celsius, information.Farenheit)
}

func celsius() {
	fmt.Println(information.Celsius)
}

func farenheit() {
	fmt.Println(information.Farenheit)
}

func networking() {
	soon("networking")
}

func dump() {
	soon("json")
}

func soon(action string) {
	fmt.Printf("Action %q: in construction ...\n", action)
}
