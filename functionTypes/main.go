package main

import "fmt"

type calFunc func(float64) float64 // type!!! funtionType --> an alias for function types

// func calWithTax(price float64) float64 {
// 	fmt.Print("calculated with tax -->")
// 	return price + (price * 0.2)
// }

// func calWithoutTax(price float64) float64 {
// 	fmt.Print("actual price without tax -->")
// 	return price
// }

func printPrice(product string, price float64, calculator calFunc) {
	fmt.Println("product:", product, "price:", calculator(price))
}

func selectCal(price float64) calFunc {
	if price > 100 {
		var WithTax calFunc = func (price float64) float64 {
			return price + ( price * 0.2)
		}
		return WithTax
	} else {
		var WithoutTax calFunc = func(price float64) float64 {
			return price
		}
		return WithoutTax
	}
}

func main() {

	product := map[string]float64{
		"book":   150.50,
		"course": 50.50,
	}

	for key, value := range product {
		printPrice(key, value, selectCal(value))
	}
	

}
