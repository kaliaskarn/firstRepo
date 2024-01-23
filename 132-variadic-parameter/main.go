package main

import (
	"fmt"
)

func main() {
	//all the values in sum are inpnutted in the main function, became the function sum requires an integer inside it, and it was declared
	//by func sum(ii ...int) then it also returns and integer.
	//we assign the variable x to be equal to sum, which is a function that will printout first the slice inputted and also
	//and then the type of the slice
	//then theres a loop where n is declared to be equal to zero and then in a loop ii is the variable name for inttegers that will be
	//inside the sum function, and ii is whatever we inputted inside the "sum()"
	// values of ii, all the integers are ranged over, and each value is added to the value of n which is zero
	//which is equal to zero and then returned
	x := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
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
