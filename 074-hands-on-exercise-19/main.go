package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	fmt.Printf("the value of x is %v\n", x)

	if x > 0 && x < 100 {
		fmt.Println("x is between 0 and 100")
	} else if x > 101 && x < 200 {
		fmt.Println("x is between 101 and 200")
	} else if x > 201 && x < 250 {
		fmt.Println("x is between 201 and 250")
	}
}
