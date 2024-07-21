package main

func reduceint(a []int, f func(int, int) int) {
}

// equivalent of the main test
func ReduceInt(a []int, f func(int, int) int) {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2}
	reduceint(as, mul)
	reduceint(as, sum)
	reduceint(as, div)
}
