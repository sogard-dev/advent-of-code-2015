package day24

import (
	"math"
	"slices"

	"github.com/sogard-dev/advent-of-code-2015/utils"
)

func part1(input string) uint64 {
	return solve(input, 3)
}

func part2(input string) uint64 {
	return solve(input, 4)
}

func solve(input string, satchels int) uint64 {
	nums := utils.GetAllNumbers(input)
	f := func(a int, b int) int {
		return b - a
	}
	slices.SortFunc(nums, f)
	idealWeight := sum(nums, -1) / satchels

	var smallestQe uint64 = math.MaxUint64
	smallestLe := math.MaxInt
	cb := func(qe uint64, used int) {
		size := length(used)
		if size < smallestLe {
			smallestLe = size
			smallestQe = qe
			println("Smallest qe: ", qe)
		} else if size == smallestLe && qe < smallestQe {
			smallestQe = qe
			println("Smallest qe: ", qe)
		}
	}

	var solveFront func(toPack []int, ideal int, used int)
	solveFront = func(toPack []int, ideal int, used int) {
		le := length(used)
		qe := qe(toPack, used)
		if le > smallestLe {
			return
		} else if le == smallestLe && qe > smallestQe {
			return
		}

		c := sum(toPack, used)
		if c == ideal {
			cb(qe, used)
		}

		for i, n := range toPack {
			if c+n <= ideal && !utils.HasBit(used, i) {
				solveFront(toPack, ideal, utils.SetBit(used, i))
			}
		}
	}

	solveFront(nums, idealWeight, 0)
	return smallestQe
}

func length(used int) int {
	acc := 0
	for i := range 32 {
		if utils.HasBit(used, i) {
			acc += 1
		}
	}
	return acc
}

func sum(nums []int, used int) int {
	acc := 0
	for i, n := range nums {
		if utils.HasBit(used, i) {
			acc += n
		}
	}
	return acc
}

func qe(nums []int, used int) uint64 {
	acc := uint64(1)
	for i, n := range nums {
		if utils.HasBit(used, i) {
			acc *= uint64(n)
		}
	}
	return acc
}
