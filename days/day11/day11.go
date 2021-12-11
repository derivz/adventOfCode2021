package day11

import (
	"adventOfCode2021/utils"
	"fmt"
	"strings"
)

const n = 10

type Grid [n][n]int

func (g *Grid) markPoint(i int, j int) {
	if i < 0 || i > 9 || j < 0 || j > 9 {
		return
	}
	g[i][j] += 1
	if g[i][j] == n {
		for a := -1; a < 2; a++ {
			for b := -1; b < 2; b++ {
				g.markPoint(i+a, j+b)
			}
		}
	}
}

func Task1(lines *[]string) {
	grid := *formatLines(lines)
	res := 0
	for step := 0; step < 100; step++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				grid.markPoint(i, j)
			}
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] > 9 {
					res++
					grid[i][j] = 0
				}
			}
		}
	}
	fmt.Println("task 1: ", res)
}

func Task2(lines *[]string) {
	grid := *formatLines(lines)
	step := 0
	for true {
		step++
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				grid.markPoint(i, j)
			}
		}
		res := 0
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] > 9 {
					res++
					grid[i][j] = 0
				}
			}
		}
		if res == n*n {
			break
		}
	}
	fmt.Println("task 2: ", step)
}

func formatLines(linesPointer *[]string) *Grid {
	var grid Grid
	for i, line := range *linesPointer {
		for j, el := range strings.Split(line, "") {
			grid[i][j] = utils.ToInt(el)
		}
	}
	return &grid
}
