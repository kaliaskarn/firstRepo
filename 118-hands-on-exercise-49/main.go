package main

import "fmt"

func main() {
	mp := make(map[string][]string)
	mp["bond_james"] = []string{"shaken, not stirred", "martinis", "fast cars"}
	mp["moneypenny_jenny"] = []string{"`intelligence", "literature", "computer science"}
	mp["no_dr"] = []string{"cats", "ice cream", "sunsets"}
	fmt.Println(mp)
	fmt.Println("============================")

	for k, v := range mp {
		fmt.Println(k)
		for i, a := range v {
			fmt.Println(i, a)
		}
	}

	mp["fleming_ian"] = []string{"steaks", "cigars", "espionage"}
	for k, v := range mp {
		fmt.Println(k, v)
	}

}
