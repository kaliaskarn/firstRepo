package main

import (
	"fmt"
)

func main() {
	ar := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for k, v := range ar {
		fmt.Printf("index value type: %v, %v, %T\n", k, v, v)
	}
}
