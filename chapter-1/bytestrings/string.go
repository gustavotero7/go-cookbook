package bytestrings

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// SearchString shows a number of methods for searching a string
func SearchString() {
	s := "this is a test"

	// returns true because s contains the wor this
	fmt.Println(strings.Contains(s, "this"))

	// Returns true beause s contains the leter a
	// would also match if it contained b or c
	fmt.Println(strings.ContainsAny(s, "abc"))

	// Returns true because s ends with test
	fmt.Println(strings.HasSuffix(s, "test"))
}

// ModifyString _
func ModifyString() {
	s := "simple string"

	// Prints [simple string]
	fmt.Println(strings.Split(s, " "))

	// prints "Simple String"
	fmt.Println(strings.Title(s))

	// prints "simple string"; all trailing and leading space is removed
	s = " simple string "
	fmt.Println(strings.TrimSpace(s))
}

// StringReader demostrates how to create an io.Reader interface quickly with a string
func StringReader() {
	s := "simple string"
	r := strings.NewReader(s)

	// Prints s on Stdout
	io.Copy(os.Stdout, r)
}
