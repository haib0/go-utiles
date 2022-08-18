package play

import (
	"testing"
	"unsafe"
)

/*
Size and alignment guarantees [https://go.dev/ref/spec#Size_and_alignment_guarantees]
For the numeric types, the following sizes are guaranteed:

type                                 size in bytes

byte, uint8, int8                     1
uint16, int16                         2
uint32, int32, float32                4
uint64, int64, float64, complex64     8
complex128                           16

The following minimal alignment properties are guaranteed:

1. For a variable x of any type: unsafe.Alignof(x) is at least 1.
2. For a variable x of struct type: unsafe.Alignof(x) is the largest of all the values unsafe.Alignof(x.f) for each field f of x, but at least 1.
3. For a variable x of array type: unsafe.Alignof(x) is the same as the alignment of a variable of the array's element type.

A struct or array type has size zero if it contains no fields (or elements, respectively) that have a size greater than zero. Two distinct zero-size variables may have the same address in memory.
*/
func TestAlignment(t *testing.T) {
	s1 := struct {
		a int8
		b int16
		c int32
	}{}

	s2 := struct {
		a int8
		c int32
		b int16
	}{}

	t.Log(unsafe.Sizeof(s1)) // 8
	t.Log(unsafe.Sizeof(s2)) // 12
}

func TestUnsafe(t *testing.T) {
	s1 := struct {
		x int
		y int
	}{}
	s2 := struct {
		x int16
		y int8
	}{}

	t.Log(unsafe.Sizeof(s1)) // 16
	t.Log(unsafe.Sizeof(s2)) // 4, not 3 because of align

	t.Log(unsafe.Alignof(s1)) // 8
	t.Log(unsafe.Alignof(s2)) // 2

	t.Log(unsafe.Offsetof(s1.x)) // 0
	t.Log(unsafe.Offsetof(s1.y)) // 8
	t.Log(unsafe.Offsetof(s2.x)) // 0
	t.Log(unsafe.Offsetof(s2.y)) // 4

	t.Log(unsafe.Sizeof([]struct{}{}))  // 24
	t.Log(unsafe.Alignof([]struct{}{})) // 8

	t.Log(unsafe.Sizeof(struct{}{}))  // 0
	t.Log(unsafe.Alignof(struct{}{})) // 1
}
