package day22

import (
	"math"
)

type gameState struct {
	playerHp    int
	playerMana  int
	bossHp      int
	bossDamange int
	poison      int
	shield      int
	recharge    int
	playerTurn  bool
}

type spell struct {
	cost     int
	damage   int
	healing  int
	poison   int
	shield   int
	recharge int
}

func (ds spell) canCast(state *gameState) bool {
	if ds.cost >= state.playerMana {
		return false
	}

	if ds.poison > 0 {
		return state.poison == 0
	}
	if ds.recharge > 0 {
		return state.recharge == 0
	}
	if ds.shield > 0 {
		return state.shield == 0
	}

	return true
}

func (ds spell) action(state gameState) gameState {
	state.playerMana -= ds.cost
	state.playerHp += ds.healing
	state.bossHp -= ds.damage
	state.poison = max(state.poison, ds.poison)
	state.recharge = max(state.recharge, ds.recharge)
	state.shield = max(state.shield, ds.shield)
	return state
}

func (ds spell) getCost() int {
	return ds.cost
}

func part1(playerHitpoints, playerMana, bossHitpoints, bossDamage int) int {
	return solve(gameState{
		playerHp:    playerHitpoints,
		playerMana:  playerMana,
		bossHp:      bossHitpoints,
		bossDamange: bossDamage,
		playerTurn:  true,
	}, false)
}

func part2(playerHitpoints, playerMana, bossHitpoints, bossDamage int) int {
	return solve(gameState{
		playerHp:    playerHitpoints,
		playerMana:  playerMana,
		bossHp:      bossHitpoints,
		bossDamange: bossDamage,
		playerTurn:  true,
	}, true)
}

func solve(initGameSate gameState, hardmode bool) int {
	spells := []spell{
		{
			cost:   173,
			poison: 6,
		},
		{
			cost:     229,
			recharge: 5,
		},
		{
			cost:   53,
			damage: 4,
		},
		{
			cost:   113,
			shield: 6,
		},
		{
			cost:    73,
			damage:  2,
			healing: 2,
		},
	}

	lowest := math.MaxInt
	playerWon := func(manacost int) {
		lowest = min(lowest, manacost)
	}

	var simulate func(gs gameState, manaSpent int)
	simulate = func(gs gameState, manaSpent int) {
		if manaSpent > lowest {
			return
		}

		isPlayerTurn := gs.playerTurn
		gs.playerTurn = !gs.playerTurn

		//handle effects
		if isPlayerTurn && hardmode {
			gs.playerHp -= 1
		}
		if gs.recharge > 0 {
			gs.playerMana += 101
			gs.recharge -= 1
		}
		if gs.shield > 0 {
			gs.shield -= 1
		}
		if gs.poison > 0 {
			gs.poison = gs.poison - 1
			gs.bossHp -= 3
		}

		if gs.playerHp <= 0 {
			return
		}
		if gs.bossHp <= 0 {
			playerWon(manaSpent)
			return
		}

		//handle turn
		if isPlayerTurn {
			for _, spell := range spells {
				if spell.canCast(&gs) {
					simulate(spell.action(gs), manaSpent+spell.getCost())
				}
			}
		} else {
			armor := 0
			if gs.shield > 0 {
				armor += 7
			}
			gs.playerHp -= max(1, gs.bossDamange-armor)
			simulate(gs, manaSpent)
		}
	}

	simulate(initGameSate, 0)

	return lowest
}
