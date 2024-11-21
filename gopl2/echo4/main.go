package main

import (
	"echo4/tempconv"
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "ommit trainling newline")
var sep = flag.String("s", " ", "seperator")

func main() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}

	// expermineting new function.
	p := new(int) // p is of type *int, a pointer to an int value.
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)

	// package tempconv
	fmt.Println(tempconv.CToF(24))

}
