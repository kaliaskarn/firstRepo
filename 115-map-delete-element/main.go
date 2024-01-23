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
	mp["black hoes"] = 66

	fmt.Println("================")
	fmt.Println(mp)

	delete(mp, "number of hoes")
	fmt.Println("================")

	for k, v := range mp {
		fmt.Println(k, v)
	}

}
