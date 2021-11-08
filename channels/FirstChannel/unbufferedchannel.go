package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 1; i < 10; i++ {
			ch <- i
		}
	}()

	go func() {
		for {
			fmt.Println(<-ch)
		}
	}()

	time.Sleep(time.Second)
}
