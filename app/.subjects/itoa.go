package main

import "fmt"

func itoa(n int) string {
	return " " //placeholder
}

// equivalent of the main test
func Itoa() {
	fmt.Println(itoa(12345))     // "12345"
	fmt.Println(itoa(0))         // "0"
	fmt.Println(itoa(-1234))     // "-1234"
	fmt.Println(itoa(987654321)) // "987654321"
}
