package play

import (
	"testing"
)

func TestDeadlock(t *testing.T) {
	ch := make(chan bool)
	ch <- true
	<-ch
}
