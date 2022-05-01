package main

import (
	"fmt"
	"time"
)

// This is an example of race condition
// 2 goroutines tries to read&write sharedInt and there is no access control.

var sharedInt int = 0
var unusedValue int = 0

func runSimpleReader() {
	for {
		var val int = sharedInt
		if val%10 == 0 {
			unusedValue = unusedValue + 1
		}
		fmt.Println(unusedValue)
	}

}

func runSimpleWriter() {
	for {
		sharedInt = sharedInt + 1
		fmt.Println(sharedInt)
	}
}

func main() {
	go runSimpleReader()
	go runSimpleWriter()
	time.Sleep(10 * time.Second)
}

/*
main() has two anonymous functions runs in concurrently.
Sometimes the increment happens first and then the printing happens.
We cant predict what happens first.
The two routines shares the same variable which cause the race condition.
*/
