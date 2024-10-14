package day21

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 78, part1(
		unit{
			hitpoints: 100,
			armor:     0,
			damage:    0,
		},
		unit{
			hitpoints: 104,
			damage:    8,
			armor:     1,
		}))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 148, part2(
		unit{
			hitpoints: 100,
			armor:     0,
			damage:    0,
		},
		unit{
			hitpoints: 104,
			damage:    8,
			armor:     1,
		}))
}
