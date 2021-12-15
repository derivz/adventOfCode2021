package main

import (
	"adventOfCode2021/days/day14"
	"adventOfCode2021/utils"
)

func main() {
	data := utils.PrepareData(14)
	println(data)
	day14.Task1(data)
	day14.Task2(data)
}
