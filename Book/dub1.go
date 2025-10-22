package Book

import (
	"bufio"
	"fmt"
	"os"
)

func Duplicate1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter lines (press Ctrl+D or Ctrl+Z to end):")
	for input.Scan() {
		line := input.Text()
		counts[line]++
	}

	fmt.Println("\nDuplicate lines:")
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d %s\n", n, line)
		}
	}

	// Check for errors
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}
