package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	favIceCream []string
}

func main() {
	p := person{
		"ashish", "sinha", []string{"ice", "cream"},
	}
	fmt.Println(p)
	p.modify()
	fmt.Println(p)
}

type Human interface {
	modify()
}

func (p *person) modify() {
	p.firstName = "kumar"

}
