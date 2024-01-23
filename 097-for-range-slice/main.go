package main

import (
	"fmt"
)

func main() {
	a := []string{
		"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)"}

	fmt.Println(len(a))
	fmt.Printf("%T \n", a)

	for k, i := range a {
		fmt.Printf("index is %v \t, and the string is %v \n", k, i)
	}

	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])
	fmt.Println(a[3])

	fmt.Println(len(a))

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

}
