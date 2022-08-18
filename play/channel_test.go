package play

import (
	"fmt"
	"testing"
)

func loopp(c chan int, name string) {
	for x := range c {
		fmt.Println(name, x)
	}
}

func TestChannel(t *testing.T) {
	// c1 := make(chan int)
	c2 := make(chan int, 10)
	// go loopp(c1, "c1")
	go loopp(c2, "c2")

	for i := 0; i < 10; i++ {
		// c1 <- i
		c2 <- i
	}
}
