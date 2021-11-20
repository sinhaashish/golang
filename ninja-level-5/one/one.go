package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	favIceCream []string
}

var m map[string]person

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

	m = make(map[string]person)

	m[p.lastName] = p
	m[p1.lastName] = p1

	for _, v := range m {
		fmt.Println(v.firstName)
		fmt.Println(v.firstName)
		for i, val := range v.favIceCream {
			fmt.Println(i, val)
		}
		fmt.Println("-------")
	}

}
