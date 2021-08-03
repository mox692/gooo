package main

import (
	"fmt"
	"os"

	"github.com/mox692/gooo/gobin"
)

var fileName = "/Users/kimuramotoyuki/go/src/github.com/mox692/gooo/gobin/cmd/main"

func main() {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}

	lib, err := gobin.Parse(f)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}

	for _, v := range lib {
		fmt.Printf("Name: %s, Version: %s", v.Name, v.Version)
	}
}
