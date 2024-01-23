package main

import (
	"fmt"
	"log"
)

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x)

	num := 42
	numPtr := &num
	fmt.Printf("Type of numPtr: %T\n", numPtr)
	fmt.Printf("Value of numPtr: %v\n", numPtr)

	fmt.Println(*numPtr)

	*numPtr = 66
	log.Println(num)
}
