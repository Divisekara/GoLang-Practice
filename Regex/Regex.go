package main

import (
	"fmt"
	"regexp"
)

func main() {
	pattern := "^\\* \\* ([1-9]|[1-2][0-9]|3[0-1]) ([1-9]|1[0-2]) \\* (19[7-9][0-9]|20[0-9][0-9])$"
	reg, _ := regexp.Compile(pattern)
	frequencies := []string{
		// correct patterns expecting true
		"* * 1 1 * 1970",
		"* * 10 1 * 1970",
		"* * 19 1 * 1970",
		"* * 30 1 * 1970",
		"* * 31 1 * 1970",
		"* * 1 1 * 1970",
		"* * 29 10 * 1970",
		"* * 30 11 * 1970",
		"* * 31 12 * 1970",
		"* * 31 12 * 1970",
		"* * 29 10 * 1979",
		"* * 30 11 * 1999",
		"* * 10 10 * 2000",
		"* * 10 10 * 2000",
		"* * 11 10 * 2099",
		"* * 12 10 * 2098",
		"* * 11 10 * 2022",
		"* * 11 10 * 2022",
		"* * 11 10 * 2022",
	}
	for _, freq := range frequencies {
		match := reg.MatchString(freq)
		fmt.Println(freq, " = ", match)
	}

	fmt.Println("\nwrong frequencies\n")

	xfrequencies := []string{
		// correct patterns expecting true
		"* * 0 1 * 1970",
		"* * 32 1 * 1970",
		"* * 99 1 * 1970",
		"* * 100 1 * 1970",
		"* *  1 * 1970",
		"* * # 1 * 1970",
		"* * 0 10 * 1970",
		"* * 31 13 * 1970",
		"* * 99 12 * 1970",
		"* *  12 * 1970",
		"* * 29 % * 1979",
		"* * 30 11 * 1969",
		"* * 10 10 * 20222",
		"* * 10 10 * 2100",
		"* * 11 10 * 3099",
		"* * 12 10 * 098",
		"* * 11 10 * ",
		"  * * 11 10 * 2022",
		"* * 11 10 * 2022 ",
		"* * 11 10 * 2022$",
		"* * * * * *",
		"* * 7/7 * * *",
		"* * * * 4 *",
		"* * 7/14 * * *",
		"* * * * 5#1 *",
		"* * * * 1#2 *",
		"* * * * 4L *",
		"* * 12 * * *",
		"* * 10 7/2 * *",
		"* * 20 3/3 * *",
		"* * L 3/3 * *",
		"* *  31  3  * *",
	}
	for _, freq := range xfrequencies {
		match := reg.MatchString(freq)
		fmt.Println(freq, " = ", match)
	}
	/*
		+ plus sign = match at least one
		? question mark = optional mark
		* star mark = combination of plus sign and question mark , it is optional as well as can match as much as possible in a raw
		. stop mark =  period character t.. = t and two any characters







	*/

	/*
		str1 := "asitha@gmail.com"
		match1, err := regexp.MatchString("asitha", str1)
		if err != nil {
			panic(err)
		}
		fmt.Println(match1)

		str2 := "asitha divisekara()"
		match2, err := regexp.MatchString("asitha()", str2)
		if err != nil {
			panic(err)
		}
		fmt.Println(match2)

		re1, _ := regexp.Compile("geek")
		str3 := "I love geeksforgeeks"
		// here the output of the
		match3 := re1.FindStringIndex(str3)
		fmt.Println(match3)

		re2, _ := regexp.Compile("ege")
		str4 := "egege"
		match4 := re2.FindStringIndex(str4)
		// this is a special occasion since we get 3 as the index of second occurrence,
		// but it should be 2
		fmt.Println("match4 - ", match4)

		str5 := "I love computer science"
		match5 := re1.FindStringIndex(str5)
		fmt.Println(match5)

		str6 := "20024-vani_gupta"
		re3, _ := regexp.Compile("[0-9]+-v.*g")
		match6 := re3.FindString(str6)
		fmt.Println("match 6 is - ", match6)
	*/
}

/*
A regular expression or regex is a special sequence of characters that defines
a search pattern that is used for matching specific text. There's a built-in
package for regex in golang.
It contains all list of actions like
1. filtering
2. replacing
3. validating
4. extracting
*/

/*
to store complicated regular expressions for reuse later, Compile()
method parses a regular expression and return a regexp object if successful which can be
used to match the text.
*/
