package filedirs

import (
	"errors"
	"io"
	"os"
)

// Operate manipulates files and directories
func Operate() error {
	// this 0777 is similar to what you'd see with chown
	// on a command line this will create a directory
	// /tmp/example, you may also use an absolute path
	// Instead of a realtive one
	if err := os.Mkdir("example_dir", os.FileMode(0777)); err != nil {
		return err
	}

	// go to the /tmp directory
	if err := os.Chdir("example_dir"); err != nil {
		return err
	}

	// f is a generic file object
	// it also implements multiple interfaces and can be used as a reader or writer
	// if the correct bits are set when opening
	f, err := os.Create("test.txt")
	if err != nil {
		return err
	}

	// we write a known-length value to the file and validate it wrote correctly
	value := []byte("Hellon")
	count, err := f.Write(value)
	if err != nil {
		return err
	}
	if count != len(value) {
		return errors.New("Incorrect length returned from write")
	}

	if err := f.Close(); err != nil {
		return err
	}

	// Read the file
	f, err = os.Open("test.txt")
	if err != nil {
		return err
	}

	io.Copy(os.Stdout, f)

	if err := f.Close(); err != nil {
		return err
	}

	// go to the /temp directory
	if err := os.Chdir(".."); err != nil {
		return err
	}

	// cleanup, os.RemoveAll can be dangerous if you point at the wrong directory,
	// use user input, and especially if you run as root
	if err := os.RemoveAll("example_dir"); err != nil {
		return err
	}

	return nil
}
