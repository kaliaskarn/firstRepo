package main

import (
	"fmt"
)

func main() {

	x := foo()
	fmt.Println(x)

	z := prac()
	fmt.Println(z)

	g := bar()
	fmt.Println(g)
}

func foo() int {
	return 69
}

func prac() func() string {
	return func() string {
		return "fuck all my homies"
	}
}

func bar() func() int {
	return func() int {
		return 43
	}
}
