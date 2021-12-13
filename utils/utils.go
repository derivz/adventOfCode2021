package utils

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ToInt(number string) int {
	val, e := strconv.Atoi(number)
	if e != nil {
		fmt.Println(e)
	}
	return val
}

func IntAbs(num int) int {
	if num > 0 {
		return num
	}
	return -num
}

func ReadLines(fileName string) *[]string {
	var res []string
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, line)
	}
	return &res
}

func Pow(a int, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func SortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func StringSliceMap(slice []string, fn func(string) string) []string {
	res := make([]string, len(slice))
	for i, s := range slice {
		res[i] = fn(s)
	}
	return res
}

func StringIntersection(s1 string, s2 string) string {
	var res []string
	for _, letter := range strings.Split(s2, "") {
		if strings.Contains(s1, letter) {
			res = append(res, letter)
		}
	}
	return strings.Join(res, "")
}

func FormatIntSequences(lines *[]string, separator string) *[][]int {
	var newLines [][]int
	for _, line := range *lines {
		var numbers []int
		for _, el := range strings.Split(line, separator) {
			numbers = append(numbers, ToInt(el))
		}
		newLines = append(newLines, numbers)
	}
	return &newLines
}
