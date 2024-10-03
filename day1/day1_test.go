package day1

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2015/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 74, part1(utils.GetInput(t, "day1.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 1795, part2(utils.GetInput(t, "day1.txt")))
}
