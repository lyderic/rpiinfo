package main

import (
	"flag"
	"fmt"
)

func usage() {
	fmt.Printf("rpi-info v.%s (c) Lyderic Landry, London 2019\n", VERSION)
	fmt.Println("Usage: rpi-info [options] <command> <command> ...")
	fmt.Println("Commands:")
	for _, command := range commands {
		fmt.Println(command)
	}
	fmt.Println("Options:")
	flag.PrintDefaults()
}
