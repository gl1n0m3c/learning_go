package slices

import "fmt"

func SlicesCopy() {
	slice := []int{1, 2, 3, 4, 5}
	a := make([]int, 3, 6)
	fmt.Println(a)
	copy(a, slice[:3])
	fmt.Println(slice)
	fmt.Println(a)
}
