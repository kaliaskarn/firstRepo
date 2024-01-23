package main

import (
	"fmt"
)

func main() {
	e := foo(
		exercise{
			first:  9,
			second: 10,
		},
	)
	fmt.Println(e)

	i, s := bar()
	fmt.Println(i, s)

}

type exercise struct {
	first  int
	second int
}

func foo(e exercise) int {
	sum := e.first + e.second
	return sum
}

func bar() (int, string) {
	return 43, "Nigger"

}
