package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

var digitsMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func readInput() string {
	stdin, err := io.ReadAll(os.Stdin)

	if err != nil {
		panic(err)
	}
	str := string(stdin)
	return str
}

func parseDigit(input string) (int, error) {
	i, ok := digitsMap[input]
	if ok == false {
		return 0, errors.New("err")
	}
	return i, nil
}

func getDigit(line string) []int {
	digits := []int{}
	for i, c := range line {
		if dig, err := parseDigit(string(c)); err == nil {
			digits = append(digits, dig)
		} else {
			for spelling, number := range digitsMap {
				if strings.HasPrefix(line[i:], spelling) {
					digits = append(digits, number)
				}
			}
		}
	}
	return digits
}

func main() {
	lines := strings.Split(readInput(), "\n")
	var answer int
	for _, l := range lines {
		matches := getDigit(l)
		if len(matches) == 0 {
			continue
		}
		first := matches[0]
		last := matches[len(matches)-1]
		num := (first * 10) + last
		answer = answer + num
		// fmt.Println(l, first, last, num)
	}
	fmt.Println(answer)
}
