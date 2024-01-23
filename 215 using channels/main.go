package main

import "fmt"

func main() {
	c := make(chan int)

	go foo(c)

	bar(c)

	// recieve(c)

	fmt.Println("exiting this")

}

func foo(c chan<- int) {
	c <- 42
}

func bar(c <-chan int) {
	fmt.Println(<-c)
}

// this also works
//  receive channel
// func receive(c <-chan int) {
// 	fmt.Println("the value received from the channel:", <-c)
// }
