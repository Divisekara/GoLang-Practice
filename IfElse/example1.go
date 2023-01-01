package main

import "fmt"

const ten = 10

func main() {
	if ten == 10 && "a" == "b" {
		fmt.Println("This is always false")
	} else if 11 == 11 {
		fmt.Println("This is always true")
	} else {
		fmt.Println("This is something else man")
	}

}

// in golang there's no ternary if else
