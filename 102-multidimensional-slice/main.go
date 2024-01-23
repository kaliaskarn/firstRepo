package main

import (
	"fmt"
)

func main() {
	jb := []string{"nigga", "georgie", "males", "dudes", "men"}
	jm := []string{"bros", "men", "cakes", "gyatt", "gandalf"}

	fmt.Println(jb)
	fmt.Println(jm)
	fmt.Println("-------------- ")

	xs := [][]string{jb, jm}
	fmt.Println(xs)

}
