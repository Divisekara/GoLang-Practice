package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var ianLetters = []string{"i", "a", "n"}

func findLetter(stringToSearch string, letterToFind string) int {
	lowerIndex := strings.Index(stringToSearch, letterToFind)
	if lowerIndex == -1 {
		return strings.Index(stringToSearch, strings.ToUpper(letterToFind))
	}

	return lowerIndex
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter any string: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("ERROR: Failed to read input.")
		os.Exit(0)
	}

	prevFoundIndex := -1
	for _, letter := range ianLetters {
		index := findLetter(input, letter)
		if index == -1 || prevFoundIndex > index {
			fmt.Println("Not found!")
			os.Exit(0)
		}

		prevFoundIndex = index
	}

	fmt.Println("Found!")
}
