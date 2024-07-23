package main

import (
	"fmt"
)

func printifnot(str string) string {
	return " " //placeholder
}

// equivalent of the main test
func PrintIfNot() {
	fmt.Print(printifnot("abcdefz"))
	fmt.Print(printifnot("abc"))
	fmt.Print(printifnot(""))
	fmt.Print(printifnot("14"))
}
