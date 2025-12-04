package main

import (
	"bufio"
	"fmt"
	"os"
)

// reverse returns the rune-wise reversed string of s.
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter 3 lines (press Enter after each). Finish early with Ctrl+Z Enter on PowerShell:")

	// Buffer the lines first, then print the reversed lines all at once.
	var lines []string
	for i := 0; i < 3; i++ {
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "error reading input:", err)
			} else {
				fmt.Fprintln(os.Stderr, "EOF reached before 3 lines (printing any lines read so far)")
			}
			break
		}
		lines = append(lines, scanner.Text())
	}

	// Now print the reversed lines
	for _, l := range lines {
		fmt.Println(reverse(l))
	}
}
