// package main in string-100 determine using string

// rune, string, unicode, encoding, UTF-8
// iterating over string
// a string being a sequence of byte
// a rune being a unicode code point determining a single value
// (character) --> unicode code point
package main

import (
	"fmt"
	"runtime"
	"strings"
)

func benchmark() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%dKB\n", m.Alloc/1024)
}

func stringfrombyte() {
	s := string([]byte{0xE6, 0xB1, 0x89})
	fmt.Println(s)
}

func stringToRune(s string) {
	r := []rune(s)
	for i, v := range r {
		fmt.Printf("%d position. %c char\n", i, v)
	}
}

func stringToRune2(s string) {
	for i, v := range s {
		fmt.Printf("%d pos, %c char\n", i, v)
	}
}

func stringsTrim(s string, remove string) {
	fmt.Printf(
		"trim right %s\ntrim left %s\ntrim suffix %s\ntrim prefix %s\n",
		strings.TrimRight(s, remove),
		strings.TrimLeft(s, remove),
		strings.TrimSuffix(s, remove),
		strings.TrimPrefix(s, remove),
	)
}

func concatWithoutGrow(values []string) string {
	sb := strings.Builder{}
	for _, v := range values {
		_, _ = sb.WriteString(v) // appends a string to the builder struct
	}
	return sb.String()
}

func concatWithGrow(values []string) string {
	sb := strings.Builder{}
	sb.Grow(10)
	for _, v := range values {
		_, _ = sb.WriteString(v)
	}
	return sb.String()
}

func main() {
	stringfrombyte() // 汉
	stringToRune("hållo汉")
	stringToRune2("hållo汉") // be careful about the position
	stringsTrim("xoxo123xoxo", "xo")

	runtime.GC()
	s := []string{"amir", "alaeifar"}
	fmt.Println(concatWithoutGrow(s))
	benchmark()
	runtime.GC()
	d := []string{"amir", "alaeifar"}
	fmt.Println(concatWithGrow(d))
	benchmark()

	var builder strings.Builder
	for i := 3; i >= 1; i-- {
		fmt.Fprintf(&builder, "%d,", i)
	}
	builder.WriteString("finish")
	fmt.Println(builder.String())

	str := "amir"
	func(s string) {
		n := strings.TrimRight(s, "r")
		fmt.Println(n)
	}(str)
	fmt.Println(str)
}
