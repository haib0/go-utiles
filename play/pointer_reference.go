package main

import "fmt"

func ComparePointerReference() {
	var x = 0
	y := &x
	fmt.Println(y)
}
