package Book

import (
	"fmt"
	"os"
	"strings"
)

func Dup3() {
	count := make(map[string]int)
	for _, fiilename := range os.Args[1:] {
		data, err := os.ReadFile(fiilename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup:%v \n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			count[line]++
		}
	}
	for v, n := range count {
		if n > 1 {
			fmt.Printf("num:%d duplicate line%s", n, v)
		}
	}
}
