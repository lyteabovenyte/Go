package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s string
	sep := " "
	for i := 1; i < len(os.Args); i++ {
		// skip file name --> if we want to access build location,
		// we can access it via os.Args[0]
		s += os.Args[i] + sep
	}
	fmt.Println(s)

	fmt.Println(os.Args[1:])                            // a list of all the arguments
	fmt.Printf("the type of os.Args --> %T\n", os.Args) // the type of os.Args --> []string

	fmt.Printf("echo2 result -->%s\n", echo2())
}

func echo2() string {
	s := " "
	sep := 1
	for _, arg := range os.Args[1:] {
		s += strconv.Itoa(sep) + ")" + arg + " "
		sep++
	}
	return s
}
