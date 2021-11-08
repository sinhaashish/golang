package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		wg.Done()
	}()

	go func() {
		for i := 10; i < 20; i++ {
			ch <- i
		}
		wg.Done()
	}()

	go func() {

		wg.Wait()
		close(ch)
	}()

	for n := range ch {
		fmt.Println(n)
	}

}
