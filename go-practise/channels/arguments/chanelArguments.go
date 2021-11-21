package main

import "fmt"

func main() {
	c := putInChannel()
	sum := sumFromChannel(c)
	for i := range sum {
		fmt.Println(i)
	}

}

func putInChannel() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i <= 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func sumFromChannel(ch chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range ch {
			sum = sum + n
		}
		out <- sum
		close(out)
	}()
	return out
}
