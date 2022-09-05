package play

import (
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	x := []int{5, 4, 3, 2, 1}
	sort.Ints(x[2:])
	t.Log(x)
}
