package main

import (
	"fmt"
)

func main() {
	m := []int{22, 23, 24, 25, 26}
	fmt.Println(m)
	fmt.Println("-----")

	m = append(m, 27, 28, 29, 30, 31)
	fmt.Println(m)
	fmt.Println("----")
}
