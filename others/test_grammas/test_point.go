package test_grammas

import "fmt"

type person struct {
	name string
	age int
}

func (p person) funcWithout()  {
	p.age += 1
	fmt.Println("In func without", p.age)
}

func (p *person)funcWith()  {
	p.age += 1
}

func TestPoint()  {
	p := person{
		name: "jone",
		age:  20,
	}
	p.funcWithout()
	fmt.Println("FuncWithout: ", p.age)
	p.funcWith()
	fmt.Println("FuncWith: ", p.age)
}