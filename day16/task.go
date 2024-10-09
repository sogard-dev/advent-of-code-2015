package day16

import (
	"strings"

	"github.com/sogard-dev/advent-of-code-2015/utils"
)

type ticker struct {
	name        string
	children    int
	cats        int
	samoyeds    int
	pomeranians int
	akitas      int
	vizslas     int
	goldfish    int
	trees       int
	cars        int
	perfumes    int
}

func parse(line string) ticker {
	sides := strings.SplitN(line, ": ", 2)
	children := -1
	cats := -1
	samoyeds := -1
	pomeranians := -1
	akitas := -1
	vizslas := -1
	goldfish := -1
	trees := -1
	cars := -1
	perfumes := -1

	for _, item := range strings.Split(sides[1], ", ") {
		spl := strings.Split(item, ": ")
		name := spl[0]
		num := utils.GetAllNumbers(spl[1])[0]

		switch name {
		case "children":
			children = num
		case "cats":
			cats = num
		case "samoyeds":
			samoyeds = num
		case "pomeranians":
			pomeranians = num
		case "akitas":
			akitas = num
		case "vizslas":
			vizslas = num
		case "goldfish":
			goldfish = num
		case "trees":
			trees = num
		case "cars":
			cars = num
		case "perfumes":
			perfumes = num
		default:
			panic("Shit: " + name)
		}
	}

	return ticker{
		name:        sides[0],
		children:    children,
		cats:        cats,
		samoyeds:    samoyeds,
		pomeranians: pomeranians,
		akitas:      akitas,
		vizslas:     vizslas,
		goldfish:    goldfish,
		trees:       trees,
		cars:        cars,
		perfumes:    perfumes,
	}
}

func isValid(goal, ticker ticker) bool {
	if ticker.children >= 0 && goal.children != ticker.children {
		return false
	}
	if ticker.cats >= 0 && goal.cats != ticker.cats {
		return false
	}
	if ticker.samoyeds >= 0 && goal.samoyeds != ticker.samoyeds {
		return false
	}
	if ticker.pomeranians >= 0 && goal.pomeranians != ticker.pomeranians {
		return false
	}
	if ticker.akitas >= 0 && goal.akitas != ticker.akitas {
		return false
	}
	if ticker.vizslas >= 0 && goal.vizslas != ticker.vizslas {
		return false
	}
	if ticker.goldfish >= 0 && goal.goldfish != ticker.goldfish {
		return false
	}
	if ticker.trees >= 0 && goal.trees != ticker.trees {
		return false
	}
	if ticker.cars >= 0 && goal.cars != ticker.cars {
		return false
	}
	if ticker.perfumes >= 0 && goal.perfumes != ticker.perfumes {
		return false
	}

	return true
}

func part1(input string) int {
	goal := ticker{
		children:    3,
		cats:        7,
		samoyeds:    2,
		pomeranians: 3,
		akitas:      0,
		vizslas:     0,
		goldfish:    5,
		trees:       3,
		cars:        2,
		perfumes:    1,
	}

	for _, line := range strings.Split(input, "\n") {
		ticker := parse(line)
		if isValid(goal, ticker) {
			return utils.GetAllNumbers(ticker.name)[0]
		}
	}

	return 0
}

func isValidPart2(goal, ticker ticker) bool {
	if ticker.cats >= 0 && goal.cats >= ticker.cats {
		return false
	}
	if ticker.trees >= 0 && goal.trees >= ticker.trees {
		return false
	}
	if ticker.pomeranians >= 0 && goal.pomeranians <= ticker.pomeranians {
		return false
	}
	if ticker.goldfish >= 0 && goal.goldfish <= ticker.goldfish {
		return false
	}

	if ticker.children >= 0 && goal.children != ticker.children {
		return false
	}
	if ticker.samoyeds >= 0 && goal.samoyeds != ticker.samoyeds {
		return false
	}
	if ticker.akitas >= 0 && goal.akitas != ticker.akitas {
		return false
	}
	if ticker.vizslas >= 0 && goal.vizslas != ticker.vizslas {
		return false
	}
	if ticker.cars >= 0 && goal.cars != ticker.cars {
		return false
	}
	if ticker.perfumes >= 0 && goal.perfumes != ticker.perfumes {
		return false
	}

	return true
}

func part2(input string) int {
	goal := ticker{
		children:    3,
		cats:        7,
		samoyeds:    2,
		pomeranians: 3,
		akitas:      0,
		vizslas:     0,
		goldfish:    5,
		trees:       3,
		cars:        2,
		perfumes:    1,
	}

	for _, line := range strings.Split(input, "\n") {
		ticker := parse(line)
		if isValidPart2(goal, ticker) {
			return utils.GetAllNumbers(ticker.name)[0]
		}
	}

	return 0
}
