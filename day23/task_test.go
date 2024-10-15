package day23

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2015/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 2, part1(`inc b
jio b, +2
tpl b
inc b`))
	require.Equal(t, 307, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 160, part2(utils.GetInput(t, "input.txt")))
}
