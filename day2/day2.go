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
	id       int
	draws    []Draw
	possible bool
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
	possible := true
	for _, v := range draws_splitted {
		draw := Draw{}
		draw.red, _ = strconv.Atoi(strings.Split(getColors(v, "red"), " ")[0])
		draw.green, _ = strconv.Atoi(strings.Split(getColors(v, "green"), " ")[0])
		draw.blue, _ = strconv.Atoi(strings.Split(getColors(v, "blue"), " ")[0])
		if (draw.red > limits.red) || (draw.green > limits.green) || (draw.blue > limits.blue) {
			possible = false
		}
		draws = append(draws, draw)
	}

	game := Game{
		id:       id,
		draws:    draws,
		possible: possible,
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
		if game.possible {
			games = append(games, game)
			sum_ids += game.id
		}
	}

	fmt.Println(sum_ids)

}

func main() {
	input := readInput()
	lines := strings.Split(input, "\n")
	part1(lines)
}
