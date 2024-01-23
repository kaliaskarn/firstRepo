package main

import "fmt"

var x = 50
var y = 51

func main() {
	if x > 20 && y <= 100 {
		fmt.Println("both are right")
	}

	if x > 20 && y > 100 {
		fmt.Println("both are right")
	} else {
		fmt.Println("one of them is wrong")
	}

	if x == 40 || y >= 20 {
		fmt.Println("one of them is right")
	}
}
