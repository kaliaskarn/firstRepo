package main

import (
	"fmt"
)

func main() {
	type engine struct {
		electric bool
	}

	type vehicle struct {
		engine
		make  string
		model string
		doors int
		color string
	}

	v1 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "idk what to write here",
		model: "electric",
		doors: 4,
		color: "black",
	}

	v2 := vehicle{
		engine: engine{
			electric: true},
		make:  "idk what to write here",
		model: "diesel/hybrid",
		doors: 2,
		color: "silver",
	}

	fmt.Println(v1)
	fmt.Println(v2)

}
