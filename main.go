package main

import "fmt"

func main() {
	s := []int{2, 3, 4}
	//s = append(s, 0) // Step 1: Grow the slice by one (placeholder)
	fmt.Printf("After append: %v\n", s)
	fmt.Println(s[1:])
	copy(s[1:], s)
	fmt.Printf("After copy: %v\n", s)
	s[0] = 1                      // Step 3: Set the first element
	fmt.Println(s[0], s[1], s[2]) // Output: 1 2 3 4
}
