package main

import "github.com/gustavotero7/go-cookbook/chapter-1/tempfiles"

func main() {
	if err := tempfiles.WorkWithTemp(); err != nil {
		panic(err)
	}
}
