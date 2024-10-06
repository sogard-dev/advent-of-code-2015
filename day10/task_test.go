package day10

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2015/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 2, part1(`1`, 1))
	require.Equal(t, 2, part1(`11`, 1))
	require.Equal(t, 4, part1(`21`, 1))
	require.Equal(t, 6, part1(`1211`, 1))
	require.Equal(t, 6, part1(`111221`, 1))
	require.Equal(t, 492982, part1(utils.GetInput(t, "input.txt"), 40))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 6989950, part1(utils.GetInput(t, "input.txt"), 50))
}
