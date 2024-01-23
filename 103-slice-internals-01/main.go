package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	b := a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("---------------")

	a[5] = 666

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("---------------")

}
