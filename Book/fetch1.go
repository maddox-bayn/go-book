package Book

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func Fetch() {

	client := http.Client{
		Timeout: time.Second,
	}
	for _, url := range os.Args[1:] {
		resp, err := client.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch: %v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s %v\n", url, err)
			os.Exit(1)
		}
	}
}
