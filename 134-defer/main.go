package main

import (
	"fmt"
)

func main() {
	defer boo()
	foo()
}

func boo() {
	fmt.Println("this will be defered and executed after the function surrounding it")
}

func foo() {
	fmt.Println("printing foo")
}
