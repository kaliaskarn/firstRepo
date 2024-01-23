package main

import (
	"fmt"
)

type customErr struct {
	first string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("error is made here: %v", ce.first)
}

func main() {
	c1 := customErr{
		first: "AKnur",
	}
	foo(c1)
}

func foo(e error) {
	fmt.Println(e)
}
