package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	counter := 0

	const num = 100
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			mutex.Lock()
			val := counter
			val++
			counter = val
			fmt.Println("Goroutines:", runtime.NumGoroutine())
			mutex.Unlock()
			wg.Done()
		}()

	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)

}
