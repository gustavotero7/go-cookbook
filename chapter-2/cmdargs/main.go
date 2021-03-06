package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	c := MenuConf{}
	menu := c.SetupMenu()
	menu.Parse(os.Args[1:])

	// We use arguments to switch between commands
	// flags are also an argument
	if len(os.Args) > 1 {
		// We dont care about case
		switch strings.ToLower(os.Args[1]) {
		case "version":
			c.Version()
		case "greet":
			f := c.GetSubMenu()
			if len(os.Args) < 3 {
				f.Usage()
				return
			}
			if len(os.Args) > 3 {
				f.Parse(os.Args[3:])
			}
			c.Greet(os.Args[2])
		default:
			fmt.Println("Invalid command")
			menu.Usage()
			return
		}
	} else {
		menu.Usage()
		return
	}
}
