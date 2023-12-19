package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Limitations struct {
	red   int
	blue  int
	green int
}

type Draw struct {
	red   int
	blue  int
	green int
}

type Game struct {
	id    int
	draws []Draw
	meta  Meta
}

type Meta struct {
	possible  bool
	red_min   int
	green_min int
	blue_min  int
}

func readInput() string {
	stdin, err := io.ReadAll(os.Stdin)

	if err != nil {
		panic(err)
	}
	str := string(stdin)
	return str
}

func getColors(input string, color string) string {
	rg := regexp.MustCompile(`[0-9]+ ` + regexp.QuoteMeta(color))
	return rg.FindString(input)
}

func parseGameLine(line string, limits Limitations) Game {
	splitted := strings.Split(line, ":")
	draws_splitted := strings.Split(splitted[1], ";")
	game_prefix := splitted[0]

	id, _ := strconv.Atoi(strings.Split(game_prefix, " ")[1])
	draws := []Draw{}
	meta := Meta{}
	red_min, green_min, blue_min := 0, 0, 0
	meta.possible = true
	for _, v := range draws_splitted {
		red, _ := strconv.Atoi(strings.Split(getColors(v, "red"), " ")[0])
		green, _ := strconv.Atoi(strings.Split(getColors(v, "green"), " ")[0])
		blue, _ := strconv.Atoi(strings.Split(getColors(v, "blue"), " ")[0])
		if (red > limits.red) || (green > limits.green) || (blue > limits.blue) {
			meta.possible = false
		}
		if red_min < red {
			red_min = red
		}
		if green_min < green {
			green_min = green
		}
		if blue_min < blue {
			blue_min = blue
		}
		draw := Draw{
			red:   red,
			blue:  blue,
			green: green,
		}
		draws = append(draws, draw)
	}
	meta.blue_min = blue_min
	meta.red_min = red_min
	meta.green_min = green_min
	game := Game{
		id:    id,
		draws: draws,
		meta:  meta,
	}

	// fmt.Println(game)
	return game
}

func part1(lines []string) {
	limits := Limitations{
		red:   12,
		green: 13,
		blue:  14,
	}
	var games []Game
	sum_ids := 0
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		game := parseGameLine(l, limits)
		if game.meta.possible {
			games = append(games, game)
			sum_ids += game.id
		}
	}

	fmt.Println("Sum ids:", sum_ids)
}

func part2(lines []string) {
	var games []Game

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		game := parseGameLine(l, Limitations{})
		games = append(games, game)
	}
	sum := 0
	for _, g := range games {
		sum += (g.meta.red_min * g.meta.green_min * g.meta.blue_min)
	}
	fmt.Println(sum)
}

func main() {
	input := readInput()
	lines := strings.Split(input, "\n")
	fmt.Println("---PART 1---")
	part1(lines)
	fmt.Println("---PART 2---")
	part2(lines)
}
