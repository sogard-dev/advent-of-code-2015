package day25

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2015/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 32451966, part1(4, 2))
	require.Equal(t, 2650453, part1(2978, 3083))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 0, part2(``))
	require.Equal(t, 0, part2(utils.GetInput(t, "input.txt")))
}
