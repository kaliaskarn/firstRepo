package main

import "fmt"

var x = 50
var y = 20

func main() {
	fmt.Println("testing if functions and statements and etc")

	if x >= 20 {
		fmt.Println("this number is larger than 20")
	}

	if x == 49 {
		fmt.Println("this number isnt equal to 49")
	} else {
		fmt.Println("x is actually equal to 50")
	}

	if y == 30 {
		fmt.Println("this is wrong")
	} else if y == 20 {
		fmt.Println("y is actually 20")
	}
}
