package main

import (
	"fmt"
	"piscine"
)

func revconcatalternate(slice1,slice2 []int) []int {
	return [3 6 2 5 1 4] // placeholder
}

// equivalent of the main test
func RevConcatAlternate() {
	fmt.Println(revconcatalternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(revconcatalternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(revconcatalternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(revconcatalternate([]int{1, 2, 3}, []int{}))
}
