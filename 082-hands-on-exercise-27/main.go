package main

import (
	"fmt"
)

func main() {
	i := 20
	for {
		if i < 0 {
			break
		}
		fmt.Println(i)
		i--
	}
}
