package day14

import (
	"slices"
	"strings"

	"github.com/sogard-dev/advent-of-code-2015/utils"
)

func part1(input string, rounds int) int {
	deers := parse(input)

	longest := 0
	for _, deer := range deers {
		longest = max(deer.atSecond(rounds), longest)
	}

	return longest
}

func part2(input string, rounds int) int {
	deers := parse(input)
	points := make([]int, len(deers))
	for sec := range rounds {
		longest := 1
		longestIndexes := []int{}
		for i, deer := range deers {
			distance := deer.atSecond(sec)
			if distance > longest {
				longest = distance
				longestIndexes = []int{i}
			} else if distance == longest {
				longestIndexes = append(longestIndexes, i)
			}
		}

		for _, k := range longestIndexes {
			points[k]++
		}
	}

	return slices.Max(points)
}

type deer struct {
	speed    int
	duration int
	rest     int
}

func (d deer) atSecond(sec int) int {
	cycle := d.duration + d.rest
	distancePerCycle := d.duration * d.speed
	cycles := sec / cycle
	timeIntoCycle := min(sec%cycle, d.duration)
	return cycles*distancePerCycle + timeIntoCycle*d.speed

}

func parse(input string) []deer {
	ret := []deer{}
	for _, line := range strings.Split(input, "\n") {
		numbers := utils.GetAllNumbers(line)
		speed, duration, rest := numbers[0], numbers[1], numbers[2]

		ret = append(ret, deer{
			speed:    speed,
			duration: duration,
			rest:     rest,
		})

	}
	return ret
}
