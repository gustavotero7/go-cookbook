package main

import (
	"flag"
	"fmt"
)

func main() {
	// Initialize our setup
	c := &Config{}
	c.Setup()

	// Generally call this from main
	flag.Parse()

	fmt.Print(c.GetMessage())
}
