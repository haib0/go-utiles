package play

import (
	"context"
	"testing"

	"golang.design/x/clipboard"
)

func TestClipboard(t *testing.T) {
	ch := clipboard.Watch(context.TODO(), clipboard.FmtText)
	for data := range ch {
		// print out clipboard data whenever it is changed
		t.Log(string(data))
	}
}
