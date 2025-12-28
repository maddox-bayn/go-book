package main

import "fmt"

func multiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

func main() {
	var fns []func()

	for i := 0; i < 3; i++ {
		i := i // shadow variable
		fns = append(fns, func() {
			fmt.Println(i)
		})
	}

	for _, fn := range fns {
		fn()
	}
	mult := multiplier(4)
	fmt.Println(mult(4))
}
