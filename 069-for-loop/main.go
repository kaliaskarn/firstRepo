package main

import (
	"fmt"
)

var x int = 14
var y int = 5

func main() {
	fmt.Printf("the value of x is %v\n the value of y is %v\n", x, y)

	for i := 0; i < 10; i++ {
		fmt.Printf("counting to five: %v \t first loop made from 0 to 9\n", i)
	}

	for {
		fmt.Printf("the value of y is %v \t my second loop\n", y)
		if y > 15 {
			break
		}
		y++
	}

	for y < 20 {
		fmt.Printf("y is %v \t\t\t my third loop\n", y)
		y++
	}

	for i := 2; i < 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Printf("i= %v \t printing even numbers under 20\n", i)
	}

}
