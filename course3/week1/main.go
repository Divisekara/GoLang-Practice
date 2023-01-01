package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	x := 1
	fmt.Println(x)
	err := increment(&x)
	if err != nil {
		panic(err)
	}
	fmt.Println(x)
}

func increment(ptr interface{}) (err error) {
	// type assertion
	v, ok := ptr.(*int)
	if ok != true {
		return errors.New("interface type error")
	}
	fmt.Println(ok, v, reflect.TypeOf(v))
	// incrementing the value of the pointer
	*v++
	return nil
}

/*
Parameters - inputs define at the definition of the function
Arguments - inputs given at the function call

passing arrays vs passing slices - once you update slice inside the function then the real slice changes.

*/
