package day11

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2015/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, "ghjaabcc", part1(`ghjaabcc`))
	require.Equal(t, "abcdffaa", part1(`abcdefgh`))
	require.Equal(t, "ghjaabcc", part1(`ghijklmn`))
	require.Equal(t, "cqjxxyzz", part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	newPass := part1(utils.GetInput(t, "input.txt"))
	require.Equal(t, "cqkaabcc", part1(iterate(newPass)))
}
