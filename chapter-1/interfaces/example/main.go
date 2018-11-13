package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/gustavotero7/go-cookbook/chapter-1/interfaces"
)

func main() {
	in := bytes.NewReader([]byte("qwertyuiopplkjhgfdsaZXCVBNM"))
	log.Println("LEN: ", in.Len())
	out := &bytes.Buffer{}
	fmt.Print("stdout on copy = ")
	if err := interfaces.Copy(in, out); err != nil {
		panic(err)
	}
	fmt.Println("out bytes buffer = ", out.String())

	fmt.Print("stdout on pipe example = ")
	if err := interfaces.PipeExample(); err != nil {
		panic(err)
	}
}
