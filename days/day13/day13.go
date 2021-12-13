package day13

import (
	"adventOfCode2021/utils"
	"fmt"
	"strings"
)

func (paper *Paper) makeFold(fold Fold) {
	for _, point := range *paper {
		if fold.dir == "x" && point.x > fold.num {
			point.x = fold.num*2 - point.x
		}
		if fold.dir == "y" && point.y > fold.num {
			point.y = fold.num*2 - point.y
		}
	}
}

func Task1(lines *[]string) {
	paper, folds := formatData(lines)
	paper.makeFold((*folds)[0])
	printPaper(paper, true)
}

func Task2(lines *[]string) {
	paper, folds := formatData(lines)
	for _, fold := range *folds {
		paper.makeFold(fold)
	}
	printPaper(paper, false)
}

type Point struct {
	x, y int
}

type Fold struct {
	dir string
	num int
}

type Paper []*Point

func printPaper(paper *Paper, onlyCount bool) {
	var maxx, maxy int
	points := map[Point]bool{}
	for _, pointPointer := range *paper {
		point := *pointPointer
		if point.x > maxx {
			maxx = point.x
		}
		if point.y > maxy {
			maxy = point.y
		}
		points[point] = true
	}
	var board [][]string
	for i := 0; i <= maxy; i++ {
		var row []string
		for j := 0; j <= maxx; j++ {
			if _, prs := points[Point{j, i}]; prs {
				row = append(row, "#")

			} else {
				row = append(row, ".")
			}
		}
		board = append(board, row)
	}
	fmt.Println(strings.Join(make([]string, 30), "-"), len(points))
	if !onlyCount {
		for _, row := range board {
			fmt.Println(strings.Join(row, ""))
		}
	}
}

func formatData(lines *[]string) (*Paper, *[]Fold) {
	var paper Paper
	var folds []Fold
	for _, line := range *lines {
		if strings.Contains(line, ",") {
			parts := strings.Split(line, ",")
			paper = append(
				paper, &Point{
					utils.ToInt(parts[0]),
					utils.ToInt(parts[1]),
				},
			)
		}
		if strings.Contains(line, "fold") {
			parts := strings.Split(line, "=")
			folds = append(
				folds, Fold{
					parts[0][len(parts[0])-1:],
					utils.ToInt(parts[1]),
				},
			)
		}
	}
	return &paper, &folds
}
