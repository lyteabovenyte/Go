package store

const (
	defaultTaxRate = 0.2
	minthreshold   = 10
)

type TaxRate struct {
	rate, threshold float64
}

var categoryMaxPrices = map[string]float64{
	"Soccer":      250,
	"WaterSports": 150,
	"Chess":       50,
}

func init() {
	for category, price := range categoryMaxPrices {
		categoryMaxPrices[category] = price + (price * defaultTaxRate)
	}
}

func NewTaxRate(rate, threshold float64) *TaxRate {
	if rate == 0 {
		rate = defaultTaxRate
	}
	if threshold < minthreshold {
		threshold = minthreshold
	}
	return &TaxRate{rate, threshold}
}

func (TaxRate *TaxRate) CalTax(p *Product) (price float64) {
	if p.price > TaxRate.threshold {
		price = p.price + (p.price * TaxRate.rate)
	} else {
		price = p.price
	}
	if max, ok := categoryMaxPrices[p.Category]; ok && price > max {
		price = max
	}
	return
}
