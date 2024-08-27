package main

import (
	"fmt"
)

func revconcatalternate(slice1, slice2 []int) []int {
	return []int{1, 2, 3} // placeholder
}

// equivalent of the main test
func RevConcatAlternate() {
	fmt.Println(revconcatalternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(revconcatalternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(revconcatalternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(revconcatalternate([]int{1, 2, 3}, []int{}))
}
