package main

import "fmt"

func main() {
	var s []string
	fmt.Println("s --> ", &s)
	fmt.Println(s == nil)
	p := []string{}
	fmt.Println("p -->", &p)
	fmt.Println(p == nil)
}
