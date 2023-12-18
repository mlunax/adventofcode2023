package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readInput() string {
	stdin, err := io.ReadAll(os.Stdin)

	if err != nil {
		panic(err)
	}
	str := string(stdin)
	return str
}

func parseSpelledDigit(input string) string {
	digits := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	i, ok := digits[input]
	if ok == false {
		return "-1"
	}
	return i
}

func addSDStrings(input1 string, input2 string) string {
	num1 := parseSpelledDigit(input1)
	if num1 == "-1" {
		num1 = input1
	}
	num2 := parseSpelledDigit(input2)
	if num2 == "-1" {
		num2 = input2
	}
	num_str := num1 + num2
	return num_str
}

func main() {
	lines := strings.Split(readInput(), "\n")
	re := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|\d)`)
	var answer int
	for _, l := range lines {
		matches := re.FindAllString(l, -1)
		matches_len := len(matches)
		if matches_len == 0 {
			continue
		}
		first := matches[0]
		last := matches[matches_len-1]
		num_str := first + " " + last
		num, err := strconv.Atoi(num_str)
		if err != nil {
			num_str = addSDStrings(first, last)
			num, err = strconv.Atoi(num_str)
			if err != nil {
				log.Fatalln(err)
				continue
			}

		}
		answer = answer + num
		fmt.Println(l, first, last, num)
	}
	fmt.Println(answer)
}
