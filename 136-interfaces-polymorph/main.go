package main

import (
	"fmt"
)

type example struct {
	first  string
	second int
}

type secretAgent struct {
	example
	ltk bool
}

func (e example) execution() {
	fmt.Printf("My bro's name is %v and he is %v years old\n", e.first, e.second)
}

func (sa secretAgent) execution() {
	fmt.Printf("I am a secret agent. Name: %v, Age: %v, License to Kill: %v\n", sa.first, sa.second, sa.ltk)
}

type Interfacing interface {
	execution()
}

func exectutingInterface(inter Interfacing) {
	inter.execution()
}

type gayExample struct {
	first  string
	second int
}

func (gay gayExample) gayMethod() {
	fmt.Println(gay.first, gay.second)
}

type gayest interface {
	gayMethod()
}

func executingGaymethod(gayyy gayest) {
	gayyy.gayMethod()
}

func main() {
	e1 := example{
		first:  "Nurik",
		second: 25, // Assuming an age for e1
	}

	sa1 := secretAgent{
		example: example{
			first:  "Nigga Bond",
			second: 35,
		},
		ltk: true,
	}

	gay1 := gayExample{
		first:  "this nigga is a gay test",
		second: 32,
	}

	//e1.execution()
	//sa1.execution()
	exectutingInterface(e1)
	exectutingInterface(sa1)
	executingGaymethod(gay1)
}
