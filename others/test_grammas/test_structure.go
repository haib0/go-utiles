package test_grammas

import "fmt"

func TestStructure()  {
	s := NewS()
	fmt.Println(s)
}

type S struct {
	i int
	N string
	m string
}

func NewS() S {
	s := S{i: 1, N: "1", m: "un"}
	return s
}