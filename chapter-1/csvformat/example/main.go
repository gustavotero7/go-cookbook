package main

import (
	"fmt"

	"github.com/gustavotero7/go-cookbook/chapter-1/csvformat"
)

func main() {
	if err := csvformat.AddMoviesFromText(); err != nil {
		panic(err)
	}

	if err := csvformat.WriteCSVOutput(); err != nil {
		panic(err)
	}

	bufer, err := csvformat.WriteCSVBuffer()
	if err != nil {
		panic(err)
	}

	fmt.Println("Buffer = ", bufer.String())
}
