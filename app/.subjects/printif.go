package main

import (
	"fmt"
)

func printif(str string) string {
	return " " //placeholder
}

// equivalent of the main test
func PrintIf() {
	fmt.Print(printif("abcdefz"))
	fmt.Print(printif("abc"))
	fmt.Print(printif(""))
	fmt.Print(printif("14"))
}
