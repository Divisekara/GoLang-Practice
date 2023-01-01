package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	var s string
	var p Person
	var sli []Person
	var names []string

	fmt.Print("Enter the file name = ")
	_, err := fmt.Scan(&s)
	f, err := os.Open(s)
	if err != nil {
		panic(errors.New("file opening error"))
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		names = strings.Split(line, " ")
		p.FirstName = names[0]
		p.LastName = names[1]
		sli = append(sli, p)
	}

	for _, v := range sli {
		fmt.Println(v.FirstName, v.LastName)
	}

}
