package main

import "fmt"

type Pricer interface {
	getPrice() float64
}

type Book struct {
	title string
	price float64
}

type Game struct {
	name    string
	version int
	price   float64
}

func (b *Book) getPrice() float64 {
	return b.price
}

func (g *Game) getPrice() float64 {
	return g.price
}

type items []Pricer

func main() {

	b1 := Book{
		title: "distributed services with go",
		price: 25.49,
	}

	b2 := Book{
		title: "pro go",
		price: 74.51,
	}

	g1 := Game{
		name:    "mlbb",
		version: 2,
		price:   75.00,
	}

	card := items{&b1, &b2, &g1}

	total := 0
	for _, c := range card {
		total += int(c.getPrice())
	}

	fmt.Println("total:", total)
}
