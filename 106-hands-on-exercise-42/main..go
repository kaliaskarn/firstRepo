package main

import (
	"fmt"
)

func main() {
	ar := [5]int{10, 11, 12, 13, 14}

	/*another way to assign values is


	ar := [5]int{}

	if i := 0; i < 5; i++{
		x[i]=i
	}


	*/
	fmt.Println(ar)
	fmt.Println("--------")

	for k, v := range ar {
		fmt.Printf("index: %v, type: %T, and value: %v\n", k, v, v)
	}
}
