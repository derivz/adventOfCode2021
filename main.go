package main

import (
	"adventOfCode2021/days"
	"adventOfCode2021/utils"
)

func main() {
	data := utils.PrepareData(1)
	days.Day1(data)
	days.Day1_1(data)
	data = utils.PrepareData(2)
	days.Day2(data)
	days.Day2_1(data)
	data = utils.PrepareData(3)
	days.Day3(data)
	days.Day3_1(data)
}
