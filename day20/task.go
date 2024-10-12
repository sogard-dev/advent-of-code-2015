package day20

func part1(rawInput int) int {
	return solve(rawInput/10, func(i, j int) bool {
		return true
	})
}

func part2(rawInput int) int {
	return solve(rawInput/11, func(i, j int) bool {
		return j < i*50
	})
}

func solve(input int, shouldContinue func(int, int) bool) int {
	sieve := make([]int, input)
	for i := 1; i <= input; i++ {
		for j := i; shouldContinue(i, j) && j < input; j += i {
			sieve[j] += i
		}

		if sieve[i] >= input {
			return i
		}
	}
	return 0
}
