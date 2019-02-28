package main

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/lyderic/tools"
)

func getFileString(f string) (s string, err error) {
	if _, err = os.Stat(f); os.IsNotExist(err) {
		return
	}
	content, err := ioutil.ReadFile(f)
	if err != nil {
		return
	}
	s = strings.TrimSpace(string(content))
	return
}

func debug(s string, a ...interface{}) {
	if dbg {
		tools.PrintColorf(tools.YELLOW, s, a...)
	}

}
