package day6

import (
	"regexp"
	"strconv"
	"strings"
)

const SIZE = 1000

type funcs struct {
	toggle func(*[SIZE][SIZE]int, int, int)
	on     func(*[SIZE][SIZE]int, int, int)
	off    func(*[SIZE][SIZE]int, int, int)
}

func part1(input string) int {
	return solve(input, funcs{
		toggle: func(lights *[SIZE][SIZE]int, x, y int) {
			if lights[x][y] == 1 {
				lights[x][y] = 0
			} else {
				lights[x][y] = 1
			}
		},
		on: func(lights *[SIZE][SIZE]int, x, y int) {
			lights[x][y] = 1
		},
		off: func(lights *[SIZE][SIZE]int, x, y int) {
			lights[x][y] = 0
		},
	})
}

func part2(input string) int {
	return solve(input, funcs{
		toggle: func(lights *[SIZE][SIZE]int, x, y int) {
			lights[x][y] += 2
		},
		on: func(lights *[SIZE][SIZE]int, x, y int) {
			lights[x][y] += 1
		},
		off: func(lights *[SIZE][SIZE]int, x, y int) {
			lights[x][y] = max(0, lights[x][y]-1)
		},
	})
}

func solve(input string, f funcs) int {
	lights := [SIZE][SIZE]int{}

	re := regexp.MustCompile("[0-9]+")
	for _, line := range strings.Split(input, "\n") {
		numbers := re.FindAllString(line, 4)
		startX, _ := strconv.Atoi(numbers[0])
		startY, _ := strconv.Atoi(numbers[1])
		endX, _ := strconv.Atoi(numbers[2])
		endY, _ := strconv.Atoi(numbers[3])

		var consumer func(*[SIZE][SIZE]int, int, int)

		if strings.Contains(line, "turn off") {
			consumer = f.off
		} else if strings.Contains(line, "turn on") {
			consumer = f.on
		} else if strings.Contains(line, "toggle") {
			consumer = f.toggle
		} else {
			panic("This should not happen: " + line)
		}

		for x := startX; x <= endX; x++ {
			for y := startY; y <= endY; y++ {
				consumer(&lights, x, y)
			}
		}
	}

	var lightsOn int
	for x := 0; x < SIZE; x++ {
		for y := 0; y < SIZE; y++ {
			lightsOn += lights[x][y]
		}
	}
	return lightsOn
}
