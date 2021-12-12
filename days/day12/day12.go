package day12

import (
	"fmt"
	"strings"
	"unicode"
)

var solutions map[string]int

func (steps *Steps) solve(node string, path string) {
	if node == "end" {
		solutions[path] += 1
		return
	}
	for _, to := range (*steps)[node] {
		if unicode.IsUpper(rune(to[0])) || !strings.Contains(path, to) {
			steps.solve(to, path+"-"+to)
		}
	}
}

func (steps *Steps) solve2(node string, path string, visited bool) {
	if node == "end" {
		solutions[path] += 1
		return
	}
	for _, to := range (*steps)[node] {
		switch {
		case to == "start":
		case unicode.IsUpper(rune(to[0])) || !strings.Contains(path, to):
			steps.solve2(to, path+"-"+to, visited)
		case !visited:
			steps.solve2(to, path+"-"+to, true)

		}
	}
}

func Task1(lines *[]string) {
	steps := *formatLines(lines)
	solutions = map[string]int{}
	steps.solve("start", "start")
	fmt.Println(len(solutions))
}

func Task2(lines *[]string) {
	steps := *formatLines(lines)
	solutions = map[string]int{}
	steps.solve2("start", "start", false)
	fmt.Println(len(solutions))
}

type Steps map[string][]string

func formatLines(linesPointer *[]string) *Steps {
	steps := Steps{}
	for _, line := range *linesPointer {
		lineParts := strings.Split(line, "-")
		a := lineParts[0]
		b := lineParts[1]
		steps[a] = append(steps[a], b)
		steps[b] = append(steps[b], a)
	}
	return &steps
}
