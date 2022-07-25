package grammas

import "fmt"

func RunSlice() {
	a := [5]int{0, 1, 2, 3}

	u := make([]int, 2, 6)
	fmt.Printf("a: %v, u (%d/%d): %v\n", a, len(u), cap(u), u)

	u = a[1:3]

	fmt.Printf("a: %v, u (%d/%d): %v\n", a, len(u), cap(u), u)

	u = append(u, 4)
	fmt.Printf("a: %v, u (%d/%d): %v\n", a, len(u), cap(u), u)

	u = append(u, 5)
	fmt.Printf("a: %v, u (%d/%d): %v\n", a, len(u), cap(u), u)

	u = append(u, 6)
	fmt.Printf("a: %v, u (%d/%d): %v\n", a, len(u), cap(u), u)

	u = append(u, 7)
	fmt.Printf("a: %v, u (%d/%d): %v\n", a, len(u), cap(u), u)
}
