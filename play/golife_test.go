package play

import (
	"testing"
	"time"
)

func gof1() {
	time.Sleep(3 * time.Second)
	print("gof1 still alive\n")
}

func gof2() {
	print("gof2 a\n")
	go gof1() // won't die
	print("gof2 b\n")
}

func TestGolife(t *testing.T) {
	go gof2()
	time.Sleep(5 * time.Second)
}
