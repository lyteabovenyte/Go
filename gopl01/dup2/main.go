// package main introduce some kind of
// "streaming" approach that read files and split
// them into lines
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	totalCount := make(map[string]int)
	args := os.Args[1:]
	if len(args) == 0 {
		count := make(map[string]int)
		countLines(os.Stdin, count)
		countLines(os.Stdin, totalCount)
	} else {
		for _, f := range args {
			file, err := os.Open(f)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			// making count local to the exact file.
			count := make(map[string]int)
			countLines(file, count)
			for line, n := range count {
				if n > 1 {
					fmt.Printf("%s file: %s: %d times\n", file.Name(), line, n)
				}
			}
			// aggregate's count for each line in total
			totalFile, err := os.Open(f)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			} else {
				countLines(totalFile, totalCount)
			}
			file.Close()
		}
	}

	fmt.Println("total aggregate counts in files:")
	for item, dup := range totalCount {
		if dup > 1 {
			fmt.Printf("%s: %d times\n", item, dup)
		}
	}
}

func countLines(f *os.File, count map[string]int) {
	buf := bufio.NewScanner(f)
	for buf.Scan() {
		if buf.Text() == "" {
			break
		}
		count[buf.Text()]++
	}
}
