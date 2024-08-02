package main

import (
	"fmt"
	"strconv"
)

func main() {
	val1 := 100
	val2 := "amir"

	f1 := strconv.Itoa(val1)
	fmt.Printf("the type is %T and the value is %v", f1, f1)
	if f2, err := strconv.Atoi(val2); err == nil {
		fmt.Printf("the type is %T and the value is %v", f2, f2)
	} else {
		fmt.Println(err)
	}

}