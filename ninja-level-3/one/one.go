package main

import "fmt"

func main() {
	x := 1
	for {

		if x < 100 {
			fmt.Println(x)
			x++
		}
	}
}
