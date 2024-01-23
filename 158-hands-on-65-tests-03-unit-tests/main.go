package main

import (
	"fmt"
)

func main() {
	fmt.Println(Rapist("Michael Jackson"))
}

func Rapist(victim string) string {
	return fmt.Sprint(victim)
}
