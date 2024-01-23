package main

import (
	"fmt"
)

func main() {
	fmt.Println(factorial(4))
	fmt.Println(loop(4))

}

func factorial(n int) int {
	fmt.Println("this is a factorial ", n)
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func loop(n int) int {
	total := 1
	for n > 0 {
		total *= n
		n--
	}
	return total
}
