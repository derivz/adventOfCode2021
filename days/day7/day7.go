package day7

import (
	"adventOfCode2021/utils"
	"fmt"
	"strings"
)

func formatData(lines *[]string) *[]int {
	var numbers []int
	for _, el := range strings.Split((*lines)[0], ",") {
		numbers = append(numbers, utils.ToInt(el))
	}
	return &numbers
}

func intMax(nums *[]int) int {
	res := 0
	for _, num := range *nums {
		if num > res {
			res = num
		}
	}
	return res
}

func calculateTotal(nums *[]int, dest int, part2 bool) int {
	total := 0
	for _, num := range *nums {
		count := utils.IntAbs(num - dest)
		if part2 {
			total += count * (count + 1) / 2
		} else {
			total += count
		}
	}
	return total
}

func findDestination(nums *[]int, part2 bool) int {
	left := 0
	right := intMax(nums)
	mid := right / 2
	for left <= right {
		value := calculateTotal(nums, mid, part2)
		valueLeft := calculateTotal(nums, mid-1, part2)
		if valueLeft < value {
			right = mid - 1
			mid = (left + right) / 2
			continue
		}

		valueRight := calculateTotal(nums, mid+1, part2)
		if valueRight < value {
			left = mid + 1
			mid = (left + right) / 2
			continue
		}

		return mid
	}
	return left
}

func Task1(lines *[]string) {
	nums := formatData(lines)
	dest := findDestination(nums, false)
	fmt.Println(dest, calculateTotal(nums, dest, false))
}

func Task2(lines *[]string) {
	nums := formatData(lines)
	dest := findDestination(nums, true)
	fmt.Println(dest, calculateTotal(nums, dest, true))
}
