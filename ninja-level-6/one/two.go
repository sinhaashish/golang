package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	f := foo(x...)
	fmt.Println(f)
	n := bar(x)
	fmt.Println(n)

}

func foo(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func bar(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}
