package main

import (
	"adventOfCode2021/days/day15"
	"adventOfCode2021/utils"
)

func main() {
	data := utils.PrepareData(15)
	println(data)
	day15.Task1(data)
	day15.Task2(data)
}
