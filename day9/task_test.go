package day9

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2015/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 605, part1(`London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141`))
	require.Equal(t, 141, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 982, part2(`London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141`))
	require.Equal(t, 736, part2(utils.GetInput(t, "input.txt")))
}
