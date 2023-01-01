package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name  string `json:"name"`
	Addr  string `json:"addr"`
	Phone string `json:"phone"`
}

func main() {
	p1 := person{
		Name:  "asitha",
		Addr:  "mawanella",
		Phone: "+94123456789",
	}
	fmt.Println(p1)
	bArr, _ := json.Marshal(p1)
	fmt.Println(bArr)
	fmt.Println(string(bArr))

	var p2 person
	_ = json.Unmarshal(bArr, &p2)
	// json object and the go struct needs to be fit

}

/*

Json Marshalling
Marshal() - returns json representation as []byte - byte array
Unmarshal() - take a byte array which contains a json object in it and convert it into a golang object


file access is linear access not random access since there is a mechanical delay.
Basic operations in file access
open - get handle for access
read - read bytes into []bytes
Write - write []byte into file
close - release handle
*/
