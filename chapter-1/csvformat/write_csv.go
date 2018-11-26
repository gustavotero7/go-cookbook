package csvformat

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
)

// A Book has an author and title
type Book struct {
	Author string
	Title  string
}

// Books is a named type for an array of books
type Books []Book

// ToCSV takes a set of books and writes to an io.Writer
// it returns any errors
func (books *Books) ToCSV(w io.Writer) error {
	n := csv.NewWriter(w)
	// Write heades
	if err := n.Write([]string{"Author", "Title"}); err != nil {
		return err
	}

	// Write books data
	for _, book := range *books {
		if err := n.Write([]string{
			book.Author,
			book.Title,
		}); err != nil {
			return err
		}
	}
	n.Flush()
	return n.Error()
}

// WriteCSVOutput initializes a set of books
// and writes ToCSV() result to os.Stdout
func WriteCSVOutput() error {
	b := Books{
		Book{
			Author: "F Scott Fitzgerald",
			Title:  "The Great Gatsby",
		},
		Book{
			Author: "J D Salinger",
			Title:  "The Catcher in the Rye",
		},
	}
	return b.ToCSV(os.Stdout)
}

// WriteCSVBuffer returns a bug¿ffer csv for a set of books
func WriteCSVBuffer() (*bytes.Buffer, error) {
	b := Books{
		Book{
			Author: "F Scott Fitzgerald",
			Title:  "The Great Gatsby",
		},
		Book{
			Author: "J D Salinger",
			Title:  "The Catcher in the Rye",
		},
	}
	w := &bytes.Buffer{}
	err := b.ToCSV(w)
	return w, err
}
