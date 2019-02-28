package main

import (
	"flag"
	"fmt"
)

func usage() {
	fmt.Printf("rpi-info v.%s (c) Lyderic Landry, London 2019\n", VERSION)
	fmt.Println("Usage: rpi-info <args>")
	fmt.Println("Arguments:")
	flag.PrintDefaults()
}
