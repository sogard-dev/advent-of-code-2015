package day17

import (
	"github.com/sogard-dev/advent-of-code-2015/utils"
)

func part1(input string, storing int) int {
	containers := utils.GetAllNumbers(input)
	return solve(containers, 0, storing, 0, func(int) bool { return true })
}

func part2(input string, storing int) int {
	containers := utils.GetAllNumbers(input)

	smallest := len(containers)
	solve(containers, 0, storing, 0, func(picked int) bool {
		smallest = min(smallest, picked)
		return true
	})

	return solve(containers, 0, storing, 0, func(picked int) bool {
		return picked == smallest
	})
}

func solve(containers []int, picked int, remaining int, moving int, picker func(int) bool) int {
	if remaining == 0 {
		if picker(picked) {
			return 1
		} else {
			return 0
		}
	}

	if moving+1 < len(containers) {
		sum := 0
		if containers[moving] <= remaining {
			sum += solve(containers, picked+1, remaining-containers[moving], moving+1, picker)
		}
		sum += solve(containers, picked, remaining, moving+1, picker)
		return sum
	} else if remaining-containers[moving] == 0 {
		if picker(picked + 1) {
			return 1
		} else {
			return 0
		}
	}

	return 0
}
