package main

import "fmt"

func DrawShape(s interface{}) {
	switch sh := s.(type) {
	case string:
		fmt.Println(sh)
	default:
		fmt.Println("divisekara")
	}
}

func main() {
	DrawShape(10)
}
