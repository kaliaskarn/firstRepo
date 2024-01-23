package main

import (
	"fmt"
	"runtime"
	"sync"
)

func murder(target string) {
	fmt.Println("the targer was killed, name:", target)
}

var victims = []string{"Baby", "Selena Gomez", "Bieber"}

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("CPUs:", runtime.NumGoroutine())

	for _, v := range victims {
		murder(v)
	}
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		runtime.Gosched()
		fmt.Println("this is literally just using a new thread")
		wg.Done()
	}()

	go func() {
		runtime.Gosched()
		fmt.Println("this as well")
		wg.Done()
	}()
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("CPUs:", runtime.NumGoroutine())
	wg.Wait()

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("CPUs:", runtime.NumGoroutine())

}
