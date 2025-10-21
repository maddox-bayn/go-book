package book

import (
	"fmt"
	"os"
	"strings"
)

func PrintOsArgs() {
	var args, sep string
	for _, arg := range os.Args[1:] {
		args += sep + arg
		sep = " "
	}
	fmt.Println(args)
	fmt.Println(strings.Join(os.Args[1:], " "))
}
