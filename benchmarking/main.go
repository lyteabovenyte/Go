package main

import (
	"fmt"
	"runtime"
)

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("before GC", m.Alloc/1024, "KB")
	runtime.GC()
	fmt.Println("after GC", m.Alloc/1024, "KB")
}
