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
	fmt.Println(mp)
	fmt.Printf("number of hoes who came: %v\n", mp["number of hoes"])

	ma := make(map[string]int)
	ma["bro"] = 88
	ma["gays"] = 77
	ma["racists"] = 66
	ma["homosexuals"] = 1

}
