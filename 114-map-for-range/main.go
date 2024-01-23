package main

import (
	"fmt"
)

func main() {
	mp := map[string]int{
		"number of hoes":    21,
		"number of friends": 22,
		"number of people":  43,
		"black people":      86,
	}

	fmt.Println("================")

	for k, v := range mp {
		fmt.Println(k, v)
	}
	fmt.Println("=================================")

	for _, v := range mp {
		fmt.Println(v)
	}
	fmt.Println("=================================")

	for k := range mp {
		fmt.Println(k)
	}
	fmt.Println("=================================")
}
