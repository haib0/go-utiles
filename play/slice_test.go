package play

import (
	"testing"
)

func TestSlice(t *testing.T) {
	a := [5]int{0, 1, 2, 3}

	u := make([]int, 2, 6)
	t.Logf("a: %v, u (%d/%d): %v\n", a, len(u), cap(u), u)

	u = a[1:3]

	t.Logf("a: %v, u (%d/%d): %v\n", a, len(u), cap(u), u)

	u = append(u, 4)
	t.Logf("a: %v, u (%d/%d): %v\n", a, len(u), cap(u), u)

	u = append(u, 5)
	t.Logf("a: %v, u (%d/%d): %v\n", a, len(u), cap(u), u)

	u = append(u, 6)
	t.Logf("a: %v, u (%d/%d): %v\n", a, len(u), cap(u), u)

	u = append(u, 7)
	t.Logf("a: %v, u (%d/%d): %v\n", a, len(u), cap(u), u)
}
