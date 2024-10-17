package day25

import (
	"fmt"
	"strings"
)

func getIndex(row, column int) int {
	s, e, r := 0, 0, 1

	for {
		s = e + 1
		e += r
		r += 1

		if r == row+column {
			return s + column - 1
		}
	}
}

func part1(row, column int) int {
	seq := getIndex(row, column)
	n := 20151125
	for i := 1; i < seq; i++ {
		n = n * 252533 % 33554393
	}
	return n
}

func part2(input string) int {
	for _, line := range strings.Split(input, "\n") {
		fmt.Println(line)
	}
	return 0
}
