package main

import (
	"fmt"
)

func iscapitalized(s string) bool {
	return false // placeholder
}

// equivalent of the main test
func IsCapitalized() {
	fmt.Println(iscapitalized("Hello! How are you?"))
	fmt.Println(iscapitalized("Hello How Are You"))
	fmt.Println(iscapitalized("Whats 4this 100K?"))
	fmt.Println(iscapitalized("Whatsthis4"))
	fmt.Println(iscapitalized("!!!!Whatsthis4"))
	fmt.Println(iscapitalized(""))
}
