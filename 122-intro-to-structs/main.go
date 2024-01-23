package main

import (
	"fmt"
)

func main() {

	type boys struct {
		first    string
		last     string
		age      int
		location string
	}

	b1 := boys{
		first:    "Nurlan",
		last:     "Kaliaskar",
		age:      20,
		location: "Astana/Nur-Sultan",
	}

	fmt.Println(b1)
	fmt.Println(b1.last, b1.location)
	fmt.Printf("Type: %T, Value: %v\n", b1, b1)
}
