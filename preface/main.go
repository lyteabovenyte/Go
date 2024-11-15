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

	// example on cunstructor in golang.
	withDefault := NewComplex(15)
	fmt.Println(withDefault.value1, withDefault.value2)
	withoutCons := &complex{15, 20}
	fmt.Println(*withDefault == *withoutCons) // returns true!
	withDefault.value1 = 16
	fmt.Println(*withDefault == *withoutCons)
}

type complex struct {
	value1 int
	value2 int // default to 20.
}

func (c *complex) Add(c2 complex) {
	c.value1 += c2.value1
	c.value2 += c2.value2
}

// NewComplex constructs a new complex with default values
// for value2, but value1 is neccessary.
func NewComplex(n1 int) *complex {
	c := new(complex)
	c.value1 = n1
	c.value2 = 20
	return c
}
