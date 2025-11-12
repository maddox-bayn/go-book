package Book

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func FetchAll() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go Fetch1(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2f elasped", time.Since(start).Seconds())
}

func Fetch1(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbyte, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading : %s %v\n", url, err)
		return
	}

	sec := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%2.f %7d %s \n", sec, nbyte, url)
}
