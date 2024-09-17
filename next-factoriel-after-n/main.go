package main

import "fmt"

func main() {
	p := 1
	n := 2
	x := 0
	for i := 1; p < n; i++ {
		p = p * i
		x = i
	}
	fmt.Printf("first factoriel after %d, is %d, for %d!\n", n, p, x)
}
