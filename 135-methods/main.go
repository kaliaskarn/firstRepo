package main

import (
	"fmt"
)

type example struct {
	first  string
	second int
}

func (e example) execution() {
	fmt.Printf("My bro's name is %v and he is %v years old\n", e.first, e.second)
}

func main() {
	e1 := example{
		first:  "Nurik",
		second: 25, // Assuming an age for e1
	}

	e2 := example{
		first:  "AnotherName", // Provide a name for e2
		second: 26,
	}

	e1.execution()
	e2.execution()
}
