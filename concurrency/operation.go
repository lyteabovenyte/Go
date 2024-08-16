package main

import (
	"fmt"
	"time"
)

func CalcStoreTotal(data ProductData) float64 {
	var storeTotal float64
	var channel chan float64 = make(chan float64)

	for category, group := range data {
		go group.TotalPrice(category, channel) // puts the output into the channel
	}
	time.Sleep(time.Second * 5)
	for i := 0; i < len(data); i++ {
		// receiving from channels is a blocking operation.
		fmt.Println("channel is receiving...")
		value := <-channel // gets the channel's data which is filled by the separated goroutines.
		fmt.Println("channel received the value:", value)
		storeTotal += value
		time.Sleep(time.Second)

	}
	fmt.Println("total:", ToCurrency(storeTotal))
	return storeTotal
}

func (group ProductGroup) TotalPrice(category string, resultChannel chan float64) {
	var total float64
	for _, p := range group {
		// fmt.Println(category, "product:", p.Name)
		total += p.Price
		time.Sleep(time.Second)
	}
	fmt.Println(category, "subtotal:", ToCurrency(total))
	fmt.Println("channel sending...")
	resultChannel <- total
	fmt.Println("channel sending compeleted...")
}
