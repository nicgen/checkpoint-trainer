package main

import "fmt"

func split(s, sep string) []string {
	return []string{"placeholder"} //placeholder
}

// equivalent of the main test
func Split() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", split(s, "HA"))
}
