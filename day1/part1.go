package main

import (
	"io"
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

func main() {
	input := readInput()
	lines := strings.Split(input, "\n")
	re := regexp.MustCompile(`\d`)
	var answer int
	for _, l := range lines {
		matches := re.FindAllString(l, -1)
		matches_len := len(matches)
		if matches_len == 0 {
			continue
		}
		first := matches[0]
		last := matches[matches_len-1]
		num_str := first + last
		num, err := strconv.Atoi(num_str)
		if err != nil {
			panic(err)
		}
		answer = answer + num
	}
	println(answer)
}
