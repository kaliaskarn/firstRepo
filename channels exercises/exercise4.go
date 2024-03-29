package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	recieve(c, q)

	fmt.Println("about to exit")
}

func recieve(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			{
				fmt.Println("printed number:", v)
			}
		case <-q:
			return
		}
	}
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()
	return c
}
