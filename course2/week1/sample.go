package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd1 := exec.Command("git", "add", ".")
	_, err := cmd1.Output()

	cmd2 := exec.Command("git", "commit", "-m", "week 3 assignment done", `--date="2019-02-06 1:1:1"`)
	_, err = cmd2.Output()

	cmd3 := exec.Command("git", "push",
		"https://Divisekara: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@github.com/Divisekara/GOlang-practice.git",
		"--all")
	_, err = cmd3.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	//fmt.Println(string(out))
}
