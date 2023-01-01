package main

import "fmt"

func main() {
	hello("asitha", "divisekara", "indrajith", "mudiyanselage")
}

// hello this is variadic function
func hello(args ...string) int {
	for i, v := range args {
		fmt.Println(i, v)
	}

	return 0
}

// using functions developer can maintain, encapsulation, readability and maintainability
// in computer science there's a statement as "a function should do just one thing, but it should do that thing damn well."
