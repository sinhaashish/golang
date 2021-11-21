package main

import "fmt"

func main() {
	p1 := person{"ashish"}
	p2 := person{"sinha"}
	saySomething(p1)
	saySomething(&p2)
}

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println(" the person speaking is ", p.first)

}

type Human interface {
	speak()
}

func saySomething(h Human) {
	h.speak()
}
