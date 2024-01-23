package main

import "fmt"

func add[T int | float64](a, b T) T {
	return a + b
}

func add1(a int, b int) int {
	return a * b
}

func main() {
	p := add(2, 3)
	fmt.Println(p)

	b := add(42.22, 33.69)
	fmt.Println(b)

	fmt.Println(add1(2, 2))

}
