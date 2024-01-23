package main

import (
	"encoding/json"
	"fmt"
)

type agent struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := agent{
		First: "Jamie",
		Last:  "Lanister",
		Age:   22,
	}

	p2 := agent{
		First: "Cersei",
		Last:  "Lanister",
		Age:   21,
	}

	both := []agent{p1, p2}
	fmt.Println(both)

	bs, err := json.Marshal(both)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
