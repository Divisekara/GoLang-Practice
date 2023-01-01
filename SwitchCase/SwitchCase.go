package main

import "fmt"

func main() {
	number := 3
	switch number {
	case 1:
		fmt.Println("number is 1")
	case 2:
		fmt.Println("number is 2")
	case 3:
		fmt.Println("number is 3")
	default:
		fmt.Println("number has nothing to do with this. He is minding is own")
	}
}
