package day6

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2015/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 1000000, part1("turn on 0,0 through 999,999"))
	require.Equal(t, 1000, part1("toggle 0,0 through 999,0"))
	require.Equal(t, 0, part1("turn off 499,499 through 500,500"))
	require.Equal(t, 400410, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 1, part2("turn on 0,0 through 0,0"))
	require.Equal(t, 2000000, part2("toggle 0,0 through 999,999"))
	require.Equal(t, 0, part2(utils.GetInput(t, "input.txt")))
}
