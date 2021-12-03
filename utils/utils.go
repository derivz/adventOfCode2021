package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ToInt(number string) int {
	val, e := strconv.Atoi(number)
	if e != nil {
		fmt.Println(e)
	}
	return val
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
