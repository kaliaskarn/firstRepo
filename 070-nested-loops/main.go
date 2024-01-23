package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("--_")
		for j := 0; j < 5; j++ {
			fmt.Printf("%v for outer loop \t %v for inner loop\n", i, j)
		}
	}
}
