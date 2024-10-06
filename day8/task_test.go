package day8

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2015/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 3, part1(`"\\x27"`))
	require.Equal(t, 6, part1(`"\\\x27"`))
	require.Equal(t, 5, part1(`"\x27"`))
	require.Equal(t, 3, part1(`"aaa\"aaa"`))
	require.Equal(t, 2, part1(`"abc"`))
	require.Equal(t, 4, part1(`"\\\\"`))
	require.Equal(t, 2, part1(`"j"`))
	require.Equal(t, 4, part1(`"\\\\zkisyjpbzandqikqjqvee"`))
	require.Equal(t, 12, part1(`""
"abc"
"aaa\"aaa"
"\x27"`))
	require.Equal(t, 1333, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 4, part2(`""`))
	require.Equal(t, 4, part2(`"abc"`))
	require.Equal(t, 6, part2(`"aaa\"aaa"`))
	require.Equal(t, 5, part2(`"\x27"`))
	require.Equal(t, 19, part2(`""
"abc"
"aaa\"aaa"
"\x27"`))
	require.Equal(t, 2046, part2(utils.GetInput(t, "input.txt")))
}
