package grammas

import "fmt"

type Weekday int

const (
	Mon Weekday = iota
	Tue
	Wen
)

func RunIota() {
	fmt.Println(Mon + 1)
}
