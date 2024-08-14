package main

import (
	. "composition/store"
	"fmt"
)

func main() {
	p := NewProduct("amir", "senior", 150.00, WithTaxRate(0.2))

	r := NewReasearcher(p, true, 2, WithWantedSalary(190.00))

	fmt.Println(r.GetPrice())

	if r.GetPrice() > r.Wanted {
		fmt.Println("can be hired")
	} else {
		fmt.Println("sorry. wanted salary is above the budget.")
	}
}
