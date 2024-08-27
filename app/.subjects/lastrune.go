package main

import "github.com/01-edu/z01"

func lastrune(s string) rune {
	return '0' //placeholder
}

func LastRune() {
	z01.PrintRune(lastrune("Hello!"))
	z01.PrintRune(lastrune("Salut!"))
	z01.PrintRune(lastrune("Ola!"))
	z01.PrintRune(lastrune("Ola!♥"))
	z01.PrintRune('\n')
}
