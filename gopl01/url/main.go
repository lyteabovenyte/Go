// fetching url's content and printing out the content
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "status: %v, error dial via http: %s\n", resp.Status, err)
			os.Exit(1)
		}
		// b, err := io.ReadAll(resp.Body)
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "copy from response body to stdount. statuscode: %v, error: %s",
				resp.StatusCode,
				err,
			)
		}
	}
}
