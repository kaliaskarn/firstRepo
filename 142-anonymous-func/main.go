package main

import (
	"fmt"
)

func main() {
	bitch()

	func() {
		fmt.Println("im not understanding interfaces")
	}()

	func(s string) {
		fmt.Println("why are all the bitches remind me of ", s)
	}("aknur")
}

func bitch() {
	fmt.Println("printed using anonymous ")
}
