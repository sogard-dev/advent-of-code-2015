package day2

import (
	"strconv"
	"strings"
)

func part1(input string) int {
	var totalSft int
	for _, line := range strings.Split(input, "\n") {
		l, w, h := parse(line)
		totalSft += 2*l*w + 2*w*h + 2*h*l
		totalSft += min(l*w, w*h, h*l)
	}
	return totalSft
}

func part2(input string) int {
	var totalSft int
	for _, line := range strings.Split(input, "\n") {
		l, w, h := parse(line)
		totalSft += 2 * min(l+w, h+w, l+h)
		totalSft += l * w * h
	}
	return totalSft
}

func parse(line string) (int, int, int) {
	spl := strings.Split(line, "x")
	l, _ := strconv.Atoi(spl[0])
	w, _ := strconv.Atoi(spl[1])
	h, _ := strconv.Atoi(spl[2])

	return l, w, h
}
