package day12

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2015/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 6, part1(`[1,2,3]`))
	require.Equal(t, 3, part1(`[[[3]]]`))
	require.Equal(t, 0, part1(`{"a":[-1,1]}`))
	require.Equal(t, 0, part1(`[-1,{"a":1}]`))
	require.Equal(t, 0, part1(`[]`))
	require.Equal(t, 0, part1(`{}`))
	require.Equal(t, 111754, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 4, part2(`{"test":[1,{"c":"red","b":2},3]}`))
	require.Equal(t, 65402, part2(utils.GetInput(t, "input.txt")))
}
