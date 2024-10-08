package day14

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2015/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 1120, part1(`Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`, 1000))
	require.Equal(t, 2696, part1(utils.GetInput(t, "input.txt"), 2503))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 689, part2(`Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`, 1000))
	require.Equal(t, 1084, part2(utils.GetInput(t, "input.txt"), 2503))
}
