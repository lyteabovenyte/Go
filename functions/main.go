package main

import "fmt"

func printPrice() *float64 {
	kayak := 27.500
	tax := kayak * 0.2
	fmt.Println("kayak: ", kayak, "tax: ", tax)
	return &kayak
}

func PrintSupplier(product string, supp ...string) {
	if supp == nil {
		fmt.Println("product: ", product, "supplier: ", "None")
	} else {
		for _, s := range supp {
			fmt.Println("product: ", product, "supplier: ", s)
		}
	}
}

func swapValues(a, b int) {
	fmt.Println("a -->", a, "b -->", b)
	temp := a
	a = b
	b = temp
	fmt.Println("swapped -->", "a -->", a, "b -->", b)
}

func swapWithPointer(a, b *int) {
	fmt.Println("before swap -->", "a --> ", *a, "b -->", *b)
	temp := *a
	*a = *b
	*b = temp
	fmt.Println("after swap -->", "a -->", *a, "b -->", *b)
}

func productPrice(product string, price, tax int) (pWithTax int) {
	pWithTax = price * tax
	return // it will return the specified return variable.
}

func main() {
	fmt.Println("about to call the printPrice func ...")
	k := printPrice()
	fmt.Println(*k)
	fmt.Println("called the function.")

	fmt.Println("---------------")
	PrintSupplier("soccer Ball")

	fmt.Println("-----exploring variadic---------")

	names := []any{"amir", "mahkam", "ala", 2}

	fmt.Println(names...)

	fmt.Println("------------------")
	x := 2
	y := 4
	swapValues(x, y)
	fmt.Println(x, y)

	fmt.Println("------pointer swap--------")
	swapWithPointer(&x, &y)
	fmt.Println("after swap -->", x, y)

	fmt.Println("--------------")

	re := productPrice("book", 12, 2)
	fmt.Println(re)
}
