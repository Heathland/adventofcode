package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var gameDelimiter = ":"
var setDelimiter = ";"
var dataDelimiter = ","
var valueDelimiter = " "

type (
	Set struct {
		Blue  int
		Red   int
		Green int
	}
	Game []Set
)

func readSet(s string) (set Set) {
	data := strings.Split(s, dataDelimiter)
	for _, value := range data {
		v := strings.Split(strings.TrimSpace(value), valueDelimiter)
		var err error
		switch strings.ToLower(v[1]) {
		case "red":
			set.Red, err = strconv.Atoi(v[0])
		case "green":
			set.Green, err = strconv.Atoi(v[0])
		case "blue":
			set.Blue, err = strconv.Atoi(v[0])
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	return
}

func readGame(s string) (g Game) {
	gameInfo := strings.Split(s, gameDelimiter)[1]
	sets := strings.Split(gameInfo, setDelimiter)
	g = make([]Set, len(sets))

	for k, set := range sets {
		g[k] = readSet(set)
	}
	return
}

func validateGame(g Game, vs Set) bool {
	for _, set := range g {
		if set.Red > vs.Red {
			return false
		}
		if set.Green > vs.Green {
			return false
		}
		if set.Blue > vs.Blue {
			return false
		}
	}
	return true
}

func getMinNeededSum(g Game) int {
	minSet := Set{}
	for _, set := range g {
		if set.Red > minSet.Red {
			minSet.Red = set.Red
		}
		if set.Green > minSet.Green {
			minSet.Green = set.Green
		}
		if set.Blue > minSet.Blue {
			minSet.Blue = set.Blue
		}
	}

	return minSet.Red * minSet.Green * minSet.Blue
}

func main() {
	lines := strings.Split(input, "\n")

	games := make([]Game, len(lines))
	for k, line := range lines {
		games[k] = readGame(line)
	}

	var totalPossible int
	var totalMinNeeded int
	for k, game := range games {
		validationSet := Set{Red: 12, Green: 13, Blue: 14}
		if validateGame(game, validationSet) {
			totalPossible += k + 1
		}
		totalMinNeeded += getMinNeededSum(game)
	}

	fmt.Printf("Result part 1: %d\n", totalPossible)
	fmt.Printf("Result part 2: %d\n", totalMinNeeded)
}
