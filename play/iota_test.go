package play

import (
	"testing"
)

type Weekday int

const (
	Mon Weekday = iota
	Tue
	Wen
)

func TestIota(t *testing.T) {
	t.Log(Mon + 1)
}
