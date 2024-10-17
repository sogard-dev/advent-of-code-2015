package day24

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2015/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, uint64(99), part1(`1
	2
	3
	4
	5
	7
	8
	9
	10
	11`))
	//require.Equal(t, 11846773891, int(part1(utils.GetInput(t, "input.txt"))))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 80393059, int(part2(utils.GetInput(t, "input.txt"))))
}
