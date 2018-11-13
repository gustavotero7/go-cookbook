package main

import "github.com/gustavotero7/go-cookbook/chapter-1/bytestrings"

func main() {
	if err := bytestrings.WorkWithBuffer(); err != nil {
		panic(err)
	}
	bytestrings.SearchString()
	bytestrings.ModifyString()
	bytestrings.StringReader()
}
