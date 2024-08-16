package main

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

type DispatchNotification struct {
	Customer string
	*Product
	Quantity int
}

var Customers = []string{"amir", "mahkam", "seyf", "arman", "arash"}

func DispatchOrders(channel chan<- DispatchNotification) { // specify a send-only channel. (send to channel)
	rand.Seed(uint64(time.Now().UTC().UnixNano()))
	orderCount := rand.Intn(5) + 5
	fmt.Println("order count:", orderCount)
	for i := 0; i < orderCount; i++ {
		channel <- DispatchNotification{
			Customer: Customers[rand.Intn(len(Customers)-1)],
			Product:  ProductList[rand.Intn(len(ProductList)-1)],
			Quantity: rand.Intn(10),
		}
		// if i == 1 {
		// 	notification := <-channel
		// 	fmt.Println("notification from i == 1:", notification.Customer)
		// }
		time.Sleep(time.Second)
	}
	close(channel)
}

// func receiveDispatcher(channel <-chan DispatchNotification) { // a receiver channle carrying DispatchNotification values. (receive from the channel)
// 	for detail := range channel {
// 		fmt.Println("detials for:", detail.Customer, "quantity:", detail.Quantity, "product:", detail.Product)
// 	}
// 	fmt.Println("channel has been closed.")
// }
