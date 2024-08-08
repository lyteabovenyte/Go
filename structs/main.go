package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Product struct {
	name, category string
	price          float64
}

// produce constant using structs is just doable using functions that return a struct with initialized values.
func football() Product {
	return Product{
		name:     "footbal",
		category: "sports",
		price:    40.59,
	}
}

type Supplier struct {
	name, city string
}

type Company struct {
	name, category string
	tax            float64
	supp           *Supplier
}


// using anonymous struct types as function parameter --> anon struct don't have name.
func WriteNames(val struct {
	name, category string
	price          float64
}) {
	fmt.Println("name:", val.name)
}

func calTax(product *Product) *Product {
	if product.price > 100 {
		product.price += product.price * 0.2
	}
	return product // returns a pointer to a product. --> store the memory address of the product.
}

func calculate(product Product) Product {
	if product.price > 100 {
		product.price += product.price * 0.2
	}
	return product
}

func newProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

func companyCons(name, category string, tax float64, supplier *Supplier) *Company {
	return &Company{
		name: name,
		category: category,
		tax: tax,
		supp: supplier,
	}
}

func main() {

	kayak := Product{
		name:     "kayak",
		category: "waterSport",
		price:    35.50,
	}

	fmt.Printf("the type is %T and the value is %v\n", kayak, kayak)
	fmt.Println(kayak.name, kayak.category, kayak.price)
	kayak.price = 50.00
	fmt.Println(kayak.price)

	fmt.Println("--------------")

	fmt.Println(football())

	fmt.Println("------------")

	swim := Product{
		name:     "swimming",
		category: "sports",
		price:    50.00,
	}

	var lifeJacket Product

	needs := map[string]Product{
		"football":   football(), // produce a constant product for football.
		"swimming":   swim,
		"lifeJacket": lifeJacket, // the zero value for the types.
		"nil":        Product{},
	}

	fmt.Println("---------")

	if val, ok := needs["lifeJacket"]; ok {
		fmt.Println(val)
	} else {
		fmt.Println("not initialized.")
	}

	fmt.Println("----------")

	var transaction float64 = 0
	for _, val := range needs {
		transaction += float64(val.price)
	}

	fmt.Printf("full transaction: %v\n", transaction)

	fmt.Println("------------")

	type StockLevel struct {
		Product
		Alternative Product
		count       int
	}

	footabllStock := StockLevel{
		Product: Product{name: "amir", category: "cs", price: 50.00},
		Alternative: Product{
			name:     "alt amir",
			category: "alt-cs",
			price:    50.50,
		},
		count: 5,
	}

	fmt.Println(footabllStock.price)
	fmt.Println("product: ", footabllStock.Product, "alt-product:", footabllStock.Alternative)

	fmt.Println("--------------")

	WriteNames(football())

	fmt.Println("------------")

	var builder strings.Builder

	json.NewEncoder(&builder).Encode(struct {
		ProductName  string
		ProductPrice float64
	}{
		ProductName:  football().name,
		ProductPrice: football().price,
	})

	fmt.Printf("the type is %T, and the value is %v\n", builder.String(), builder.String())

	fmt.Println("----------------")

	p := calTax(&Product{
		name:     "p",
		category: "test-case",
		price:    102.05,
	})

	p2 := calculate(Product{
		name:     "p",
		category: "test-case",
		price:    102.05,
	})

	fmt.Printf("pointer app --> the type is %T and the value is %v\n", *p, p)
	fmt.Printf("functional app --> the type is %T and the value is %v\n", p2, p2)

	fmt.Println("--------------")

	prods := [2]*Product{
		newProduct("amir", "senior", 250.00), // returns a pointer to a product.
		newProduct("seyf", "senior", 270.00), // returns a pointer to a product.
	}

	for _, p := range prods {
		fmt.Printf("the type is %T and the value is %v\n", p, p)
	}

	fmt.Println("------------")

	s1 := &Supplier{
		name: "ship",
		city: "US-California",
	}

	s2 := &Supplier{
		name: "plane",
		city: "IR-Tehran",
	}

	comps := [2]*Company{
		companyCons("amazon", "shop", 0.2, s1),
		companyCons("skybooks", "bookshelf", 0.1, s2),
	}

	for _, company := range comps {
		fmt.Println(company.supp.city)
	}

	fmt.Println("-------------")

	

}
