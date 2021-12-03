package day1

import "adventOfCode2021/utils"

func Task1(lines *[]string) {
	var data []int
	for _, line := range *lines {
		data = append(data, utils.ToInt(line))
	}
	count := 0
	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			count += 1
		}
	}
	println(count)
}

func Task2(lines *[]string) {
	var data []int
	for _, line := range *lines {
		data = append(data, utils.ToInt(line))
	}
	count := 0
	for i := 3; i < len(data); i++ {
		if data[i] > data[i-3] {
			count += 1
		}
	}
	println(count)
}
