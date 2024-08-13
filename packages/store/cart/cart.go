package cart

import "packages/store"

type Cart struct {
	CustomerName string
	Products     []store.Product
}

func (c *Cart) GetTotal() (total float64) {
	for _, p := range c.Products {
		total += p.Price()
	}
	return
}

func (c *Cart) GetTotalWithTax(p *store.Product, t *store.TaxRate) (total float64) {
	for _, p := range c.Products {
		total += t.CalTax(&p)
	}
	return
}
