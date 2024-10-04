package day3

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2015/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 2, part1(">"))
	require.Equal(t, 4, part1("^>v<"))
	require.Equal(t, 2, part1("^v^v^v^v^v"))
	require.Equal(t, 2081, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 3, part2("^v"))
	require.Equal(t, 3, part2("^>v<"))
	require.Equal(t, 11, part2("^v^v^v^v^v"))
	require.Equal(t, 2341, part2(utils.GetInput(t, "input.txt")))
}
