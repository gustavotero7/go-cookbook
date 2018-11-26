package csvformat

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

// Movie will hold parsed CSV
type Movie struct {
	Title    string
	Director string
	Year     int
}

// ReadCSV gives shows some examples of processing CSV
// that is passed in as an io.Reader
func ReadCSV(b io.Reader) ([]Movie, error) {

	r := csv.NewReader(b)

	// These are some optional configuration options
	r.Comma = ';'
	r.Comment = '-'

	var movies []Movie

	// Grab and ignore the header for now
	// we may also wanna use this for a dictionary key or
	// some other form of lookup
	_, err := r.Read()
	if err != nil && err != io.EOF {
		return nil, err
	}

	// Loop until it's all processed
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		year, err := strconv.ParseInt(record[2], 10, 64)
		if err != nil {
			return nil, err
		}

		m := Movie{
			Title:    record[0],
			Director: record[1],
			Year:     int(year),
		}
		movies = append(movies, m)
	}

	return movies, nil
}

// AddMoviesFromText uses the CSV parser with a string
func AddMoviesFromText() error {
	// This is an example of us taking a string, converting
	// it into a buffer, and reading it with the CSV pkg
	in := `
		- First hour headers
		movie title;director;year released

		- Then some data
		The Crow;Alex Proyas;1994
		Star Wars: Episode VIII;Rian Johnson; 2017
	`

	b := bytes.NewBufferString(in)
	m, err := ReadCSV(b)
	if err != nil {
		return nil
	}
	fmt.Printf("%#v", m)
	return nil
}
