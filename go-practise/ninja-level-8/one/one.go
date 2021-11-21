package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println(" In mian ")
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func(x int) {
			fmt.Println(" In go routine", x)
			wg.Done()
		}(10)
	}
	wg.Wait()

}
