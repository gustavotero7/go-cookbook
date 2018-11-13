package interfaces

import (
	"fmt"
	"io"
	"os"
)

// Copy copies data from in to out first directly,
// then using a bbuffer. It also writes to stdout
func Copy(in io.ReadSeeker, out io.Writer) error {
	// we write to out, but also Stdout
	w := io.MultiWriter(out, os.Stdout)

	// a standard copy, this can be dangerous if there's a
	// lot of data in in
	if _, err := io.Copy(w, in); err != nil {
		return err
	}

	// Go to the begining of the stream
	in.Seek(0, 0)

	// Buffered write using 64 byte chunks
	buf := make([]byte, 8)
	if _, err := io.CopyBuffer(w, in, buf); err != nil {
		return err
	}

	// New line
	fmt.Println()
	return nil
}
