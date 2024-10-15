package day22

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 900, part1(50, 500, 51, 9))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 1216, part2(50, 500, 51, 9))
}
