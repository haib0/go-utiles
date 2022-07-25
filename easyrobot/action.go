package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

var cmdKey = []rune{41, 33, 64, 35}

type Action struct {
	events []hook.Event
	stop   bool
}

func NewAction() *Action {
	return &Action{
		events: []hook.Event{},
		stop:   false,
	}
}

func (a *Action) End() {
	a.stop = true
}

func (a *Action) Reset() {
	a.events = []hook.Event{}
	a.stop = false
}

func (a *Action) record(e hook.Event) {
	for _, c := range cmdKey {
		if e.Keychar == c {
			return
		}
	}
	fmt.Println("Record: ", e)
	a.events = append(a.events, e)
}

// func (a *Action) pop() {
// 	if len(a.events) > 0 {
// 		a.events = a.events[0 : len(a.events)-1]
// 	}
// }

func (a *Action) Print() {
	for _, e := range a.events {
		fmt.Println(e)
	}
}

func (a *Action) Act() {
	a.stop = false
	go a.ActForever()
}

func (a *Action) ActForever() {
	for {
		if a.stop {
			a.stop = false
			fmt.Println("INFO: Stop Successfully")
			return
		}
		a.ActOnce()
	}
}

func (a *Action) ActOnce() {
	events := a.events
	if len(events) == 0 {
		return
	}
	lastWhen := events[0].When
	for _, e := range events {
		if a.stop {
			fmt.Println("INFO: Stop One Event")
			return
		}
		d := e.When.Sub(lastWhen)
		lastWhen = e.When
		time.Sleep(d)

		switch e.Kind {
		case hook.MouseDown:
			{
				button := ""
				if e.Button == 1 {
					button = "left"
				} else if e.Button == 2 {
					button = "right"
				} else {
					button = "wheelLeft"
				}
				fmt.Println("Mouse Click", int(e.X), int(e.Y), button)
				robotgo.MoveClick(int(e.X), int(e.Y), button)
			}
		case hook.KeyDown:
			{
				key := string(e.Keychar)
				fmt.Println("Key Press", key)
				robotgo.KeyPress(key)
			}
		}
	}
}
