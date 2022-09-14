package play

import (
	"encoding/json"
	"testing"
)

func TestMarshal(t *testing.T) {
	type S struct {
		A int
		B *int
		C float64
		d func() string // is lower case
		e chan struct{} // is lower case
	}

	s := S{
		A: 1,
		B: nil,
		C: 12.15,
		d: func() string {
			return "NowCoder"
		},
		e: make(chan struct{}),
	}

	_, err := json.Marshal(s)
	if err != nil {
		t.Error("err occurred")
		return
	}
	t.Log("everything is ok")
}
