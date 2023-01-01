package main

import "fmt"

var funcVar func(int) int

// treating a function as a first class object
func incFn(x int) int {
	return x + 1
}

func main() {
	funcVar = incFn
	fmt.Println(funcVar(1))
}
