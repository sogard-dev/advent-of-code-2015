package day13

import (
	"strings"

	"github.com/sogard-dev/advent-of-code-2015/utils"
)

func part1(input string) int {
	return solve(input, []string{})
}

func part2(input string) int {
	return solve(input, []string{"me"})
}

func solve(input string, extraNames []string) int {
	m := map[string]map[string]int{}

	for _, line := range strings.Split(input, "\n") {
		spl := strings.Split(line, " ")
		name1 := spl[0]
		name2 := strings.ReplaceAll(spl[10], ".", "")
		change := spl[2]
		num := utils.GetAllNumbers(line)[0]
		if change == "lose" {
			num = -num
		}
		if _, ok := m[name1]; !ok {
			m[name1] = map[string]int{}
		}
		m[name1][name2] = num
	}

	names := []string{}
	for k := range m {
		names = append(names, k)
	}
	names = append(names, extraNames...)

	maxScore := 0
	for p := make([]int, len(names)); p[0] < len(p); utils.NextPerm(p) {
		perm := utils.GetPerm(names, p)
		score := getScore(m, perm)
		maxScore = max(maxScore, score)
	}

	return maxScore
}

func getScore(m map[string]map[string]int, n []string) int {
	score := 0
	size := len(n)
	for i := 0; i < size; i++ {
		prevName := (i + size - 1) % size
		nextName := (i + size + 1) % size
		me := n[i]
		next := n[nextName]
		prev := n[prevName]

		score += m[me][prev]
		score += m[me][next]
	}

	return score
}
