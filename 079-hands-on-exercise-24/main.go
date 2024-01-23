package main

import (
	"fmt"
	"math/rand"
)

/*func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("i = %v\t counting from 0 to 99\n", i)
	}

}
*/

func main() {
	for i := 0; i < 100; i++ {
		x := rand.Intn(10)
		y := rand.Intn(10)

		fmt.Printf("iteration is %v \t and x and y is equal to %v and %v \t", i, x, y)

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
}
