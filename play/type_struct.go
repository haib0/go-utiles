package play

type FuncA func(x int) int

func (f FuncA) Handle(x int) int {
	return f(x)
}

type FuncB struct {
	f func(x int) int
}

func (f FuncB) Handle(x int) int {
	return f.f(x)
}
