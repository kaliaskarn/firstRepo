package main

import (
	"fmt"
)

func main() {
	inputs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(foo(inputs...))

	fmt.Println(bar(inputs))
}

func foo(num ...int) int {
	t := 0
	for _, v := range num {
		t += v
	}
	return t

}

func bar(num2 []int) int {
	t := 0
	for _, v := range num2 {
		t += v
	}
	return t

}
