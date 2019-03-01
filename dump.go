package main

import "encoding/json"

func getJsonDump() (output string) {
	raw, err := json.MarshalIndent(information, "", "  ")
	if err != nil {
		panic(err)
	}
	output = string(raw)
	return
}
