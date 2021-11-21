package main

import "fmt"

type TEST int

var x TEST
var y int

func main() {
	fmt.Printf(" \n Type  %T", x)
	fmt.Printf(" \n value before  %v", x)
	x = 42
	fmt.Printf(" \n value after  %v \n ", x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n ", y)
}
