package libs

import (
	"fmt"
	"sync"
)

type student struct {
	Name string
	Age  int
}

func JustA() {
	m := make(map[string]*student)
	students := []*student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "Hwang", Age: 22},
	}
	l := sync.RWMutex{}
	wg := sync.WaitGroup{}
	for i, stu := range students {
		wg.Add(1)
		go func() {
			defer wg.Done()
			l.Lock()
			local := stu
			fmt.Println(local)
			m[local.Name] = local
			l.Unlock()
		}()
		fmt.Printf("For Index: %d \n", i)
	}
	wg.Wait()
	fmt.Printf("%v", m)
}
