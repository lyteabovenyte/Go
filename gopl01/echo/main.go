package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func main() {
	var s string
	sep := " "
	for i := 1; i < len(os.Args); i++ {
		// skip file name --> if we want to access build location,
		// we can access it via os.Args[0]
		s += os.Args[i] + sep
	}
	fmt.Println(s)

	fmt.Println(os.Args[1:])                            // a list of all the arguments
	fmt.Printf("the type of os.Args --> %T\n", os.Args) // the type of os.Args --> []string

	fmt.Printf("echo2 result -->%s\n", echo2())
	fmt.Printf("echo with join --> %s\n", echoWithJoin())

	// ex2:
	for i, ar := range os.Args[1:] {
		fmt.Println(i, ")", ar)
	}

	// ex3:
	testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			slice := []string{"amir", "alaeifar", "ala", "alaeifar", "some", "other", "name"}
			s := ""
			for _, arg := range slice {
				s += arg
			}
		}
		fmt.Println("reporting b.ReportAlloc()")
		b.ReportAllocs()
	})
}

func echo2() string {
	s := " "
	for index, arg := range os.Args[1:] {
		s += strconv.Itoa(index+1) + ")" + arg + " "
	}
	return s
}

func echoWithJoin() string {
	return strings.Join(os.Args[1:], " ")
}

/*
Join function using builder pattern for efficient memory management.
func Join(elems []string, sep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return elems[0]
	}

	var n int
	if len(sep) > 0 {
		if len(sep) >= maxInt/(len(elems)-1) {
			panic("strings: Join output length overflow")
		}
		n += len(sep) * (len(elems) - 1)
	}
	for _, elem := range elems {
		if len(elem) > maxInt-n {
			panic("strings: Join output length overflow")
		}
		n += len(elem)
	}

	var b Builder
	b.Grow(n)
	b.WriteString(elems[0])
	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(s)
	}
	return b.String()
}
*/

// var slice = []string{"amir", "alaeifar"}

// // expermineting the time difference between the two.
// func benchmarkMyEcho2(b *testing.B) {
// 	for i := 1; i <= b.N; i++ {
// 		s := ""
// 		for index, arg := range slice {
// 			s += strconv.Itoa(index) + arg
// 		}
// 	}
// }

// func benchamrkMyEchoJoin(b *testing.B) {
// 	for i := 1; i <= b.N; i++ {
// 		strings.Join(slice, ", ")
// 	}
// }
