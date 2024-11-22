package main

import "fmt"

// main experiments shadowing concept
func main() {
	x := "hello"
	for _, x := range x {
		x = x + 'A' - 'a'
		fmt.Printf("%c", x) // uppercase the letters
	}
}
