package main

import "fmt"

type multiplication interface {
	int | float64
}

func multiply[T multiplication](a, b T) T {
	return a * b
}

func main() {
	fmt.Println(multiply(2, 2))

}
