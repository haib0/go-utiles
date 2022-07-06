package others

import (
	"context"

	"golang.design/x/clipboard"
)

func RunClipboard() {
	ch := clipboard.Watch(context.TODO(), clipboard.FmtText)
	for data := range ch {
		// print out clipboard data whenever it is changed
		println(string(data))
	}
}
