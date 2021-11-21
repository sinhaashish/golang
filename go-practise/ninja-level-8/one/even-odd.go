package main

import (
	"fmt"
)

func main() {
	odd := make(chan int)
	even := make(chan int)
	done := make(chan bool)
	fmt.Println("hghh")

	go send(odd, even, done)
	receive(odd, even, done)
	fmt.Println("about to exit")
	//time.Sleep(20)
}

func send(odd, even chan<- int, done chan<- bool) {
	fmt.Println("hghh")
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	//	close(even)
	//	close(odd)
	done <- true
}

func receive(odd, even <-chan int, done <-chan bool) {
	fmt.Println("hghh")
	for {
		select {
		case v := <-even:
			fmt.Println("the value received from the even channel:", v)
		case v := <-odd:
			fmt.Println("the value received from the odd channel:", v)
		case <-done:
			return
		}
	}

}
