package main

import (
	"fmt"
)

func main() {

	/*type person struct{
		first    string
		last     string
		age      int
		location string
	}
	*/

	p1 := struct {
		first    string
		last     string
		age      int
		location string
	}{
		first:    "Nigglet",
		last:     "Dark one",
		age:      69,
		location: "Washinghton nigga street",
	}
	fmt.Println(p1)
	fmt.Printf("%T \t %v\n", p1, p1)

}
