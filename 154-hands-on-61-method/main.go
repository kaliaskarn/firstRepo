package main

import (
	"fmt"
)

type person struct {
	first string
	age   int
}

func main() {
	p1 := person{
		first: "Michael NIgger",
		age:   69,
	}

	p1.speak()

}

func (p person) speak() {
	fmt.Println(p.first, "is my real name and my age is ", p.age)
}
