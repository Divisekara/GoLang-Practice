package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	sound      string
}

type AnimalInterface interface {
	Eat()
	Move()
	Speak()
}

func (aniname Animal) Eat() {
	fmt.Println(aniname.food + "\n")
	return
}

func (aniname Animal) Move() {
	fmt.Println(aniname.locomotion + "\n")
	return
}

func (aniname Animal) Speak() {
	fmt.Println(aniname.sound + "\n")
	return
}

func main() {
	m := make(map[string]Animal)
	m["Cow"] = Animal{"grass", "walk", "moo"}
	m["Bird"] = Animal{"worms", "fly", "peep"}
	m["Snake"] = Animal{"mice", "slither", "hsss"}

	var generalAni AnimalInterface

	for {
		var command, requestAniname, requestAniType string
		fmt.Print("> ")
		fmt.Scan(&command, &requestAniname, &requestAniType)

		if command == "query" {
			generalAni = m[requestAniname]
			switch requestAniType {
			case "eat":
				generalAni.Eat()
			case "move":
				generalAni.Move()
			case "speak":
				generalAni.Speak()
			}

		} else if command == "newanimal" {
			m[requestAniname] = m[requestAniType]
			fmt.Println("Created it!\n")

		} else {
			fmt.Println("Incorrect Input!!END")
			break
		}
	}
}
