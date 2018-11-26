package main

import (
	"flag"
	"fmt"
)

// Config will be the holder for our flags
type Config struct {
	subject      string
	isAwesome    bool
	howAwesome   int
	countTheWays CountTheWays
}

// Setup initializes a config from flags that are passed in
func (c *Config) Setup() {
	// You can set a flag directly like so:
	// var someVar = flag.String("flag_name", "default_val", "description")
	// but in practice puttng it in a struct is genereally better
	// Longhand
	flag.StringVar(&c.subject, "subject", "", "subject is a string, it defaults to empty")
	// Shorthand
	flag.StringVar(&c.subject, "s", "", "subject is a string, it defaults to empty (shorthand)")

	flag.BoolVar(&c.isAwesome, "isAwesome", false, "Is it awesome or what?")
	flag.IntVar(&c.howAwesome, "howAwesome", 10, "how awesome out of 10?")

	// Custom variable type
	flag.Var(&c.countTheWays, "c", "comma separated list of integers")
}

// GetMessage uses all of the internal
// config vars and returns a sentence
func (c *Config) GetMessage() string {
	msg := c.subject
	if c.isAwesome {
		msg += " is awesome"
	} else {
		msg += " is NOT awesome"
	}

	msg = fmt.Sprintf("%s with a certainty of %d out of 10. Let me count the ways %s", msg, c.howAwesome, c.countTheWays.String())
	return msg
}
