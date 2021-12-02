// Given a valid (IPv4) IP address, return a defanged version of that IP address.

// A defanged IP address replaces every period "." with "[.]".

package main

import (
	"fmt"
	"strings"
)

func main() {
	address := "1.1.1.1"
	fmt.Println(defangIPaddr(address))

}

func defangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}
