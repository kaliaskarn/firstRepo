package main

import "fmt"

func main() {
	x := exercise()
	fmt.Println(x(42))

}

func exercise() func(int) int {
	return func(s int) int {
		fmt.Println("my number is:", s)
		return s
	}
}
