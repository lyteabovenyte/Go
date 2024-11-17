// kinda batch processing approach
// reading the whole file into memory
// and process them at once
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	count := make(map[string]int)
	filename := os.Args[1:]
	for _, file := range filename {
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3, error: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			count[line]++
		}
	}
	for d, n := range count {
		if n > 1 {
			fmt.Printf("%s: %d\n", d, n)
		}
	}
}
