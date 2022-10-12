package play

import "testing"

func TestNilSlice(t *testing.T) {
	var s1 []int
	var s2 []int = nil
	var s3 = []int{}

	t.Logf("s1==nil? %v", s1 == nil) // true
	t.Logf("s2==nil? %v", s2 == nil) // true
	t.Logf("s3==nil? %v", s3 == nil) // false
}
