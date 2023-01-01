package main

import (
	"fmt"
	"net/http"
	"reflect"
)

func main() {
	fmt.Println("asitha")
	var x [5]int
	fmt.Println(x)
	fmt.Println(x[2])

	var y [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(y)

	//size inferring
	z := [...]int{1, 2, 3, 4, 99999}
	fmt.Println(z)

	for i, v := range z {
		fmt.Println(i, v)
	}

	a := []string{"a", "b", "c", "d"}
	a = append(a, "a", "b", "c", "d", "e", "f")
	fmt.Println(reflect.TypeOf(a), len(a), cap(a))

	// 2 argument version of make prepare the slice's length is equal to the capacity
	b := make([]int, 10)
	fmt.Println(b)

	// 3 argument version of make specify the capacity and the length where length <= capacity
	c := make([]int, 4, 4)
	fmt.Println(c)

	c = append(c, 1)
	fmt.Println(cap(c))

	var ab map[float32]float64
	fmt.Println("ab = ", ab, reflect.TypeOf(ab))
	ab = make(map[float32]float64, 4)
	fmt.Println(ab)

	d := map[string]string{}
	e := map[int]string{}
	f := map[interface{}]interface{}{}
	// so fo the maps any type of keys and any type of values can be there
	fmt.Println(d, e, f)

	// map literal definition
	g := map[string]string{"a": "asitha", "b": "some text"}
	fmt.Println(g)

	h := map[rune]rune{'a': 'b'}
	fmt.Println(h)

	// modifying a key value or adding a key value
	g["b"] = "johnson"
	g["c"] = "divisekara"
	delete(g, "f") // there's no key name 'f' but it ain't giving any errors
	fmt.Println(g["a"], g["b"], g["c"])
	delete(g, "b")
	fmt.Println(g)

	// two value assignment tests for the existence of the key
	id, ok := g["a"]
	if !ok {
		panic("this aint right")
	}
	fmt.Println(id, ok)
	// deleting a key value from map

	for k, v := range g {
		fmt.Println(k, v)
	}

	type person struct {
		name    string
		address string
		grades  map[string]int
		friends interface{}
	}

	var p1 person
	p1.name = "asitha"
	p1.address = "divisekara"
	p1.grades = map[string]int{"maths": 100, "science": 78, "history": 98}
	// this is awesome since it allows to add person struct inside the person struct as an interface
	p1.friends = person{
		name:    "saman",
		address: "bill",
		grades:  nil,
		friends: nil,
	}
	fmt.Println(p1)

	xa := [...]int{1, 2, 3, 4, 5}
	ya := xa[0:2]
	za := xa[1:4]
	// this is confusing since when tailed from start the capacity goes down. and tailed from end doesn't change capacity
	fmt.Println(len(ya), cap(ya), len(za), cap(za))

	//con, _ := net.Dial("tcp", "uci.edu:80")
	com, _ := http.Get("www.uci.edu")
	fmt.Println("this is connection = ", com)

}

/*

composite data types are data types that aggregate other data types together. then they bring lots of different data types together.
1. arrays - fixed length series of elements of a chosen one basic data type

2. slices - the length of the array can change. Slice is a window on an underlying array, variable size upto the whole array
the pointer indicate the start of the slice
Length indicates the number of elements in the slice len()
Capacity is the maximum number of elements from start of slice to end of array cap()
size of the slice can be increased by the append function

3, Hash tables - key value pairs
Each value is associated with a unique key
Hash function is used to compute the slot for a key'
advantages in hash tables - faster lookups than lists. the time taken for the lookup is constant and for the lists its linear time, can have arbitrary keys
disadvantages - two keys can be hash to same slot where a collision happens, but its very rare which is negligible


3. maps - map is the golang implementation of the hash table

4. structs - is another aggregate data type


*/
