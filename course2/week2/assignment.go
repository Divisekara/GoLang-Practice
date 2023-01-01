package main

import "fmt"

func main() {
	var a float64
	fmt.Print("Enter acceleration = ")
	fmt.Scan(&a)

	var v float64
	fmt.Print("Enter Initial velocity = ")
	fmt.Scan(&v)

	var d float64
	fmt.Print("Enter intial displacement = ")
	fmt.Scan(&d)

	fn := GenDisplaceFn(a, v, d)
	fmt.Println("displacement after 3 seconds = ", fn(3))
	fmt.Println("displacement after 5 seconds = ", fn(5))
}

func GenDisplaceFn(a, v, d float64) func(float64) float64 {
	return func(t float64) float64 {
		return a*t*t/2 + v*t + d
	}
}
