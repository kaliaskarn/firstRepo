package main

import "fmt"

func inDelta(n *int) {
	*n = 66
}

func inSlice(ii []int) {
	ii[4] = 99
}

func inMap(xi map[string]int, s string) {
	xi[s] = 89
}

func main() {
	a := 42
	fmt.Println(a)
	inDelta(&a)
	fmt.Println(a)

	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(b)
	inSlice(b)
	fmt.Println(b)

	c := make(map[string]int)
	c["bro"] = 33
	fmt.Println(c["bro"])
	inMap(c, "bro")
	fmt.Println(c["bro"])

}
