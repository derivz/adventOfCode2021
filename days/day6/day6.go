package day6

import (
	"adventOfCode2021/utils"
	"fmt"
	"strings"
)

func calculateFishPopulation(daysTillBreed int, daysLeft int, memory *map[int]int) int {
	if daysLeft < 0 {
		return 0
	}
	result := 1
	daysLeft -= daysTillBreed

	memorisedResult, prs := (*memory)[daysLeft]
	if prs {
		return memorisedResult
	}
	for i := 0; i <= daysLeft/7; i++ {
		result += calculateFishPopulation(8, daysLeft-i*7-1, memory)
	}
	(*memory)[daysLeft] = result
	return result
}

func formatData(lines *[]string) *[]int {
	var numbers []int
	for _, el := range strings.Split((*lines)[0], ",") {
		numbers = append(numbers, utils.ToInt(el))
	}
	return &numbers
}

func Task1(lines *[]string) {
	memory := make(map[int]int)
	nums := *formatData(lines)
	total := 0
	for _, num := range nums {
		total += calculateFishPopulation(num, 80, &memory)
	}
	fmt.Println(total)
}

func Task2(lines *[]string) {
	memory := make(map[int]int)
	nums := *formatData(lines)
	total := 0
	for _, num := range nums {
		total += calculateFishPopulation(num, 256, &memory)
	}
	fmt.Println(total)
}
