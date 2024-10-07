package day11

import "strings"

func part1(input string) string {
	for {
		twoPairs := hasTwoPairs(input)
		abc := hasAbc(input)
		noBad := hasGoodLetter(input)
		if twoPairs && abc && noBad {
			return input
		}

		input = iterate(input)
	}
}

func iterate(input string) string {
	ret := []rune(input)
	for i := len(input) - 1; i >= 0; i-- {
		r := input[i]
		if r == 'z' {
			ret[i] = 'a'
		} else {
			ret[i] = rune(r + 1)
			return string(ret)
		}
	}
	panic("We are out of chars! ")
}

func hasTwoPairs(input string) bool {
	for i := range len(input) - 1 {
		if input[i] == input[i+1] {
			for j := i + 2; j < len(input)-1; j++ {
				if input[i] != input[j] && input[j] == input[j+1] {
					return true
				}
			}
		}
	}
	return false
}

func hasAbc(input string) bool {
	for i := range len(input) - 2 {
		next := input[i] + 1
		nextNext := input[i] + 2
		if next == input[i+1] && nextNext == input[i+2] {
			return true
		}
	}
	return false
}

func hasGoodLetter(input string) bool {
	return !strings.ContainsAny(input, "iol")
}

func part2(input string) int {
	return 0
}
