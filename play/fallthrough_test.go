package play

import "testing"

// Test "fallthrough"
//
// Output:
//
//	fallthrough_test.go:12: x is less than 3
//	fallthrough_test.go:15: x is less than 4
//	fallthrough_test.go:18: x is less than 5
//	fallthrough_test.go:21: DEFAULT
func TestFallthrough(t *testing.T) {
	x := 2
	switch x {
	case 1:
		t.Log("x is less than 2")
		fallthrough
	case 2:
		t.Log("x is less than 3")
		fallthrough
	case 3:
		t.Log("x is less than 4")
		fallthrough
	case 4:
		t.Log("x is less than 5")
		fallthrough
	default:
		t.Log("DEFAULT")
	}
}
