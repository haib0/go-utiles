package test_grammas

import "fmt"

func ComparePointerReference()  {
	var x = 0
	y := &x
	fmt.Println(y)
}

type s struct {
	i int
	n int
}

func funcPointer(i *int)  {

}

func funcReference()  {

}