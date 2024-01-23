package main

import (
	"fmt"
)

type person struct {
	first string

	age int
}

func (p *person) speak() {
	fmt.Println(p.first, "just said that her age is", p.age)
}

// func (p person) talk(b string) {
// 	fmt.Println(p.first, "will never meet nurlan even after")
// }

type human interface {
	speak()
	// talk()
}

func saySomething(h human) {
	h.speak()
	// h.talk()

}

func main() {
	p1 := &person{
		first: "Aknur",
		age:   19,
	}
	saySomething(p1)

}
