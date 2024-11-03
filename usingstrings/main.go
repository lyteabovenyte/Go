package main

import "fmt"

func main() {
	fmt.Println("kayak:", Kayak.Name, "price:", Kayak.Price)
	fmt.Println("______________")

	fmt.Printf("name: %v, and price $%4.2f\n", Kayak.Name, Kayak.Price)

	fmt.Println("-----------------")

	name, _ := getProduct(2)
	fmt.Println(name)

	_, e := getProduct(15)
	fmt.Println(e.Error())
	fmt.Println(e)

	p := Products[1]

	fmt.Printf("p --> %#v", p)
}

func getProduct(index int) (name string, err error) {
	if len(Products) > index {
		name = fmt.Sprintf("product name: %v\n", Products[index].Name)
	} else {
		err = fmt.Errorf("error: product not found for index %v", index)
	}
	return
}
