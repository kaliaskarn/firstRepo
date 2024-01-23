package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 43; i++ {

		x := rand.Intn(5)

		fmt.Printf("iteration: %v \t value of x is: %v \t", i, x)

		switch {
		case x >= 2 && x < 5:
			fmt.Println("x is between 2 and four")
		case x >= 0 && x <= 3:
			fmt.Println("x is between 0 and 3")
		case x >= 1 && x <= 2:
			fmt.Println("x is either 1 or 2")
		}
	}

}
