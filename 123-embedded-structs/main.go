package main

import (
	"fmt"
)

func main() {

	type person struct {
		first    string
		last     string
		age      int
		location string
	}

	type secretAgent struct {
		person
		ptk bool
	}

	sa1 := secretAgent{ // Use ':=' for declaration
		person: person{ // Correcting the nested struct initialization
			first:    "James",
			last:     "Bond",
			age:      44,
			location: "Great Britain",
		},
		ptk: true, // Added missing comma
	}
	fmt.Println(sa1)
}
