package main

import (
	"fmt"
	"os"
)

func check(fn string, err error) {
	if err != nil {
		fmt.Printf("error in %s: %s\n", fn, err)
		os.Exit(1)
	}
}
