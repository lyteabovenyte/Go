package main

import (
	"fmt"
)

// type CategoryError struct {
// 	requestedCategory string
// }

// // implementing error interface for CategoryError type.
// func (e *CategoryError) Error() string {
// 	return "category " + e.requestedCategory + " does not exist."
// }

type channelMessage struct {
	Category string
	Total    float64
	err      error
}

// defining a method for an alias type.
func (slice ProductSlice) TotalPrice(category string) (total float64, err error) {
	productCount := 0

	for _, p := range slice {
		if p.Category == category {
			total += p.Price
			productCount++
		}
	}
	if productCount == 0 {
		// err = &CategoryError{requestedCategory: category}
		err = fmt.Errorf("category %v not found", category)
	}
	return total, err
}

func (slice ProductSlice) TotalPriceAsync(categories []string, channel chan<- channelMessage) {
	for _, cat := range categories {
		total, err := slice.TotalPrice(cat)
		channel <- channelMessage{
			Category: cat,
			Total:    total,
			err:      err,
		}
	}
	close(channel)
}
