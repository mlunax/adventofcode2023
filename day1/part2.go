package main

import (
	"errors"
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

func parseDigit(input string) (string, error) {
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
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
	}
	i, ok := digits[input]
	if ok == false {
		return "", errors.New("No spelled digit")
	}
	return i, nil
}

func addDStrings(input1 string, input2 string) string {
	num1, err := parseDigit(input1)
	if err != nil {
		num1 = input1
	}
	num2, err := parseDigit(input2)
	if err != nil {
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
		num_str := addDStrings(first, last)
		num, err := strconv.Atoi(num_str)
		if err != nil {
			log.Fatalln(err)
		}
		answer = answer + num
		fmt.Println(l, first, last, num)
	}
	fmt.Println(answer)
}
