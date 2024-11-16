package main

import (
	"strings"
	"testing"
)

var args = []string{"amir", "alaeifar", "ala", "alaeifar", "mahkam", "some", "other", "names"}

func traditionalConcat(args []string) string {
	s := " "
	for _, arg := range args {
		s += arg
	}
	return s
}

func BenchmarkTraditionalConcat(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		traditionalConcat(args)
	}
}

func BenchmarkPerformantConcat(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		strings.Join(args, " ") // using builder pattern under the hood.
	}
}
