package play

import (
	"testing"
)

func TestPointerReference(t *testing.T) {
	var x = 0
	y := &x
	t.Logf("x=%d\ty=%d\t*y=%d", x, y, *y)

	*y = 1
	x = 2
	t.Logf("x=%d\ty=%d\t*y=%d", x, y, *y)
}
