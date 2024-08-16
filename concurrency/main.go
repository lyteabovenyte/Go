package main

import (
	"fmt"
	"time"
)

// using select to send to a channel.
func enumerateProducts(channel chan<- *Product) {
	for _, p := range ProductList {
		select {
		case channel <- p:
			fmt.Println("product saved.", p.Name)
		default:
			fmt.Println("product discarding...", p.Name)
			time.Sleep(time.Second)
		}
	}
	close(channel)
}

func main() {
	// dispatchChannel := make(chan DispatchNotification, 20)

	// var receiverChannel <-chan DispatchNotification = dispatchChannel
	// var senderChannel chan<- DispatchNotification = dispatchChannel

	// go DispatchOrders(chan<- DispatchNotification(dispatchChannel))
	// receiveDispatcher((<-chan DispatchNotification)(dispatchChannel))

	productChannel := make(chan *Product, 3)
	go enumerateProducts(productChannel)

	time.Sleep(time.Second)

	for p := range productChannel {
		fmt.Println("product Received.", p.Name)
	}

	// openChannels := 2

	//	for {
	//		select {
	//		case detail, ok := <-dispatchChannel:
	//			if ok {
	//				fmt.Println("details for Notification:", detail.Customer)
	//			} else {
	//				fmt.Println("dispatch channel has been closed")
	//				dispatchChannel = nil
	//				openChannels--
	//			}
	//		case product, ok := <-productChannel:
	//			if ok {
	//				fmt.Println("product:", product.Name)
	//			} else {
	//				fmt.Println("product channel has been closed.")
	//				productChannel = nil
	//				openChannels--
	//			}
	//		default:
	//			if openChannels == 0 {
	//				fmt.Println("all values received and channels are closed")
	//				goto alldone
	//			} else {
	//				fmt.Println("waiting for new messages. channels are open...")
	//				time.Sleep(time.Second)
	//			}
	//		}
	//	}
	//
	// alldone:
	//
	//	fmt.Println("All Done!!")
}
