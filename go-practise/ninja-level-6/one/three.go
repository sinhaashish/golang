package main

import "fmt"

func main() {

	defer func(a int) {
		fmt.Println(" In defer function ", a)
	}(42)

	fmt.Println(" hello world")

}
