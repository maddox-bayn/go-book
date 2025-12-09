package Book

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

func Bayn() {
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

	// Create a buffered writer and write all reversed lines to it,
	// then flush once at the end for better performance.
	bw := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := bw.Flush(); err != nil {
			fmt.Fprintln(os.Stderr, "error flushing output:", err)
		}
	}()

	// Now write the reversed lines to the buffer
	for _, l := range lines {
		if _, err := bw.WriteString(reverse(l) + "\n"); err != nil {
			fmt.Fprintln(os.Stderr, "error writing output:", err)
			break
		}
	}
}
