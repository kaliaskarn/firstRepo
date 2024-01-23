package main

import (
	"fmt"
	"runtime"
	"sync"
)

var gs int = 100
var counter int = 0

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			runtime.Gosched()
			counter++
			fmt.Println(counter)
			wg.Done()
			mu.Unlock()
		}()

	}
	wg.Wait()

}
