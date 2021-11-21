package main

import (
	"fmt"
	"log"
)

// func main() {
// 	c := factorial(4)
// 	for n := range c {
// 		fmt.Println(n)
// 	}
// }

func main() {
	defer func() {
		fmt.Println("world")
	}()

	fmt.Println("hello")
	log.Fatal("oops")
}

// func factorial(n int) chan int {
// 	ch := make(chan int)
// 	go func() {
// 		result := 1
// 		for i := n; i > 0; i-- {
// 			result *= i
// 		}
// 		ch <- result
// 		close(ch)
// 	}()
// 	return ch
// }
