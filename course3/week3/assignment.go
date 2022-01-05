package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func sortPartition(arr []int, ch chan []int, wg sync.WaitGroup) {
	defer wg.Done()
	l := len(arr)

	for i := 0; i < l; i++ {
		for j := 0; j < l-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	ch <- arr
}

func merge(arr1, arr2 []int) []int {
	var out []int
	l1 := len(arr1)
	l2 := len(arr2)
	var p1, p2 int
	for i := 0; i < l2+l1; i++ {
		if p1 == l1 {
			out = append(out, arr2[p2:]...)
			break
		} else if p2 == l2 {
			out = append(out, arr1[p1:]...)
			break
		} else if arr1[p1] <= arr2[p2] {
			out = append(out, arr1[p1])
			p1++
		} else {
			out = append(out, arr2[p2])
			p2++
		}
	}

	return out

}

func main() {
	var wg sync.WaitGroup
	//arr := []int{6, 3, 9, 1, 0, 3, 6, 4, 23, 45, 67, 99}

	fmt.Println("Enter integers seperated by spaces - ")
	br := bufio.NewReader(os.Stdin)
	a, _, _ := br.ReadLine()
	ns := strings.Split(string(a), " ")
	fmt.Println(ns)

	var arr []int
	for i := 0; i < len(ns); i++ {
		n, _ := strconv.Atoi(ns[i])
		arr = append(arr, n)
	}

	l := len(arr) / 4
	arr1 := arr[0:l]
	arr2 := arr[l : l*2]
	arr3 := arr[l*2 : l*3]
	arr4 := arr[l*3:]

	ch := make(chan []int)

	wg.Add(4)
	go sortPartition(arr1, ch, wg)
	go sortPartition(arr2, ch, wg)
	go sortPartition(arr3, ch, wg)
	go sortPartition(arr4, ch, wg)
	s1 := <-ch
	s2 := <-ch
	s3 := <-ch
	s4 := <-ch

	close(ch)

	//fmt.Println(s1, s2, s3, s4)

	s5 := merge(s1, s2)
	//fmt.Println(s5)
	s6 := merge(s3, s4)
	//fmt.Println(s6)
	s7 := merge(s5, s6)
	fmt.Println(s7)

}
