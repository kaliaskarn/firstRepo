package main

import "fmt"

func main() {

	f1 := map[string]int{
		"michael":  1,
		"jonathan": 2,
		"marcus":   3,
	}

	d1 := []string{"beer", "whiskey", "cider"}

	p1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first:     "Nigglet",
		friends:   f1,
		favDrinks: d1,
	}
	fmt.Println(p1)
	fmt.Println("==============================")

	for k, v := range p1.friends {
		fmt.Println(p1.first, "- friends -", k, v)
	}

	for _, v := range p1.favDrinks {
		fmt.Println(p1.first, "- drinks -", v)
	}
}
