package main

import (
	"fmt"
)

func main() {
	m := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("full slice: %#v \n", m)
	fmt.Println("----")

	m = append(m[:5], m[6:]...)
	fmt.Printf("slice without 5: %#v \n ", m)
	fmt.Println("---")
}
