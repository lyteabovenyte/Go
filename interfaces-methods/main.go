package main

import "fmt"

func newProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

type productList []*Product // type alias --> using methods for any types.

func (product *productList) calCategoryTotals() map[string]float64 {
	total := make(map[string]float64)
	for _, p := range *product {
		total[p.category] += p.price
	}
	return total
}

func (p *Product) addDiscount(rate float64) {
	p.price = p.price - (p.price * rate)
}

func (p *productList) addDiscount(rate float64) {
	for _, product := range *p {
		product.price = product.price - (product.price * rate)
	}
}

type Expenses interface {
	getName() string
	getCost(annual bool) float64
}

func calTotal(expenses []Expenses) (total float64) {
	total = 0
	for _, item := range expenses {
		total += item.getCost(true)
	}
	return
}

func main() {
	products := productList{
		newProduct("amir", "senior", 160.00),
		newProduct("seyf", "senior", 190.00),
		newProduct("arman", "economy", 220.00),
	}

	for cat, price := range products.calCategoryTotals() {
		fmt.Println("category:", cat, "price:", price)
	}

	fmt.Println("------after------")

	for _, p := range products {
		p.addDiscount(0.05)
	}

	for cat, price := range products.calCategoryTotals() {
		fmt.Println("category:", cat, "price:", price)
	}

	fmt.Println("XXXXXXXXXX===============after split packagin=============XXXXXXXXXXXx")

	amir := Product{"amir", "senior", 150.00}
	insurance := Service{"insurance for amir", 12, 25.00}

	fmt.Println("insurance price:", insurance.durationMonths*int(insurance.monthlyFee), "for -->", amir.name)

	expenses := []Expenses{
		Product{"seyf", "senior", 190.00},
		Service{"seyf insurance", 36, 25.00},
	}

	fmt.Println("total -->", calTotal(expenses))

	fmt.Println("--------after account---------")

	ex_for_amir := []Expenses{
		Product{"book", "cs", 25.50},
		Service{"AWS", 12, 12.50},
	}

	amir_acc := Account{
		AccountNumber: 1,
		expenses:      ex_for_amir,
	}

	fmt.Println("total cost for amir with ID:", amir_acc.AccountNumber, "-->", calTotal(amir_acc.expenses))

	fmt.Println("----------------")

	pr := Product{"lamp", "blue-yellow", 10.00}

	var exp Expenses = &pr

	pr.price = 12.00

	fmt.Println("pr.price -->", pr.price)
	fmt.Println("expense.price -->", exp.getCost(false))

	fmt.Println("-----------------")

	for _, ele := range ex_for_amir {
		switch value := ele.(type) {
		case Service:
			fmt.Println("service -->", value.description, value.getCost(true))
		case Product:
			fmt.Println("product -->", value.name, value.getCost(false))
		default:
			fmt.Println("this type doesn't implement the Expense interface.")
		}
	}
}
