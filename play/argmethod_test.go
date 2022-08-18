package play

import (
	"testing"
)

type person struct {
	name string
	age  int
}

func (p person) funcWithout() {
	p.age += 1
	_ = p.age
}

func (p *person) funcWith() {
	p.age += 2
}

func TestPoint(t *testing.T) {
	p := person{
		name: "jone",
		age:  20,
	}

	p.funcWithout()
	t.Log("FuncWithout: ", p.age)

	p.funcWith()
	t.Log("FuncWith: ", p.age)
}
