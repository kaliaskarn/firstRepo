package main

import "fmt"

func diabetes() func() int {
	x := 69
	return func() int {
		return x * x
	}
}

func main() {
	x1 := diabetes()
	fmt.Println(x1())
}
