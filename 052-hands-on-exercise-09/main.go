package main

import "fmt"

var a = 11

const b = 12

func main() {
	fmt.Printf("the value of a is %v and the type of it is %T\n", a, a)
	fmt.Printf("the value of b is %v and the type of it is %T\n", b, b)
	c := 13
	fmt.Printf("the value of c is %v and the type of it is %T\n", c, c)
}
