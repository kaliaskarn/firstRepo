package main

import (
	"fmt"
)

func main() {
	type person struct {
		firstName   string
		lastName    string
		favIceCream []string
	}
	p1 := person{
		firstName:   "Nurlan",
		lastName:    "Kaliaskar",
		favIceCream: []string{"Chocolate", "Milk", "Phistacho"},
	}

	p2 := person{
		firstName:   "Nurali",
		lastName:    "Kaliaskar",
		favIceCream: []string{"Vanilla", "Strawberry", "Mix of fruits"},
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println("========================================")

	for _, v := range p1.favIceCream {
		fmt.Println(p1.firstName, "favorite is", v)
	}

	fmt.Println("============================")
	for _, v := range p2.favIceCream {
		fmt.Println(p2.firstName, "favorite is", v)
	}

}
