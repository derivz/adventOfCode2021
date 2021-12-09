package day9

import (
	"adventOfCode2021/utils"
	"fmt"
	"sort"
	"strings"
)

func (f *Field) findLowPoints() []Point {
	var res []Point
	for i, line := range f.grid {
		for j, num := range line {
			switch {
			case j > 0 && num >= f.grid[i][j-1]:
			case j < f.height-1 && num >= f.grid[i][j+1]:
			case i > 0 && num >= f.grid[i-1][j]:
			case i < f.width-1 && num >= f.grid[i+1][j]:
			default:
				res = append(res, Point{i, j})
			}
		}
	}
	return res
}

func (f *Field) getBasinSize(p Point) int {
	if f.grid[p.x][p.y] == 9 {
		return 0
	}
	if _, visited := f.visited[p]; visited {
		return 0
	}
	res := 1
	if p.y > 0 && f.grid[p.x][p.y] < f.grid[p.x][p.y-1] {
		res += f.getBasinSize(Point{p.x, p.y - 1})
	}
	if p.y < f.height-1 && f.grid[p.x][p.y] < f.grid[p.x][p.y+1] {
		res += f.getBasinSize(Point{p.x, p.y + 1})
	}
	if p.x > 0 && f.grid[p.x][p.y] < f.grid[p.x-1][p.y] {
		res += f.getBasinSize(Point{p.x - 1, p.y})
	}
	if p.x < f.width-1 && f.grid[p.x][p.y] < f.grid[p.x+1][p.y] {
		res += f.getBasinSize(Point{p.x + 1, p.y})
	}
	f.visited[p] = true
	return res
}

func Task1(lines *[]string) {
	field := formatData(lines)
	lowPoints := field.findLowPoints()
	res := 0
	for _, lowPoint := range lowPoints {
		res += 1 + field.grid[lowPoint.x][lowPoint.y]
	}
	fmt.Println(res)
}

func Task2(lines *[]string) {
	field := formatData(lines)
	lowPoints := field.findLowPoints()
	var basins []int
	for _, point := range lowPoints {
		basins = append(basins, field.getBasinSize(point))
	}
	sort.Ints(basins)
	n := len(basins)
	fmt.Println(basins[n-1] * basins[n-2] * basins[n-3])
}

// boring crap below

type Point struct {
	x, y int
}

type Field struct {
	grid    [][]int
	width   int
	height  int
	visited map[Point]bool
}

func newField(width int, height int) *Field {
	var f Field
	f.grid = [][]int{}
	f.width = width
	f.height = height
	f.visited = map[Point]bool{}
	return &f
}

func formatData(lines *[]string) *Field {
	width := len(*lines)
	height := len((*lines)[0])
	field := *newField(width, height)
	for _, line := range *lines {
		var numbers []int
		for _, num := range strings.Split(line, "") {
			numbers = append(numbers, utils.ToInt(num))
		}
		field.grid = append(field.grid, numbers)
	}
	return &field
}
