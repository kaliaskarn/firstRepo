package main

import (
	"fmt"
)

func main() {
	ar := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(ar[:5])
	fmt.Println(ar[5:])
	fmt.Println(ar[2:7])
	fmt.Println(ar[1:6])
}
