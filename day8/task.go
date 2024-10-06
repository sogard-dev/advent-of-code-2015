package day8

import "strings"

func part1(input string) int {
	var charsInString int
	var charsInMemory int
	for _, line := range strings.Split(input, "\n") {
		charsInString += len(line)

		i := 1
		length := 0
		var prev byte
		for i < len(line)-1 {
			thisChar := line[i]
			toSetPrev := thisChar

			if prev == '\\' {
				if thisChar == '"' {
					toSetPrev = 0
				} else if thisChar == '\\' {
					toSetPrev = 0
				} else if thisChar == 'x' {
					toSetPrev = 0
					i += 2
				} else {
					panic("Unknown escape")
				}
			} else {
				length += 1
			}

			prev = toSetPrev
			i++
		}

		charsInMemory += length
	}
	return charsInString - charsInMemory
}

func part2(input string) int {
	var charsInString int
	var charsEscaped int
	for _, line := range strings.Split(input, "\n") {
		charsInString += len(line)

		i := 0
		length := 0
		for i < len(line) {
			thisChar := line[i]

			switch thisChar {
			case '"', '\\':
				length += 2
			default:
				length += 1
			}

			i++
		}

		charsEscaped += length + 2
	}
	return charsEscaped - charsInString
}
