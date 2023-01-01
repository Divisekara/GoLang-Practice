package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var n string
	var number int
	sli := make([]int, 3)
	i := 0
	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			panic(err)
		}
		if n == "X" {
			break
		}
		number, err = strconv.Atoi(n)
		if i < 3 {
			sli[i] = number
		} else {
			sli = append(sli, number)
		}
		if err != nil {
			panic(err)
		}
		sort.Ints(sli)
		fmt.Println(sli)
		i++
	}
	fmt.Println("I quit")
}
