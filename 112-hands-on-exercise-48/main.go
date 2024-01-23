package main

import (
	"fmt"
)

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "I'm 008."}
	xr := [][]string{x, y}
	fmt.Println(xr)
	fmt.Println("=========================")

	for a, b := range xr {
		fmt.Println(a, b)
		for c, d := range b {
			fmt.Println(c, d)
		}
	}

}
