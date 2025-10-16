package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	/*var args string
	for _, arg := range os.Args[1:] {
		args += " " + arg
		//fmt.Printf("Argument %d: %s \n", i, arg)
	}
	fmt.Println(args)*/
	fmt.Println(strings.Join(os.Args[1:], " "))
}
