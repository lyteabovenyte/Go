package main

import (
	"fmt"
	"time"
	"unsafe"
)

func main() {
	fmt.Println("----------buffered channel------------")
	bufferedChannel()
	fmt.Println("----------unbuffered channel-----------")
	unBufferedChannel()
	fmt.Println("--------checking unsafe size of-------")
	var m struct{}
	fmt.Println(unsafe.Sizeof(m))

	fmt.Println("------checking whether a channel is open or not-------")
	// examining whether we are getting an actual message
	// or a closure event
	ch1 := make(chan int)
	close(ch1)
	v, open := <-ch1
	fmt.Printf("v: %v, is-open: %v\n", v, open)

	fmt.Println("-----testing merge function--------")
	channel1 := make(chan int, 2)
	channel2 := make(chan int, 2)

	finalCh := merge(channel1, channel2)
	channel1 <- 1
	channel1 <- 10
	channel2 <- 2
	close(channel1)
	channel2 <- 3
	channel2 <- 4
	channel2 <- 5
	channel2 <- 6
	channel2 <- 7
	close(channel2)

	for {
		value, open := <-finalCh
		if !open {
			fmt.Println("channel is closed")
			break
		}
		fmt.Println(value)
	}

	fmt.Println("------data race with slices-------")
	s := make([]int, 0, 1)
	fmt.Println("s", s)
	go func() {
		s1 := append(s, 5)
		fmt.Println("s1", s1)
	}()
	time.Sleep(time.Second * 1)
	go func() {
		s2 := append(s, 6)
		fmt.Println("s2", s2)
	}()
	fmt.Println("s again: ", s)
	time.Sleep(time.Second * 1)
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
		// will block until receiver is ready
		// so we'll get all the uMessageCh messages
		// to run this loop
		uMessageCh <- i
	}
	disconnetCh <- struct{}{}
}

// merge two channels
// state machine and nil channel approach
// Receiving from both channels. If one is closed,
// we assign it to nil so we only receive from one channel
func merge(ch1, ch2 chan int) (MergeCh chan int) {
	MergeCh = make(chan int, 5)

	go func() {
		for ch1 != nil || ch2 != nil {
			select {
			case v, open := <-ch1:
				if !open {
					ch1 = nil
					break // it will break the ongoing select not the for loop
				}
				MergeCh <- v
			case v, open := <-ch2:
				if !open {
					ch2 = nil
					break
				}
				MergeCh <- v
			}
		}
		close(MergeCh)
	}()
	return MergeCh
}
