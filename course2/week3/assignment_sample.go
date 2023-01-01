package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	sound      string
}

func (animal *Animal) init(food, locomotion, sound string) {
	animal.food = food
	animal.locomotion = locomotion
	animal.sound = sound

}

func (animal *Animal) Speak() {
	fmt.Println("speak like ", animal.sound)
}

func (animal *Animal) Eat() {
	fmt.Println("eats ", animal.food)
}

func (animal *Animal) Move() {
	fmt.Println("moves like ", animal.locomotion)
}

func main() {
	var cow, bird, snake, animal Animal
	cow.init("grass", "walk", "moo")
	bird.init("worms", "fly", "peep")
	snake.init("mice", "slither", "hsss")
	for {
		fmt.Println("enter animal and action separated by spaces (enter x to exit): ")
		fmt.Print("> ")
		inputReader := bufio.NewReader(os.Stdin)
		user_input, _ := inputReader.ReadString('\n')
		user_input = strings.TrimSuffix(user_input, "\r\n")
		user_input = strings.ToLower(user_input)
		if user_input == "x" {
			break
		}
		parts := strings.Split(user_input, " ")
		// fmt.Println("txt ", txt)
		// fmt.Println(parts)
		if len(parts) < 2 {
			fmt.Println("incorrect input")
			continue
		}
		animalStr := parts[0]
		action := parts[1]
		switch animalStr {
		case "cow":
			animal = cow
		case "bird":
			animal = bird
		case "snake":
			animal = snake
		default:
			fmt.Println("invalid animal. should be cow, snake, or bird")
			continue
		}
		switch action {
		case "move":
			animal.Move()
		case "eat":
			animal.Eat()
		case "speak":
			animal.Speak()
		}
	}
}
