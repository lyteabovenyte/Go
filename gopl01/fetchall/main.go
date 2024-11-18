package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	start := time.Now()
	// could be great to have buffered channel with anticipation on uris
	ch := make(chan string) // unbuffered channel which blocks the send until one receive happens
	for _, uri := range os.Args[1:] {
		go fetch(uri, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("total elapsed: %.2fs\n", time.Since(start).Seconds())
}

func fetch(uri string, ch chan<- string) {
	start := time.Now()
	res, err := http.Get(uri)
	if err != nil {
		ch <- fmt.Sprintf("fetch error: %v\n", err)
		return
	}
	f, err := os.Create(url.QueryEscape(uri))
	if err != nil {
		ch <- fmt.Sprintf("creating file: %v\n", err)
	}
	nbytes, err := io.Copy(f, res.Body) // writing the result into a file
	res.Body.Close()

	if closeErr := f.Close(); err != nil {
		err = closeErr
	}

	if err != nil {
		ch <- fmt.Sprintf("copy error: %v\n", err)
		return
	}
	elapsed := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", elapsed, nbytes, uri)
}
