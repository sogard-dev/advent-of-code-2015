package day2

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2015/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 58, part1("2x3x4"))
	require.Equal(t, 43, part1("1x1x10"))
	require.Equal(t, 1586300, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 34, part2("2x3x4"))
	require.Equal(t, 14, part2("1x1x10"))
	require.Equal(t, 3737498, part2(utils.GetInput(t, "input.txt")))
}
