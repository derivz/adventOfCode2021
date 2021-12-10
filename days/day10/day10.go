package day10

import (
	"fmt"
	"sort"
	"strings"
)

func diagnoseLine(line string) (stackPointer *[]rune, corruptedChar rune) {
	var stack []rune
	for _, symbol := range line {
		if strings.ContainsRune("([{<", symbol) {
			stack = append(stack, symbol)
		} else if len(stack) > 0 && stack[len(stack)-1] == getMatchingChar(symbol) {
			stack = stack[:len(stack)-1]
		} else {
			return &stack, symbol
		}
	}
	return &stack, 0
}

func Task1(lines *[]string) {
	points := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	res := 0
	for _, line := range *lines {
		_, corrupted := diagnoseLine(line)
		if corrupted > 0 {
			res += points[corrupted]
		}
	}
	fmt.Println("task 1: ", res)
}

func Task2(lines *[]string) {
	points := map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}
	var results []int
	for _, line := range *lines {
		stackPointer, corrupted := diagnoseLine(line)
		if corrupted == 0 {
			res := 0
			stack := *stackPointer
			for i := len(stack) - 1; i >= 0; i-- {
				res = res*5 + points[(stack)[i]]
			}
			results = append(results, res)
		}
	}
	sort.Ints(results)

	fmt.Println("task 2: ", results[len(results)/2])
}

func getMatchingChar(s rune) rune {
	match := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}
	v, _ := match[s]
	return v
}
