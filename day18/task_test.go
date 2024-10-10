package day18

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2015/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 4, part1(`.#.#.#
...##.
#....#
..#...
#.#..#
####..`, 4))
	require.Equal(t, 1061, part1(utils.GetInput(t, "input.txt"), 100))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 17, part2(`.#.#.#
...##.
#....#
..#...
#.#..#
####..`, 5))
	require.Equal(t, 1006, part2(utils.GetInput(t, "input.txt"), 100))
}
