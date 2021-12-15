package day14

import (
	"fmt"
	"math"
	"strings"
)

func (moves *Moves) makeMove(pairs *Pairs) *Pairs {
	newPairs := make(Pairs)
	for pair, value := range *pairs {
		if letter, prs := (*moves)[pair]; prs {
			newPairs[string(pair[0])+letter] += value
			newPairs[letter+string(pair[1])] += value
		} else {
			newPairs[pair] += value
		}
	}
	return &newPairs

}

func (pairs *Pairs) countRunes(start string) int {
	counter := make(map[uint8]int)
	counter[start[0]] += 1
	counter[start[len(start)-1]] += 1
	for p, v := range *pairs {
		counter[p[0]] += v
		counter[p[1]] += v
	}
	min, max := math.MaxInt64, 0
	for _, v := range counter {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return (max - min) / 2
}

func Task1(lines *[]string) {
	start, moves, pairsP := formatData(lines)
	for i := 0; i < 10; i++ {
		pairsP = moves.makeMove(pairsP)
	}
	fmt.Println(pairsP.countRunes(start))
}

func Task2(lines *[]string) {
	start, moves, pairsP := formatData(lines)
	for i := 0; i < 40; i++ {
		pairsP = moves.makeMove(pairsP)
	}
	fmt.Println(pairsP.countRunes(start))
}

func formatData(linesPointer *[]string) (string, *Moves, *Pairs) {
	start := (*linesPointer)[0]
	moves := make(Moves)
	pairs := make(Pairs)
	for _, line := range *linesPointer {
		if strings.Contains(line, "->") {
			parts := strings.Split(line, " -> ")
			moves[parts[0]] = parts[1]
		}

	}
	for i := 1; i < len(start); i++ {
		pairs[start[i-1:i+1]] += 1
	}
	return start, &moves, &pairs
}

type Moves map[string]string
type Pairs map[string]int
