package errcheck

import (
	"fmt"
	"io"
	"os"
)

// Close will check the error returned from the given io.Closer and print the
// error and exit on error
func Close(c io.Closer) {
	if err := c.Close(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Exit will print a value and then exit
func Exit(v interface{}) {
	switch v.(type) {
	case error:
		fmt.Println(v)
	case string:
		fmt.Printf("error: %s\n", v)
	default:
		fmt.Printf("error (unknown): %s\n", v)
	}
	os.Exit(1)
}
