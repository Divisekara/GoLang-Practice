package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Print(" ", i)
	}
	fmt.Println()

	myArr := []string{"asitha", "diviseakra", "indrajith", "mudiyanselage"}
	fmt.Println("length of the array is ", len(myArr))
	for i, v := range myArr {
		fmt.Println(i, v)
	}
}
