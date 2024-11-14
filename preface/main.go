package main

import "fmt"

func main() {

	// lexical scope of golang
	var x string = "hello"
	y := 2

	if true {
		x := "hi" // not the same x outside of the scope of this if.
		fmt.Printf("x in the if scope memory address: %p\n", &x)
		fmt.Printf("x in the if scope value: %v\n", x)
		y = 1
		fmt.Printf("y in the if scope memory address: %p\n", &y)
		fmt.Printf("y in the if scope value: %v\n", y)
	}
	fmt.Printf("x OUTSIDE the if scope memory address: %p\n", &x)
	fmt.Printf("y OUTSIDE the if scope memory address: %p\n", &y)

	// operator overloading
	m, n := complex{1, 2}, complex{1, 1}
	m.Add(n)
	fmt.Println(m.value1, m.value2)
}

type complex struct {
	value1 int
	value2 int
}

func (c *complex) Add(c2 complex) {
	c.value1 += c2.value1
	c.value2 += c2.value2
}
