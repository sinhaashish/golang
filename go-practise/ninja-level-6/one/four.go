package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	a := person{
		first: "ashis",
		last:  "sinha",
		age:   36,
	}
	b := person{
		first: "yahoo",
		last:  "google",
		age:   36,
	}
	a.speak()
	b.speak()

}

func (p *person) speak() {
	fmt.Printf(" Speaker %d , Age %d", p.first, p.age)
}

type Healp interface {
	speak()
}
