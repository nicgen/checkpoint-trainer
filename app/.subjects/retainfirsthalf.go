package main

import (
	"fmt"
)

func retainfirsthalf(str string) string {
	return " " //placeholder
}

// equivalent of the main test
func RetainFirstHalf() {
	fmt.Println(retainfirsthalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(retainfirsthalf("A"))
	fmt.Println(retainfirsthalf(""))
	fmt.Println(retainfirsthalf("Hello World"))
}
