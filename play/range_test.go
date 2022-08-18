package play

import (
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {
	rangeMap()
	rangeArray()
}

func rangeMap() {
	var m = make(map[string]int)
	m["A"] = 1
	m["C"] = 3
	m["B"] = 2
	for k, n := range m {
		fmt.Println(k, n)
	}
	fmt.Println()
}

func rangeArray() {
	var a []rune = []rune{'a', 'b', 'c'}
	for i, n := range a {
		fmt.Println(i, string(n))
	}
	fmt.Println()

	var ii int
	var nn rune
	for ii, nn = range a {
		fmt.Println(ii, nn)
	}
	fmt.Println()
}
