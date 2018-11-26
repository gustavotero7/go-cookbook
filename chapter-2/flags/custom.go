package main

import (
	"fmt"
	"strconv"
	"strings"
)

// CountTheWays is a custom type that we will read a flag in to
type CountTheWays []int

func (c *CountTheWays) String() string {
	result := ""
	for _, v := range *c {
		if len(result) > 0 {
			result += " ... "
		}
		result += fmt.Sprint(v)
	}
	return result
}

// Set will be used by flag pkg
func (c *CountTheWays) Set(value string) error {
	values := strings.Split(value, ",")
	for _, v := range values {
		i, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*c = append(*c, i)
	}
	return nil
}
