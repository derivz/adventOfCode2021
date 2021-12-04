package day4

import (
	"adventOfCode2021/utils"
	"fmt"
	"strings"
)

const n = 5

type Grid struct {
	Original [][]int
	Match    [5][5]bool
	Map      map[int][2]int
	Win      bool
}

func (grid *Grid) markNumber(num int) bool {
	ints, prs := grid.Map[num]
	if grid.Win {
		// no more marking for winner for part 2
		return false
	}
	if prs {
		i := ints[0]
		j := ints[1]
		if grid.Original[i][j] != num {
			fmt.Println("Error: ", grid.Original[i][j], " != ", num)
		}
		grid.Match[i][j] = true
		row := true
		col := true
		for k := 0; k < n; k++ {
			row = row && grid.Match[i][k]
			col = col && grid.Match[k][j]
		}
		if row || col {
			grid.Win = true
			return true
		}
	}
	return false
}

func (grid *Grid) calculateResult(num int) int {
	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if !grid.Match[i][j] {
				sum += grid.Original[i][j]
			}
		}
	}
	return sum * num
}

func newGrid() *Grid {
	var g Grid
	g.Map = make(map[int][2]int)
	g.Match = [n][n]bool{}
	return &g
}

func formatGrid(lines []string) *Grid {
	grid := newGrid()
	var numbers []int
	for _, line := range lines {
		numbers = []int{}
		strNums := strings.Split(line, " ")
		for _, num := range strNums {
			if len(num) > 0 {
				numbers = append(numbers, utils.ToInt(num))
			}
		}
		grid.Original = append(grid.Original, numbers)
	}
	grid.Map = map[int][2]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			val := grid.Original[i][j]
			v, p := grid.Map[val]
			if p {
				panic(fmt.Sprintf("ALERT!!! at [%d, %d] %d", i, j, v))
			}
			grid.Map[val] = [2]int{i, j}
		}
	}
	return grid
}

func formatData(lines *[]string) (numbers []int, grids []*Grid) {
	first := (*lines)[0]
	numStrs := strings.Split(first, ",")
	for _, numStr := range numStrs {
		if len(numStr) > 0 {
			numbers = append(numbers, utils.ToInt(numStr))
		}
	}

	var part []string
	for _, line := range (*lines)[2:] {
		if len(line) > 0 {
			part = append(part, line)
			if len(part) == n {
				grids = append(grids, formatGrid(part))
				part = []string{}
			}
		}
	}
	return numbers, grids
}

func Task1(lines *[]string) {
	numbers, grids := formatData(lines)

	for _, number := range numbers {
		for _, grid := range grids {
			if grid.markNumber(number) {
				fmt.Println("RESULT IS ", grid.calculateResult(number))
				return
			}
		}
	}
	fmt.Println("No result???")
}

func Task2(lines *[]string) {
	numbers, grids := formatData(lines)
	gridCount := len(grids)
	winCount := 0

	for _, number := range numbers {
		for _, grid := range grids {
			if grid.markNumber(number) {
				winCount += 1
				if winCount == gridCount {
					fmt.Println("RESULT IS ", grid.calculateResult(number))
					return
				}
			}
		}
	}
	fmt.Println("No result???")
}
