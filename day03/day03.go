package main

import (
	"fmt"
	"io"
	"os"
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

func part1(lines []string) {
	fmt.Println("Part 1 not implemented")
}

func part2(lines []string) {
	fmt.Println("Part 2 not implemented")
}

func main() {
	input := readInput()
	lines := strings.Split(input, "\n")
	fmt.Println("---PART 1---")
	part1(lines)
	fmt.Println("---PART 2---")
	part2(lines)
}
