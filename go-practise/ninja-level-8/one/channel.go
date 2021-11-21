package main

import "fmt"

func main() {

	c := foo()
	receive(c)

}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func foo() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}
