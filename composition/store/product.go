package store

type Product struct {
	Name, Category string
	price          float64
	taxRate        float64 // optional
}

const (
	defaultTaxRate = 0.2
)

// basic constructors ensures that fields are properly initialized, without optional fields
// func NewProduct(name, category string, price float64) *Product {
// 	return &Product{name, category, price}
// }

// type ProductGetPrice func(p *Product)

// func (p *Product) GetPrice(opt ...ProductGetPrice) ProductGetPrice {
// 	for _, option := range opt {
// 		return func(p *Product) {
// 			option(p)
// 		}
// 	}
// 	return nil
// }

// product.GetPrice(WithTaxRate(0.4))

type ProductOptions func(p *Product)

func NewProduct(name, category string, price float64, opts ...ProductOptions) *Product {
	// init frist --> then add the options via functions.
	p := &Product{
		Name:     name,
		Category: category,
		price:    price,
	}

	for _, options := range opts {
		options(p)
	}
	return p
}

func WithTaxRate(taxRate float64) ProductOptions {
	return func(p *Product) {
		p.taxRate = taxRate
	}
}

func (p *Product) GetPrice() float64 {
	if p.taxRate != 0.0 {
		return p.price + (p.price * p.taxRate)
	}
	return p.price + (p.price * defaultTaxRate)
}
