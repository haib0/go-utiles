package main

import "fmt"

type emptyCtx int

func (e *emptyCtx) string() string {
	switch e {
	case todo:
		return "context.todo"
	case background:
		return "context.background"
	}

	return "NIL"
}

var (
	todo       = new(emptyCtx)
	background = new(emptyCtx)
)

func RunSwitch() {
	fmt.Println(todo.string())
}
