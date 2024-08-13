package main

import (
	"fmt"
	_ "packages/data"
	. "packages/fmt"
	"packages/store"
	"packages/store/cart"
	"github.com/fatih/color"
)

func main() {
	p := store.NewProduct("amir", "senior", 120.00)

	fmt.Println(p.Price())
	// fmt.Printf("the type is %T and the value is %v\n", p.Price, p.Price)

	fmt.Println("----------------")

	p2 := store.NewProduct("seyf", "senior", 210.00)

	fmt.Println("price with standard tax-->", p2.Price())

	tax := store.NewTaxRate(0.30, 60)

	fmt.Println("with default tax -->", tax.CalTax(p2))

	fmt.Println("-------------")

	fmt.Println(ToCurrency(p.Price()))

	fmt.Println("---------------")

	p3 := store.NewProduct("arman", "economy", 220.00)

	t := store.NewTaxRate(0.3, 120.00)

	fmt.Println("custom tax rate -->", t.CalTax(p3))

	c := cart.Cart{
		CustomerName: "apple",
		Products:     []store.Product{*p3},
	}
	fmt.Println("price for apple with standard tax-->", ToCurrency(c.GetTotal()))
	fmt.Println("price for apple with custom tax-->", ToCurrency(c.GetTotalWithTax(p, t)))

	// color.Green("Name:\n", c.CustomerName)
	color.Green(fmt.Sprintln("XXXXXXXXXXXXXXXXXXXX"))
}
