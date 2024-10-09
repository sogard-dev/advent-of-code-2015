package day15

import (
	"strings"

	"github.com/sogard-dev/advent-of-code-2015/utils"
)

type ingredient struct {
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func parse(input string) []ingredient {
	ret := []ingredient{}
	for _, line := range strings.Split(input, "\n") {
		num := utils.GetAllNumbers(line)
		ret = append(ret, ingredient{
			capacity:   num[0],
			durability: num[1],
			flavor:     num[2],
			texture:    num[3],
			calories:   num[4],
		})
	}
	return ret
}

func part1(input string) uint64 {
	ingredients := parse(input)
	teaspoons := 100

	permutations := make([]int, len(ingredients))
	permutations[0] = teaspoons

	var highScore uint64
	visit(0, teaspoons, permutations, func() {
		scored := score(permutations, ingredients)
		highScore = max(highScore, scored)
	})

	return highScore
}

func part2(input string) uint64 {
	ingredients := parse(input)
	teaspoons := 100

	permutations := make([]int, len(ingredients))
	permutations[0] = teaspoons

	var highScore uint64
	visit(0, teaspoons, permutations, func() {
		if calories(permutations, ingredients) == 500 {
			scored := score(permutations, ingredients)
			highScore = max(highScore, scored)
		}
	})

	return highScore
}

func visit(moving int, teaspoons int, perm []int, f func()) {
	f()

	if moving < len(perm)-1 && perm[moving] > 0 && perm[moving+1] < teaspoons {
		perm[moving]--
		perm[moving+1]++
		visit(moving, teaspoons, perm, f)
		visit(moving+1, teaspoons, perm, f)
		perm[moving+1]--
		perm[moving]++
	}
}

func score(perm []int, ingredients []ingredient) uint64 {
	var capacity, durability, flavor, texture int
	for i, amt := range perm {
		capacity += ingredients[i].capacity * amt
		durability += ingredients[i].durability * amt
		flavor += ingredients[i].flavor * amt
		texture += ingredients[i].texture * amt
	}
	return uint64(max(0, capacity)) * uint64(max(0, durability)) * uint64(max(0, flavor)) * uint64(max(0, texture))
}

func calories(perm []int, ingredients []ingredient) uint64 {
	var calories int
	for i, amt := range perm {
		calories += ingredients[i].calories * amt
	}
	return uint64(max(0, calories))
}
