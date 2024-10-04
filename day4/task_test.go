package day4

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2015/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 609043, part1("abcdef"))
	require.Equal(t, 1048970, part1("pqrstuv"))
	require.Equal(t, 346386, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 9958218, part2(utils.GetInput(t, "input.txt")))
}
