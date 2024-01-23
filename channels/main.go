package main

import "fmt"

func main() {
	c := make(chan int, 1)
	c1 := make(chan int, 1)

	go func() { c1 <- 66 }()

	c <- 42

	fmt.Println(<-c)
	fmt.Println(<-c1)
}
