package main

import "fmt"

// Speaker interface
type Speaker interface {
	Speaks()
}

type Dog struct{ name string }

// Speaks implementation for the interface
func (d Dog) Speaks() {
	fmt.Println(d.name)
}

func main() {
	var s1 Speaker
	//var d1 *Dog
	d1 := Dog{name: "asitha"}
	//d2 := Dog{"asitha"}
	// as long as d1 satisfy the interface
	s1 = d1
	s1.Speaks()
	//var s2 Speaker

}

/*
dynamic type is Dog
dynamic value is d1

*/
