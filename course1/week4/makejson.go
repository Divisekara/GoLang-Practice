package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

const (
	name    = "name"
	address = "address"
)

/*
func main() {
	var s string
	arr := make(map[string]string)

	// scanning the name
	fmt.Print("Enter the name = ")
	_, err := fmt.Scan(&s)
	if err != nil {
		panic(err)
	}
	arr[name] = s

	// scanning the address
	fmt.Print("Enter the address = ")
	_, err = fmt.Scan(&s)
	if err != nil {
		panic(err)
	}
	arr[address] = s

	// marshaling the map into byte array
	bArr, err := json.Marshal(arr)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bArr))

}
*/

func main() {
	var m map[string]string
	m = make(map[string]string)
	fmt.Println("Please input your name:")
	br := bufio.NewReader(os.Stdin)
	bname, _, _ := br.ReadLine()
	name := string(bname)
	m["name"] = name
	fmt.Println("Then input your adress:")
	br1 := bufio.NewReader(os.Stdin)
	badress, _, _ := br1.ReadLine()
	adress := string(badress)
	m["adress"] = adress
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Encoding faild")
	} else {
		fmt.Println("Encoded data : ")
		fmt.Println(b)
		fmt.Println("Decoded data : ")
		fmt.Println(string(b))
	}
}
