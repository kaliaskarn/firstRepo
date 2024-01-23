package main

import "fmt"

func main() {
	c := make(chan int, 1)
	for i := 0; i < 10; i++ {
		go func() {
			for k := 0; k < 10; k++ {
				c <- k
			}
		}()
	}

	for b := 0; b < 100; b++ {
		fmt.Println(b, <-c)
	}
}
