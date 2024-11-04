package main

import (
	"fmt"
	"runtime"
)

func main() {
	printAlloc()
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println(m.Alloc/1024, "KB")
}
