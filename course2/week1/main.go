package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	for j := 0; j <= 0; j++ {
		i := 1
		for {
			fmt.Printf("Commit %d \n", i)
			changeFIle(i)
			//time.Sleep(time.Second * 1)
			// assigning the date and the month patter
			m, d := (i*3)/28+1, (i*3)%28
			//git(m, d) // here you can change the gap

			changeFIle(i * 2)
			m, d = (i*4)/28+1, (i*3)%28
			git(m, d) // here you can change the gap
			if i == 15 {
				break
			}
			i++
		}
	}
}

func changeFIle(i int) {
	f, err := os.Create("sample.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(strconv.Itoa(i))
	if err != nil {
		panic(err)
	}
}

func git(m, d int) {
	cmd1 := exec.Command("bash", "git.sh", strconv.Itoa(m), strconv.Itoa(d))
	_, err := cmd1.Output()

	//cmd1 := exec.Command("git", "add", ".")
	//_, err := cmd1.Output()
	//
	//cmd2 := exec.Command("git", "commit", "-m", "week 3 assignment done", `--date="2019-02-26 1:1:1"`)
	//_, err = cmd2.Output()
	//
	//cmd3 := exec.Command("git", "push",
	//	"https://Divisekara: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@github.com/Divisekara/GOlang-practice.git",
	//	"--all")
	//_, err = cmd3.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

}

/*

Von Neumann Bottleneck -
Concurrency might be speed up the program even when its run on single core processor - when external calls happens and core is idle, then another task can be run

*/
