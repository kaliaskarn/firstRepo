package main

import (
	"fmt"
)

func main() {

	a := [3]int{1, 2, 3}
	fmt.Println(a)

	b := [...]string{"niggas", "obama", "kanye west"}
	fmt.Println(b)

	var c [3]int
	c[0] = 7
	c[1] = 3
	fmt.Println(c)

}
