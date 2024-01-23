package main

import (
	"fmt"
)

func main() {
	m := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("slice m: %v \n", m)
	fmt.Printf("------")

	fmt.Printf("slice m: %v \n", m[0:4])
	fmt.Printf("------")

	fmt.Printf("slice m: %v \n", m[:8])
	fmt.Printf("------")

	fmt.Printf("slice m: %v \n", m[4:])
	fmt.Printf("------")
}
