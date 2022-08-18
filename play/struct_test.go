package play

import (
	"testing"
)

func TestStructure(t *testing.T) {
	s := NewS()
	t.Log(s)
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
