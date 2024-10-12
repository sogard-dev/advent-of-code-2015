package day20

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 8, part1(150))
	require.Equal(t, 786240, part1(34000000))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 831600, part2(34000000))
}
