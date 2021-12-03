package days

import (
	"adventOfCode2021/utils"
	"strings"
)

type Move struct {
	Dir  string
	Step int
}

func formatMove(line string) Move {
	words := strings.Split(line, " ")
	dir := words[0]
	step := utils.ToInt(words[1])
	return Move{dir, step}
}

func Day2(lines *[]string) {
	var moves []Move
	for _, line := range *lines {
		moves = append(moves, formatMove(line))
	}

	var hor, dep int
	for _, move := range moves {
		switch move.Dir {
		case "forward":
			hor += move.Step
		case "up":
			dep -= move.Step
		case "down":
			dep += move.Step
		}
	}
	println(hor * dep)
}

func Day2_1(lines *[]string) {
	var moves []Move
	for _, line := range *lines {
		moves = append(moves, formatMove(line))
	}

	var hor, dep, aim int
	for _, move := range moves {
		switch move.Dir {
		case "forward":
			hor += move.Step
			dep += move.Step * aim
		case "up":
			aim -= move.Step
		case "down":
			aim += move.Step
		}
	}
	println(hor * dep)
}
