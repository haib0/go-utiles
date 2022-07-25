package main

import (
	"fmt"

	hook "github.com/robotn/gohook"
)

var action = NewAction()
var recording = false
var running = false

func PrintManual() {
	manual := `-----EASY ROBOT----- 
CTRL+SHIFT+1: RECORD MOUSE CLICK AND KEYBOARD INPUT / SAVE RECORD
CTRL+SHIFT+2: PRINT ACTIONS
CTRL+SHIFT+3: DO ACTIONS FOREVER
CTRL+SHIFT+0: END ACTIONS
----EASEYOURHAND----`

	fmt.Println(manual)
}

func main() {
	Regist()
	PrintManual()
	s := hook.Start()
	defer hook.End()
	<-hook.Process(s)
}

func Regist() {
	hook.Register(hook.KeyDown, []string{}, func(e hook.Event) {
		if recording {
			action.record(e)
		}
	})

	hook.Register(hook.MouseDown, []string{}, func(e hook.Event) {
		if recording {
			action.record(e)
		}
	})

	hook.Register(hook.KeyDown, []string{"0", "ctrl", "shift"}, cbEnd)

	hook.Register(hook.KeyDown, []string{"1", "ctrl", "shift"}, cbRecord)

	hook.Register(hook.KeyDown, []string{"2", "ctrl", "shift"}, cbPrint)

	hook.Register(hook.KeyDown, []string{"3", "ctrl", "shift"}, cbAct)

}

func cbEnd(e hook.Event) {
	if recording {
		fmt.Println("!! Bo's LOVE to Fish Never END 💖")
		return
	}
	if running {
		action.End()
		running = false
		fmt.Println("---END ACT---")
		fmt.Println("---PLEASE WAIT FOR SUCCESS---")
		return
	}
	PrintManual()
}

func cbPrint(e hook.Event) {
	if recording {
		fmt.Println("---RECORDING---")
		return
	}
	if running {
		fmt.Println("---RUNING---")
		return
	}
	if !recording {
		fmt.Println("---PRINT EVENTS---")
		action.Print()
		fmt.Println("---PRINT EVENTS OVER---")
	}
}

func cbRecord(e hook.Event) {
	if running {
		fmt.Println("---RUNING---")
		return
	}
	if recording {
		fmt.Println("---SAVE RECORD---")
		recording = false
	} else {
		fmt.Println("---START RECORD---")
		recording = true
		action.Reset()
	}
}

func cbAct(e hook.Event) {
	if recording {
		fmt.Println("---RECORDING---")
		return
	}
	if running {
		fmt.Println("---RUNING---")
		return
	}
	fmt.Println("---ACT---")
	running = true
	action.Act()
}
