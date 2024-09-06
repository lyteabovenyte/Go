package main

import (
	"fmt"
)

func main() {
	categories := []string{"Watersports", "Soccer", "Running"}

	ch := make(chan channelMessage, 100)

	go Products.TotalPriceAsync(categories, ch)

	// for result := range ch { // not blocking
	// 	if result.err != nil {
	// 		fmt.Println(result.err)
	// 	} else {
	// 		fmt.Println("cat:", result.Category, "total:", ToCurrency(result.Total))
	// 	}
	// }

	// 	for {
	// 		select {
	// 		case p, ok := <-ch:
	// 			if ok {
	// 				if p.err == nil {
	// 					fmt.Println("cat:", p.Category, "total:", p.Total)
	// 				} else {
	// 					fmt.Println(p.err)
	// 				}
	// 			} else {
	// 				fmt.Println("the channel is closed")
	// 				goto alldone
	// 			}
	// 		default:
	// 			fmt.Println("waiting ...")
	// 			time.Sleep(time.Millisecond * 500)
	// 		}
	// 	}
	// alldone:
	// 	fmt.Println("all done!!!")

	defer func() {
		if arg := recover(); arg != nil {
			fmt.Printf("the type is %T, and the value is %v\n", arg, arg)
			if err, ok := arg.(error); ok {
				fmt.Println("recover", "Error", err)
			} else if str, ok := arg.(string); ok {
				fmt.Println("recover", "message:", str)
			} else {
				fmt.Println("recover:", "panic recovered.")
			}
		}
	}()

	for message := range ch {
		if message.err == nil {
			fmt.Println("cat:", message.Category, "total:", message.Total)
		} else {
			fmt.Println("error:", message.err)
		}
	}
}
