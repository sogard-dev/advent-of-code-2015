package day7

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2015/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	testString := `123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i`
	require.Equal(t, 72, part1(testString, "d"))
	require.Equal(t, 507, part1(testString, "e"))
	require.Equal(t, 492, part1(testString, "f"))
	require.Equal(t, 114, part1(testString, "g"))
	require.Equal(t, 65412, part1(testString, "h"))
	require.Equal(t, 65079, part1(testString, "i"))
	require.Equal(t, 123, part1(testString, "x"))
	require.Equal(t, 456, part1(testString, "y"))

	require.Equal(t, 956, part1(utils.GetInput(t, "input.txt"), "a"))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 40149, part2(utils.GetInput(t, "input.txt")))
}
