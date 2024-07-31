package main

import "fmt"

func atoibase(s string, base string) int {
	return 0 //placeholder
}

// equivalent of the main test
func AtoiBase() {
	fmt.Println(atoibase("125", "0123456789"))
	fmt.Println(atoibase("1111101", "01"))
	fmt.Println(atoibase("7D", "0123456789ABCDEF"))
	fmt.Println(atoibase("uoi", "choumi"))
	fmt.Println(atoibase("bbbbbab", "-ab"))
}
