package main

import "fmt"

type MyInt int

func (m MyInt) Double() int {
	fmt.Println(m * 2)
	return int(m * 2)
}

type data struct {
	p1 float64
	s1 string
}

func (d *data) printing() {
	fmt.Println(d.s1)
	fmt.Println(d.p1)
}

func main() {
	v := MyInt(3)
	v.Double()
	v.Double()

	d := data{
		p1: 10,
		s1: "asitha",
	}
	d.printing()
}
