package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	for k, v := range m {
		fmt.Printf("index: %v \t key: %v\n", k, v)

	}
}
