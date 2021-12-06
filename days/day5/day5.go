package day5

import (
	"adventOfCode2021/utils"
	"fmt"
	"math"
	"strings"
)

const n = 1000

type Move struct {
	From Point
	To   Point
}

type Point struct {
	X int
	Y int
}

type Grid [n][n]int

func (grid *Grid) countCrosses() int {
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 1 {
				count++
			}
		}
	}
	return count
}

func newPoint(pointStr string) *Point {
	var p Point
	coords := strings.Split(pointStr, ",")
	p.X = utils.ToInt(coords[0])
	p.Y = utils.ToInt(coords[1])
	return &p
}

func abs(num int) int {
	return int(math.Abs(float64(num)))
}

func boolToDStep(val bool) int {
	if val {
		return 1
	}
	return -1
}

func formatData(lines *[]string) []Move {
	var moves []Move
	for _, line := range *lines {
		points := strings.Split(line, " -> ")
		fromPoint := newPoint(points[0])
		toPoint := newPoint(points[1])
		moves = append(moves, Move{*fromPoint, *toPoint})
	}
	return moves
}

func (move *Move) isHorizontal() bool {
	return move.From.Y == move.To.Y
}

func (move *Move) markHorizontal(grid *Grid) {
	diff := abs(move.From.X - move.To.X)
	step := boolToDStep(move.From.X < move.To.X)
	for i := 0; i <= diff; i++ {
		grid[move.From.X+step*i][move.From.Y] += 1
	}
}

func (move *Move) isVertical() bool {
	return move.From.X == move.To.X
}

func (move *Move) markVertical(grid *Grid) {
	diff := abs(move.From.Y - move.To.Y)
	step := boolToDStep(move.From.Y < move.To.Y)
	for i := 0; i <= diff; i++ {
		grid[move.From.X][move.From.Y+step*i] += 1
	}
}

func (move *Move) isDiagonal() bool {
	xDiff := abs(move.From.X - move.To.X)
	yDiff := abs(move.From.Y - move.To.Y)
	return xDiff == yDiff
}

func (move *Move) markDiagonal(grid *Grid) {
	diffX := abs(move.From.X - move.To.X)
	stepX := boolToDStep(move.From.X < move.To.X)
	stepY := boolToDStep(move.From.Y < move.To.Y)
	for i := 0; i <= diffX; i++ {
		(*grid)[move.From.X+stepX*i][move.From.Y+stepY*i] += 1
	}
}

func Task1(lines *[]string) {
	grid := Grid{}
	moves := formatData(lines)

	for _, move := range moves {
		if move.isVertical() {
			move.markVertical(&grid)
		} else if move.isHorizontal() {
			move.markHorizontal(&grid)
		}
	}

	fmt.Println("result - ", grid.countCrosses())
}

func Task2(lines *[]string) {
	grid := Grid{}
	moves := formatData(lines)

	for _, move := range moves {
		if move.isVertical() {
			move.markVertical(&grid)
		} else if move.isHorizontal() {
			move.markHorizontal(&grid)
		} else if move.isDiagonal() {
			move.markDiagonal(&grid)
		}
	}

	fmt.Println("result - ", grid.countCrosses())
}
