package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	c <- 10
	fmt.Println(c)
	x := 0
	go func() {
		x++
	}()

	//time.Sleep(1 * time.Second)
	y := <-c
	fmt.Println(y)
	go func() {
		fmt.Println(x)
	}()

	time.Sleep(1 * time.Second)

}

/*
main() has two anonymous functions runs in concurrently.
Sometimes the increment happens first and then the printing happens.
We cant predict what happens first.
The two routines shares the same variable which cause the race condition.
*/
