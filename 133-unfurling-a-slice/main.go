package main

import (
	"fmt"
)

func main() {
	i := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	x := sum(i...)
	fmt.Println("the sum is", x)

}

func sum(ii ...int) int {
	fmt.Println(ii)
	fmt.Printf("%T\n", ii)
	fmt.Println("=======================")

	n := 0
	for _, v := range ii {
		n += v
	}
	return n
}
