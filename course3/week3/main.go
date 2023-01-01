package main

import (
	"fmt"
	"sync"
)

func foo(wg *sync.WaitGroup) {
	fmt.Println("New routine")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go foo(&wg)
	defer fmt.Println("asitha")
	defer fmt.Println("Divisekara")
	wg.Wait()
	fmt.Println("Main routine")
}
