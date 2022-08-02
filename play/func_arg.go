package main

import "fmt"

func changeSlice(s []int, i int) {
	s[i] = -1
}

func RunFuncArg() {
	s := make([]int, 2)
	changeSlice(s, 1) // s changed
	fmt.Println(s)
}
