package main

import "fmt"

func main() {
	var n int
	var sli []int
	for i := 0; i < 10; i++ {
		fmt.Print("Enter number - ")
		_, err := fmt.Scanln(&n)
		if err != nil {
			panic(err)
		}
		sli = append(sli, n)
	}
	bubbleSort(&sli)
	fmt.Println(sli)
}

// bubbleSort algorithm
func bubbleSort(sli *[]int) {
	for j := 0; j < len(*sli); j++ {
		for i := 0; i < len(*sli)-1; i++ {
			if (*sli)[i] > (*sli)[i+1] {
				swap(sli, i, i+1)
			}
		}
	}
}

// swap function
func swap(sli *[]int, x, y int) {
	(*sli)[x], (*sli)[y] = (*sli)[y], (*sli)[x]
}
