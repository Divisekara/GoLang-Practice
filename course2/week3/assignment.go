package main

import (
	"fmt"
)

type Animal struct {
	food  string
	move  string
	sound string
}

func (a Animal) Eat() string {
	fmt.Println(a.food)
	return a.food
}

func (a Animal) Move() string {
	fmt.Println(a.move)
	return a.move
}

func (a Animal) Speak() string {
	fmt.Println(a.sound)
	return a.sound
}

func main() {
	cow := Animal{
		food:  "grass",
		move:  "walk",
		sound: "moo",
	}

	bird := Animal{
		food:  "worms",
		move:  "fly",
		sound: "peep",
	}

	snake := Animal{
		food:  "mice",
		move:  "slither",
		sound: "hsss",
	}

	for true {
		fmt.Println("Enter animal name and information needed > ")
		var animal, info string
		fmt.Scan(&animal, &info)
		switch animal {
		case "cow":
			switch info {
			case "food":
				fmt.Println(cow.food)
			case "move":
				fmt.Println(cow.move)
			case "sound":
				fmt.Println(cow.sound)
			}
		case "bird":
			switch info {
			case "food":
				fmt.Println(bird.food)
			case "move":
				fmt.Println(bird.move)
			case "sound":
				fmt.Println(bird.sound)
			}
		case "snake":
			switch info {
			case "food":
				fmt.Println(snake.food)
			case "move":
				fmt.Println(snake.move)
			case "sound":
				fmt.Println(snake.sound)
			}
		}
	}
}
