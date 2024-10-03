package day1

import "strings"

func part1(input string) int {
	return strings.Count(input, "(") - strings.Count(input, ")")
}

func part2(input string) int {
	var floor int = 0

	for i, s := range input {
		switch s {
		case '(':
			floor += 1
		case ')':
			floor -= 1
		default:
			panic("shit")
		}

		if floor == -1 {
			return i + 1
		}
	}
	return 0
}
