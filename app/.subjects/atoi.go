package main

import (
	"fmt"
)

func atoi(s string) int {
	return 0 // placeholder
}

// equivalent of the main test
func Atoi() {
	fmt.Println(atoi("12345"))         // 12345
	fmt.Println(atoi("0000000012345")) // 12345
	fmt.Println(atoi("012 345"))       // 0
	fmt.Println(atoi("Hello World!"))  // 0
	fmt.Println(atoi("+1234"))         // 1234
	fmt.Println(atoi("-1234"))         // -1234
	fmt.Println(atoi("++1234"))        // 0
	fmt.Println(atoi("--1234"))        // 0
}
