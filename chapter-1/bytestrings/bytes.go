package bytestrings

import (
	"bufio"
	"bytes"
	"fmt"
)

// WorkWithBuffer will make use of the buffer created by the
// Buffer fuction
func WorkWithBuffer() error {
	rawString := "it's easy to encode unicode into a byte array"
	b := Buffer(rawString)

	// We can quickly convert a buffer back into bytes with
	// b.Bytes() or a string with b.String
	fmt.Println(b.String())

	// Because this is and io.Reader we can make use of
	// generic io reader functions such as
	s, err := toString(b)
	if err != nil {
		return err
	}
	fmt.Println(s)

	// We can also take our bytes and create a bytes reader
	// these readers implement io.Reader, io.ReaderAt, io.WriterTo
	// io.Seeker, io.ByteScanner, and io.RuneScanner interfaces
	reader := bytes.NewReader([]byte(rawString))

	// We can also plug it in to a scanner that allows
	// buffered reading and tokenization
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	// iterate over all scan events
	for scanner.Scan() {
		fmt.Print(scanner.Text())
	}

	return nil
}
