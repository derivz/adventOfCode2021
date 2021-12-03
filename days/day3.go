package days

import (
	"strconv"
)

const n = 12

func Day3(lines *[]string) {
	var c1 [n]int
	var c0 [n]int

	var cmin, cmax string
	for _, move := range *lines {
		for i := 0; i < n; i++ {
			if move[i] == '1' {
				c1[i] += 1
			} else {
				c0[i] += 1
			}
		}
	}
	for i := 0; i < n; i++ {
		if c1[i] >= c0[i] {
			cmax += "1"
			cmin += "0"
		} else {
			cmax += "0"
			cmin += "1"
		}
	}
	ma, _ := strconv.ParseInt(cmax, 2, 64)
	mi, _ := strconv.ParseInt(cmin, 2, 64)
	println(ma*mi)
}

func Day3_1(lines *[]string) {
	var c1 [n]int
	var c0 [n]int

	var cmin, cmax string
	for i := 0; i < n; i++ {
		for _, move := range *lines {
			if move[:i] == cmax {
				if move[i] == '1' {
					c1[i] += 1
				} else {
					c0[i] += 1
				}
			}
		}
		if c1[i] >= c0[i] {
			cmax += "1"
		} else {
			cmax += "0"
		}
	}
	for i := 0; i < n; i++ {
		c1[i] = 0
		c0[i] = 0
		for _, move := range *lines {
			if move[:i] == cmin {
				if move[i] == '1' {
					c1[i] += 1
				} else {
					c0[i] += 1
				}
			}
		}
		if c0[i] == 0 {
			cmin += "1"
		} else if c1[i] >= c0[i] || c1[i] == 0 {
			cmin += "0"
		} else {
			cmin += "1"
		}
	}
	ma, _ := strconv.ParseInt(cmax, 2, 64)
	mi, _ := strconv.ParseInt(cmin, 2, 64)
	println(ma*mi)
}
