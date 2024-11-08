package main

import (
	"fmt"
)

func main() {
	fmt.Println("----------buffered channel------------")
	bufferedChannel()
	fmt.Println("----------unbuffered channel-----------")
	unBufferedChannel()
}

func bufferedChannel() {
	bmessageCh := make(chan int, 10)
	disconnectCh := make(chan struct{}, 1)

	for i := 0; i < 10; i++ {
		bmessageCh <- i
	}
	disconnectCh <- struct{}{}

	// get all the messages from bmessageCh
	// then disconnect respectively
	for {
		select {
		case v := <-bmessageCh:
			fmt.Println(v)
		case <-disconnectCh:
			for {
				select {
				case v := <-bmessageCh:
					fmt.Println(v)
				default:
					fmt.Println("disconnect signal, return")
					return
				}
			}
		}
	}
}

func unBufferedChannel() {
	uMessageCh := make(chan int)
	disconnetCh := make(chan struct{})

	go func() {
		for {
			select {
			case v := <-uMessageCh:
				fmt.Println(v)
			case <-disconnetCh:
				fmt.Println("disconnet signal, return")
				return
			}
		}
	}()

	for i := 0; i < 10; i++ {
		uMessageCh <- i // will block until receiver is ready
	}
	disconnetCh <- struct{}{}
}
