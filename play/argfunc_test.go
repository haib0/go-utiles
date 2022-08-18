package play

import "testing"

func changeSlice(s []int, i int) {
	s[i] = -1
}

func TestFuncArg(t *testing.T) {
	s := make([]int, 2)
	t.Log("Before: ", s)
	changeSlice(s, 1) // s changed
	t.Log("After: ", s)
}
