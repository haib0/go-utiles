package main

import (
	"fmt"
)

type IntSlice []int

func (s IntSlice) change(i int) {
	s[i] = -1
}

type S1 struct {
	Val int
}

func (s *S1) pchange(i int) {
	s.Val = i
}

func (s S1) change(i int) {
	// ineffective assignment to field S1.Val
	// s.Val = i
}

func RunSlicePoint() {
	intSlice := make([]int, 10)
	IntSlice(intSlice).change(1)
	fmt.Println(intSlice)

	s := S1{Val: 1}
	fmt.Println(s)
	s.change(1)
	fmt.Println(s)
	s.pchange(2)
	fmt.Println(s)
}
