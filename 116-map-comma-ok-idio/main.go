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

	if v, ok := mp["number of hoes"]; ok {
		fmt.Println("hoes are present", v)
	} else {
		fmt.Println("PANIC! NIGGAS HOES ARE GONEEEE!!! CRYY BITCHES")
	}
}
