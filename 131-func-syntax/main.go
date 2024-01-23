package main

import (
	"fmt"
)

func main() {
	foo()

	boo("jake")

	s := moo("very very much")
	fmt.Println(s)

	s1, n := soo("Jack Hardmann", 52)
	fmt.Println(s1, n)
}

func foo() {
	fmt.Println("this has no parameters and return")
}

func boo(s string) {
	fmt.Println("my bro has an name and its:", s)
}

func moo(s string) string {
	return fmt.Sprint("my homie is this much gay:", s)
}

func soo(name string, age int) (string, int) {
	age *= 189
	return fmt.Sprint(name, " is actualy this old in my universe:"), age
}
