package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

func readInput() string {
	stdin, err := io.ReadAll(os.Stdin)

	if err != nil {
		panic(err)
	}
	str := string(stdin)
	return str
}

func readFile() string {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	str := string(file)
	return str
}

func readFileLineByLine() []string {
	var lines []string
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func parseDigit(input string) (int, error) {
	digits := map[string]int{
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
	i, ok := digits[input]
	if ok == false {
		return 0, errors.New("err")
	}
	return i, nil
}

// func addDStrings(input1 string, input2 string) string {
// 	num1, err := parseDigit(input1)
// 	if err != nil {
// 		num1 = input1
// 		log.Fatalln(err)
// 	}
// 	num2, err := parseDigit(input2)
// 	if err != nil {
// 		num2 = input2
// 		log.Fatalln(err)
// 	}
// 	num_str := num1 + num2
// 	println(num_str)
// 	return num_str
// }

func main() {
	// lines := strings.Split(readFile(), "\n")
	lines := readFileLineByLine()
	re := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|\d)`)
	var answer int
	for _, l := range lines {
		matches := re.FindAllString(l, -1)
		matches_len := len(matches)
		if matches_len == 0 {
			break
		}
		first := matches[0]
		last := matches[matches_len-1]
		num_f, err := parseDigit(first)
		if err != nil {
			log.Fatalln(err)
		}
		num_l, err := parseDigit(last)
		// num_str := addDStrings(first, last)
		// num, err := strconv.Atoi(num_str)
		if err != nil {
			log.Fatalln(err)
		}
		num := (num_f * 10) + num_l
		answer = answer + num
		fmt.Println(l, first, last, num)
	}
	fmt.Println(answer)
}
