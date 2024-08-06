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
		var WithTax calFunc = func(price float64) float64 {
			return price + (price * 0.2)
		}
		return WithTax
	} else {
		var WithoutTax calFunc = func(price float64) float64 {
			return price
		}
		return WithoutTax
	}
}

// function closure
func priceCalFactory(threshold, rate float64) calFunc {
	return func(price float64) float64 {
		if price > threshold {
			return price + (price * rate)
		}
		return price
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

	fmt.Println("------------------")

	watersportsProducts := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}

	soccerProducts := map[string]float64{
		"Soccer Ball": 19.50,
		"Stadium":     79500,
	}

	waterCalc := priceCalFactory(100, 0.2) // returns the function to calculate for waterSports
	soccerCals := priceCalFactory(50, 0.1) // returns the function to calculate for soccerSports

	for key, value := range watersportsProducts {
		printPrice(key, value, waterCalc)
	}

	for key, value := range soccerProducts {
		printPrice(key, value, soccerCals)
	}

}
