package main

import (
	"fmt"
)

func main() {
	fmt.Println("text")
	defer fmt.Println("defer 1")

	fmt.Println("text ")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")

	fmt.Println("text")

	defer fmt.Println("defer 123")

}
