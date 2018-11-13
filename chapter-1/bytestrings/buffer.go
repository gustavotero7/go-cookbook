package bytestrings

import (
	"bytes"
	"io"
	"io/ioutil"
)

// Buffer demostrates some tricks for initializing bytes Buffers
// These buffers implement an io.Reader interface
func Buffer(rawString string) *bytes.Buffer {
	// we'll start with a string encoded into raw bytes
	rawBytes := []byte(rawString)

	// There are a number of ways to create abuffer from
	// the raw bytes or from the original string
	var b = new(bytes.Buffer)
	b.Write(rawBytes)

	// Alternatively
	b = bytes.NewBuffer(rawBytes)

	// And avoiding the initial byte array altogether
	b = bytes.NewBufferString(rawString)

	return b
}

func toString(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", nil
	}
	return string(b), nil
}
