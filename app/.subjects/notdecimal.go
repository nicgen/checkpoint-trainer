package main

import (
	"fmt"
)

func notdecimal(dec string) string {
	return "" // placeholder
}

// equivalent of the main test
func NotDecimal() {
	fmt.Print(notdecimal("0.1"))
	fmt.Print(notdecimal("174.2"))
	fmt.Print(notdecimal("0.1255"))
	fmt.Print(notdecimal("1.20525856"))
	fmt.Print(notdecimal("-0.0f00d00"))
	fmt.Print(notdecimal(""))
	fmt.Print(notdecimal("-19.525856"))
	fmt.Print(notdecimal("1952"))
}
