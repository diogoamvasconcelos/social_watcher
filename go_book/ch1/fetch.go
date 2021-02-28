package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	urlPrefix = "http://"
)

func main() {
	start := time.Now()
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, urlPrefix) {
			url = fmt.Sprintf("%s%s", urlPrefix, url)
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch failed: %s: %v\n", url, err)
			continue
		}

		fmt.Printf("\n======= New URL:%s ========\nStatusCode: %s\n", url, resp.Status)
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch body read failed: %s: %v\n", url, err)
			continue
		}
	}
	fmt.Printf("\nTook: %.3fs", time.Since(start).Seconds())
}
