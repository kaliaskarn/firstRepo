package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{1, 23, 4341, 243, 443, 23, 44, 75}
	xs := []string{"Broski", "Nigger", "James", "Cmon", "Amanda", "Dr.", "Sexy bois"}
	sort.Ints(xi)
	sort.Strings(xs)
	fmt.Println(xs)
	fmt.Println(xi)
}
