package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	favIceCream []string
}


var m map[string]Vertex

func main() {

	p := person{
		firstName:   "Ashish",
		lastName:    "Sinha",
		favIceCream: []string{"ice", "cream"},
	}
	p1 := person{
		firstName:   "Ashish1",
		lastName:    "Sinha1",
		favIceCream: []string{"ice1", "cream1"},
	}

	m = make([string]person)

	m[p.lastName]= p
	m[p1.lastName]= p1

	for range k, v := 

}
