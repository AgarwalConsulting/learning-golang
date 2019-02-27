package test

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkSplit(b *testing.B) {
	var split []string
	for i := 0; i < b.N; i++ {
		split = strings.Split("one,two,three,four,five,six,seven", ",")
	}

	fmt.Println(split)
}
