package main

import (
	"fmt"
)

func main() {
	messageCh := make(chan int, 10)
	disconnectCh := make(chan struct{}, 1)

	for i := 0; i < 10; i++ {
		messageCh <- i
	}
	disconnectCh <- struct{}{}

	// get all the messages from messageCh
	// then disconnect respectively
	for {
		select {
		case v := <-messageCh:
			fmt.Println(v)
		case <-disconnectCh:
			for {
				select {
				case v := <-messageCh:
					fmt.Println(v)
				default:
					fmt.Println("disconnect signal, return")
					return
				}
			}
		}
	}
}
