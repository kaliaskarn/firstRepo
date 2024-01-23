package main

import (
	"fmt"
)

func main() {
	x := 40
	y := 20

	fmt.Printf("the value of x= %v and value of y= %v\n", x, y)

	switch {
	case x > 20:
		fmt.Println("x is more than 20")
	case x > 50:
		fmt.Println("x is not more than 50")
	case x == 40:
		fmt.Println("x is equal to 40")
	}

	switch y {
	case 22:
		fmt.Println("y is 22")
	case 51:
		fmt.Println("y is 51")
	case 20:
		fmt.Println("y is 20")
	}

	switch y {
	case 20:
		fmt.Println("y is 20")
		fallthrough
	case 21:
		fmt.Println("y is not 21 but printed because of fallthrough")
		fallthrough
	case 22:
		fmt.Println("y is not 22 but printed because of fallthrough")
		fallthrough
	case 23:
		fmt.Println("y is not 23 but printed because of fallthrough")
		fallthrough
	case 24:
		fmt.Println("y is not 24 but printed because of fallthrough")
	}

}
