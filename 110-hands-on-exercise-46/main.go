package main

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)
	fmt.Println("--------------")

	x = append(x[:3], x[7:]...)
	fmt.Println(x)

}
