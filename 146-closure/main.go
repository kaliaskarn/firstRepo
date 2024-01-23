package main

import "fmt"

func main() {
	p := practicing()
	fmt.Println(p())
	fmt.Println(p())
	fmt.Println(p())
	fmt.Println(p())
	fmt.Println(p())
	fmt.Println(p())
	fmt.Println(p())

}

func practicing() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
