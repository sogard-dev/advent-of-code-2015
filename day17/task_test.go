package day17

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2015/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 4, part1(`20 15 10 5 5`, 25))
	require.Equal(t, 654, part1(utils.GetInput(t, "input.txt"), 150))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 3, part2(`20 15 10 5 5`, 25))
	require.Equal(t, 57, part2(utils.GetInput(t, "input.txt"), 150))
}
