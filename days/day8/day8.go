package day8

import (
	"adventOfCode2021/utils"
	"fmt"
	"strings"
)

type Line struct {
	Sample    []string
	Riddle    []string
	CodeToInt map[string]int
	Codes     [10]string
}

func newLine(row string) *Line {
	var line Line
	parts := strings.Split(row, " | ")
	line.Sample = utils.StringSliceMap(strings.Split(parts[0], " "), utils.SortString)
	line.Riddle = utils.StringSliceMap(strings.Split(parts[1], " "), utils.SortString)
	line.CodeToInt = make(map[string]int)
	return &line
}

func (line *Line) MarkCode(code string, number int) {
	line.Codes[number] = code
	line.CodeToInt[code] = number
}

func (line *Line) DecodeAnswer() int {
	var res int
	for i, code := range line.Riddle {
		res += line.CodeToInt[code] * utils.Pow(10, len(line.Riddle)-1-i)
	}
	return res
}

func (line *Line) solveRiddle() int {
	for _, code := range line.Sample {
		switch len(code) {
		case 2:
			line.MarkCode(code, 1)
		case 3:
			line.MarkCode(code, 7)
		case 4:
			line.MarkCode(code, 4)
		case 7:
			line.MarkCode(code, 8)
		}
	}
	for _, code := range line.Sample {
		if len(code) == 5 {
			if utils.StringIntersection(code, line.Codes[1]) == line.Codes[1] {
				line.MarkCode(code, 3)
			} else if len(utils.StringIntersection(code, line.Codes[4])) == 3 {
				line.MarkCode(code, 5)
			} else {
				line.MarkCode(code, 2)
			}
		}
	}
	for _, code := range line.Sample {
		if len(code) == 6 {
			if utils.StringIntersection(code, line.Codes[3]) == line.Codes[3] {
				line.MarkCode(code, 9)
			} else if utils.StringIntersection(code, line.Codes[7]) == line.Codes[7] {
				line.MarkCode(code, 0)
			} else {
				line.MarkCode(code, 6)
			}
		}
	}
	return line.DecodeAnswer()
}

func formatData(rows *[]string) *[]Line {
	var lines []Line
	for _, row := range *rows {
		lines = append(lines, *newLine(row))
	}
	return &lines
}

func Task1(rows *[]string) {
	lines := *formatData(rows)
	res := 0
	for _, line := range lines {
		for _, code := range line.Riddle {
			switch len(code) {
			case 2,3,4,7:
				res += 1
			}
		}
	}
	fmt.Println(res)
}

func Task2(rows *[]string) {
	lines := *formatData(rows)
	res := 0
	for _, line := range lines {
		sol := line.solveRiddle()
		res += sol
	}
	fmt.Println( res)
}
