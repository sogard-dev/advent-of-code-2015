package day21

import (
	"math"
	"strings"

	"github.com/sogard-dev/advent-of-code-2015/utils"
)

type unit struct {
	armor     int
	hitpoints int
	damage    int
}

type item struct {
	armor  int
	damage int
	cost   int
}

func parse(input string, optional bool) []item {
	ret := []item{}
	if optional {
		ret = append(ret, item{})
	}
	for _, line := range strings.Split(input, "\n") {
		nums := utils.GetAllNumbers(line)
		ret = append(ret, item{
			cost:   nums[0],
			damage: nums[1],
			armor:  nums[2],
		})
	}
	return ret
}

func part1(player unit, boss unit) int {
	lowestCost := math.MaxInt

	solve(player, boss, func(cost int, won bool) {
		if won && cost < lowestCost {
			lowestCost = cost
		}
	})

	return lowestCost
}

func part2(player unit, boss unit) int {
	highestCost := 0

	solve(player, boss, func(cost int, won bool) {
		if !won && highestCost < cost {
			highestCost = cost
		}
	})

	return highestCost
}

func solve(player unit, boss unit, callback func(int, bool)) {
	weapons := parse(`Dagger        8     4       0
Shortsword   10     5       0
Warhammer    25     6       0
Longsword    40     7       0
Greataxe     74     8       0`, false)

	chests := parse(`Leather      13     0       1
Chainmail    31     0       2
Splintmail   53     0       3
Bandedmail   75     0       4
Platemail   102     0       5`, true)

	rings := parse(`Damage     25     1       0
Damage     50     2       0
Damage    100     3       0
Defense   20     0       1
Defense   40     0       2
Defense   80     0       3`, true)

	for _, weapon := range weapons {
		for _, chest := range chests {
			for i, ring1 := range rings {
				for _, ring2 := range rings {
					if ring1 != ring2 || i == 0 {
						cost := weapon.cost + chest.cost + ring1.cost + ring2.cost
						armor := weapon.armor + chest.armor + ring1.armor + ring2.armor
						damage := weapon.damage + chest.damage + ring1.damage + ring2.damage

						newPlayer := unit{
							hitpoints: player.hitpoints,
							armor:     player.armor + armor,
							damage:    player.damage + damage,
						}

						result := fight(newPlayer, boss)

						callback(cost, result)
					}
				}
			}
		}
	}
}

func fight(player unit, boss unit) bool {
	for {
		boss.hitpoints -= max(1, player.damage-boss.armor)
		if boss.hitpoints <= 0 {
			return true
		}

		player.hitpoints -= max(1, boss.damage-player.armor)
		if player.hitpoints <= 0 {
			return false
		}
	}
}
