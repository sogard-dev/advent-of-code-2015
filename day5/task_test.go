package day5

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2015/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 1, part1("ugknbfddgicrmopn"))
	require.Equal(t, 1, part1("aaa"))
	require.Equal(t, 0, part1("jchzalrnumimnmhp"))
	require.Equal(t, 0, part1("haegwjzuvuyypxyu"))
	require.Equal(t, 0, part1("dvszwmarrgswjxmb"))
	require.Equal(t, 238, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 0, part2("aaa"))
	require.Equal(t, 1, part2("qjhvhtzxzqqjkmpb"))
	require.Equal(t, 1, part2("xxyxx"))
	require.Equal(t, 0, part2("uurcxstgmygtbstg"))
	require.Equal(t, 0, part2("ieodomkazucvgmuy"))
	require.Equal(t, 69, part2(utils.GetInput(t, "input.txt")))
}
