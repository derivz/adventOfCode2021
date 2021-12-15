package day15

import (
	"adventOfCode2021/utils"
	"container/heap"
	"fmt"
)

func Task1(lines *[]string) {
	fmt.Println(solve(lines, 1))
}

func Task2(lines *[]string) {
	fmt.Println(solve(lines, 5))
}

func solve(lines *[]string, multiplier int) int {
	grid := IntGrid(*utils.FormatIntSequences(lines, ""))
	visited := make(map[[2]int]bool)
	myHeap := &PointHeap{Point{0, 0, 0, &grid, &visited}}
	heap.Init(myHeap)
	for myHeap.Len() > 0 {
		point := heap.Pop(myHeap).(Point)
		if point.isLast(multiplier) {
			return point.totalDistance
		}
		point.pushAdjustedPoints(myHeap, multiplier)
	}
	return 0
}

func (grid *IntGrid) getValue(i, j int) int {
	x, xMove := i%grid.height(1), i/grid.height(1)
	y, yMove := j%grid.width(1), j/grid.width(1)
	return ((*grid)[y][x]+xMove+yMove-1)%9 + 1
}

func (p Point) isLast(multiplier int) bool {
	return p.y == len(*p.gridRef)*multiplier-1 &&
		p.x == len((*p.gridRef)[0])*multiplier-1

}

func (p Point) pushAdjustedPoints(pointHeap *PointHeap, multiplier int) {
	for _, modifier := range []int{-1, 1} {
		for i := 0; i < 2; i++ {
			x := p.x
			y := p.y
			if i > 0 {
				x += modifier
			} else {
				y += modifier
			}
			if y < 0 || y > p.gridRef.height(multiplier)-1 ||
				x < 0 || x > p.gridRef.width(multiplier)-1 {
				continue
			}
			if _, prs := (*p.visitedRef)[[2]int{x, y}]; prs {
				continue
			}
			(*p.visitedRef)[[2]int{x, y}] = true
			heap.Push(pointHeap, Point{
				x:             x,
				y:             y,
				totalDistance: p.totalDistance + p.gridRef.getValue(x, y),
				gridRef:       p.gridRef,
				visitedRef:    p.visitedRef,
			})
		}
	}
}

type IntGrid [][]int

func (grid *IntGrid) height(multiplier int) int {
	return len(*grid) * multiplier
}

func (grid *IntGrid) width(multiplier int) int {
	return len((*grid)[0]) * multiplier
}

type Point struct {
	x             int
	y             int
	totalDistance int
	gridRef       *IntGrid
	visitedRef    *map[[2]int]bool
}

type PointHeap []Point

func (p PointHeap) Len() int {
	return len(p)
}

func (p PointHeap) Less(i, j int) bool {
	return p[i].totalDistance < p[j].totalDistance
}

func (p PointHeap) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PointHeap) Push(x interface{}) {
	*p = append(*p, x.(Point))
}

func (p *PointHeap) Pop() interface{} {
	x := (*p)[p.Len()-1]
	*p = (*p)[0 : p.Len()-1]
	return x
}
