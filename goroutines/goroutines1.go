package goroutines

import (
	"fmt"
)

func Goroutines1() {
	// fmt.Println(runtime.NumCPU())

	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)

	defer func() {
		panicValue := recover()
		fmt.Println(panicValue)
	}()

	makePanic()
	fmt.Println("yeah!")
}

func makePanic() {
	panic("some panic")
}
