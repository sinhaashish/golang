package main

import (
	"fmt"
	"strconv"
	"strings"
)

func complexNumberMultiply(num1 string, num2 string) string {
	X := strings.Split(num1, "+")
	Y := strings.Split(num2, "+")
	a, _ := strconv.Atoi(X[0])
	b, _ := strconv.Atoi(X[1][:len(X[1])-1])
	c, _ := strconv.Atoi(Y[0])
	d, _ := strconv.Atoi(Y[1][:len(Y[1])-1])
	fmt.Println()
	return strconv.Itoa(a*c-b*d) + "+" + strconv.Itoa(b*c+a*d) + "i"

}
