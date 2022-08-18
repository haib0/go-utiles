package main

import (
	"fmt"
	"unsafe"
)

func RunUnsafe() {
	s1 := struct {
		x int
		y int
	}{}
	s2 := struct {
		x int16
		y int8
	}{}

	fmt.Println(unsafe.Sizeof(s1)) // 16
	fmt.Println(unsafe.Sizeof(s2)) // 4, not 3 because of align

	fmt.Println(unsafe.Alignof(s1)) // 8
	fmt.Println(unsafe.Alignof(s2)) // 2

	fmt.Println(unsafe.Offsetof(s1.x)) // 0
	fmt.Println(unsafe.Offsetof(s1.y)) // 8
	fmt.Println(unsafe.Offsetof(s2.x)) // 0
	fmt.Println(unsafe.Offsetof(s2.y)) // 4
}
