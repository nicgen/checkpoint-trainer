package main

import (
	"fmt"
)

func digitlen(n, base int) int {
	return 0 //placeholder
}

// equivalent of the main test
func DigitLen() {
	fmt.Println(digitlen(100, 10))
	fmt.Println(digitlen(100, 2))
	fmt.Println(digitlen(-100, 16))
	fmt.Println(digitlen(100, -1))
}
