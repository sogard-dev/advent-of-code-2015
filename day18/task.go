package day18

import (
	"fmt"
	"math/big"
	"math/bits"
	"strings"
)

func printLights(b *big.Int, n int) {
	for y := range n {
		for x := range n {
			if b.Bit(getIndexFromCoord(y, x, n)) == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}

		}
		fmt.Println()
	}
}

func getIndexFromCoord(y, x, n int) int {
	return n*y + x
}

func bitCount(z *big.Int) int {
	var count int
	for _, x := range z.Bits() {
		count += bits.OnesCount(uint(x))
	}
	return count
}

func iterate(prevLights *big.Int, n int, part2 bool) big.Int {
	lights := big.Int{}

	isOn := func(y, x int) uint {
		return prevLights.Bit(getIndexFromCoord(y, x, n))
	}

	set := func(y, x int, v uint) {
		lights.SetBit(&lights, getIndexFromCoord(y, x, n), v)
	}

	if part2 {
		set(1, 1, 1)
		set(1, n-2, 1)
		set(n-2, 1, 1)
		set(n-2, n-2, 1)
	}

	for y := 1; y < n-1; y++ {
		for x := 1; x < n-1; x++ {
			var neighbours uint
			for dy := -1; dy < 2; dy++ {
				for dx := -1; dx < 2; dx++ {
					if dx != 0 || dy != 0 {
						neighbours += isOn(y+dy, x+dx)
					}
				}
			}
			if isOn(y, x) == 1 {
				if neighbours == 2 || neighbours == 3 {
					set(y, x, 1)
				}
			} else {
				if neighbours == 3 {
					set(y, x, 1)
				}
			}
		}
	}

	return lights
}

func parse(input string) (int, big.Int) {
	lines := strings.Split(input, "\n")
	n := len(lines) + 2
	lights := big.Int{}
	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				lights.SetBit(&lights, getIndexFromCoord(y+1, x+1, n), 1)
			}
		}
	}
	return n, lights
}

func part1(input string, steps int) int {
	n, lights := parse(input)

	for range steps {
		lights = iterate(&lights, n, false)
	}

	return bitCount(&lights)
}

func part2(input string, steps int) int {
	n, lights := parse(input)
	lights.SetBit(&lights, getIndexFromCoord(1, 1, n), 1)
	lights.SetBit(&lights, getIndexFromCoord(n-2, 1, n), 1)
	lights.SetBit(&lights, getIndexFromCoord(1, n-2, n), 1)
	lights.SetBit(&lights, getIndexFromCoord(n-2, n-2, n), 1)

	for range steps {
		lights = iterate(&lights, n, true)

	}

	return bitCount(&lights)
}
