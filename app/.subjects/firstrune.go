package main

import "github.com/01-edu/z01"

func firstrune(s string) rune {
	return '0' //placeholder
}

// equivalent of the main test
func FirstRune() {
	z01.PrintRune(firstrune("Hello!"))
	z01.PrintRune(firstrune("Salut!"))
	z01.PrintRune(firstrune("Ola!"))
	z01.PrintRune('\n')
}
