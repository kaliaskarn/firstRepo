package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Printf("x is equal to %v and y is equal to %v\n", x, y)

	if x < 4 && y < 4 {
		fmt.Println("both are less than 4")
	} else if x > 6 && y > 6 {
		fmt.Println("both are larger than 6")
	} else if x >= 4 && x <= 6 {
		fmt.Println("both are larger than 6")
	} else if y != 5 {
		fmt.Println("y is not 5")
	} else {
		fmt.Println("nont of them cases were met")
	}

	switch {
	case x < 4 && y < 4:
		fmt.Println("both are less than 4")
	case x > 6 && y > 4:
		fmt.Println("both are larger than 6")
	case x >= 4 && x <= 6:
		fmt.Println("x is greater than or equal to 4 and equal to or less than 6")
	case y != 5:
		fmt.Println("y isnt 5 ")
	default:
		fmt.Println("none are right")

	}
}
