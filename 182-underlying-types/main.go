package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type multiplication interface {
	constraints.Integer | constraints.Float
}

func multiply[T multiplication](a, b T) T {
	return a * b
}

type myBaby int

func main() {
	var v myBaby
	v = 66
	fmt.Println(multiply(v, 2))

}
