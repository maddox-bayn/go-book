package main

import (
	"fmt"
	"io"
	"os"

	"girhub.com/maddox-bayn/go-book/book"
)

func main() {
	// If command-line args are provided, print them using the book package.
	if len(os.Args) > 1 {
		book.PrintOsArgs()
		return
	}

	// Read from stdin in small chunks and report bytes read until EOF.
	b := make([]byte, 8)
	for {
		n, err := os.Stdin.Read(b)
		if n > 0 {
			fmt.Printf("got %d bytes: %q\n", n, b[:n])
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintf(os.Stderr, "read error: %v\n", err)
			os.Exit(1)
		}
	}
}
