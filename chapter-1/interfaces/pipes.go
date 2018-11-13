package interfaces

import (
	"io"
	"os"
)

func PipeExample() error {
	// The pipe reader and pipe write implement
	// io.Reader and io.Writer
	r, w := io.Pipe()

	// This needs to be run in a separate go-routine
	// as it will block waiting for the reader
	// close at the end for cleanup
	go func() {
		// For now we'll write something basic,
		// this could be also to encode json
		// base64 encode, etc.
		w.Write([]byte("test"))
		w.Close()
	}()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}

	return nil
}
