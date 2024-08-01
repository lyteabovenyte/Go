// entry point of the tools module
package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello, go...")
	PrintHello()
	for i := 0; i < 5; i++ {
		PrintNumber(i)
	}
}

func PrintHello() {
	fmt.Println("hello, go")
}

func PrintNumber(number int) {
	fmt.Println(number)
}
