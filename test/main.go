package main

import ("fmt")
// variable shadowing
func main() {
	i := 0
	if true {
		i := 1
		fmt.Println(i)
	}
	fmt.Println(i)
}
