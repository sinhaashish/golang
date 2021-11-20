package main

import "fmt"

func main() {

	sq := Square{2}
	c := Cirdle{1}
	info(&sq)
	info(&c)

}

type Square struct {
	side int
}

type Cirdle struct {
	radius int
}

func (a *Square) Area() int {
	return a.side * a.side
}

func (a *Cirdle) Area() int {
	return a.radius * 22 * a.radius
}

type Shape interface {
	Area() int
}

func info(s Shape) {
	fmt.Println(s.Area())
}
