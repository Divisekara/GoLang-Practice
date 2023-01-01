package sample

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 8, 10}
	fmt.Println(a)
	b := a[:]
	fmt.Printf("b: %v\n", b)
	c := a[:5]
	fmt.Printf("c: %v\n", c)
	d := a[3:]
	fmt.Printf("d: %v\n", d)
	e := a[3:6]
	fmt.Printf("e: %v\n", e)

	f := make([]int, 20) // creating an array with predefined sice
	fmt.Printf("f: %v\n", f)
	f[1] = 1234
	// f[100] = 123455 then we cant assign a value to an index beyond the limit but we can append values
	f = append(f, 1, 2, 3, 4, 5)
	fmt.Printf("f after append: %v\n", f)

	f = append(f, []int{99, 98, 97}...)
	fmt.Printf("f after concatenating another list: %v\n", f)
	// push method can be achieved using append function but for pop method is little bit tedious

	f = f[:len(f)-1]
	fmt.Printf("f after popping: %v\n", f)
}
