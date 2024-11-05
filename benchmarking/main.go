package benchmarking

import (
	"fmt"
	"runtime"
)

func main() {
	PrintAlloc()
}

func PrintAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println(m.Alloc/1024, "KB")
}
