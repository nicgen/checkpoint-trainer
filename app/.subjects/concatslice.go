package main

import (
	"fmt"
)

func concatslice(slice1, slice2 []int) []int {
	return [1 2 3 4 5 6]int // placeholder
}

// equivalent of the main test
func ConcatSlice() {
	fmt.Println(concatslice([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(concatslice([]int{}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(concatslice([]int{1, 2, 3}, []int{}))
}
