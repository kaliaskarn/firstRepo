package main

import "fmt"

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

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, v := range m {
		fmt.Println(v)
		for _, v2 := range v.favIceCream {
			fmt.Println(v.firstName, v.lastName, "loves this", v2)
		}
	}

}
