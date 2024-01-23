package main

import (
	"fmt"
)

func main() {
	a := make([]int, 0, 10)
	fmt.Println(cap(a))
	fmt.Println(len(a))
	fmt.Println("---------------")

	a = append(a, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(a)
	fmt.Println(cap(a))
	fmt.Println(len(a))
	fmt.Println("---------------")

}
