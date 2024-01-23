package main

import "fmt"

func main() {
	mp := make(map[string][]string)
	mp["bond_james"] = []string{"shaken, not stirred", "martinis", "fast cars"}
	mp["moneypenny_jenny"] = []string{"`intelligence", "literature", "computer science"}
	mp["no_dr"] = []string{"cats", "ice cream", "sunsets"}
	fmt.Println(mp)
	fmt.Println("============================")
	delete(mp, "bond_james")
	for k, v := range mp {
		fmt.Println(k, v)
	}

}
