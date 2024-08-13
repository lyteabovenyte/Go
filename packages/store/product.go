// Package Store is used to provide types and methods
// commonly used for online sales.
package store

var standardTax *TaxRate = NewTaxRate(0.25, 25)

// Product describe an item for sale.
type Product struct {
	Name, Category string // Name and Type of the item.
	price          float64
}

func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

func (p *Product) Price() float64 {
	return standardTax.CalTax(p)
}

func (p *Product) SetPrice(newPrice float64) {
	p.price = newPrice
}
