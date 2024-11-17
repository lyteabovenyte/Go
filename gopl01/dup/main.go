package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "" {
			break
		} else {
			count[input.Text()]++
		}
	}
	fmt.Println(count)
	for line, n := range count {
		if n > 1 {
			fmt.Printf("%s\t%d\n", line, n)
		}
	}
	// sigc := make(chan os.Signal, 1)
	// signal.Notify(sigc,
	// 	os.Interrupt,
	// 	syscall.SIGTERM,
	// 	syscall.SIGQUIT,
	// 	syscall.SIGINT,
	// )
	// go func() {
	// 	s := <-sigc
	// 	switch s {
	// 	case os.Interrupt:
	// 		for line, n := range count {
	// 			fmt.Printf("%s\t%d\n", line, n)
	// 		}
	// 	default:
	// 		fmt.Println("out.")
	// 	}
	// }()
}
