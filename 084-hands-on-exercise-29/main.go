package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 6; i++ {
		fmt.Println("---")
		for j := 0; j < 6; j++ {
			fmt.Printf("iteration of i is %v \t and value of j is %v\n", i, j)
		}

	}
}
