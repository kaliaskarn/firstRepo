package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(even, odd, quit)

	recieve(even, odd, quit)

}

func send(even, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}

func recieve(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			{
				fmt.Println("value from even channel", v)
			}
		case v := <-odd:
			{
				fmt.Println("value from odd channel", v)
			}
		case i, ok := <-quit:
			if !ok {
				fmt.Println("from comma ok", i)
				return
			} else {
				fmt.Println("from comma ok", i)
			}
		}
	}
}
