package day5

import "strings"

func part1(input string) int {
	var result int
	for _, word := range strings.Split(input, "\n") {
		if isNicePart1(word) {
			result += 1
		}
	}
	return result
}

func isNicePart1(word string) bool {
	badArr := [...]string{"ab", "cd", "pq", "xy"}
	for _, bad := range badArr {
		if strings.Contains(word, bad) {
			return false
		}
	}

	vowelCount := 0
	previousRune := ' '
	letterTwice := false

	for _, r := range word {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			vowelCount += 1
		}

		if r == previousRune {
			letterTwice = true
		}

		previousRune = r
	}
	return vowelCount >= 3 && letterTwice
}

func part2(input string) int {
	var result int
	for _, word := range strings.Split(input, "\n") {
		if isNicePart2(word) {
			result += 1
		}
	}
	return result
}

func isNicePart2(word string) bool {
	twoLetters := map[string]int{}
	hasTwo := false
	for i := range len(word) - 1 {
		two := string(word[i : i+2])
		prev, ok := twoLetters[two]
		if ok {
			if prev < i-1 {
				hasTwo = true
				break
			}
		} else {
			twoLetters[two] = i
		}
	}

	hasRepeats := false
	for i := range len(word) - 2 {
		letter := word[i]
		if letter == word[i+2] {
			hasRepeats = true
			break
		}
	}

	return hasTwo && hasRepeats
}
