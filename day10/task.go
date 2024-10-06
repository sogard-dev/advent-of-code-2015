package day10

import "strconv"

type block struct {
	number    int
	repeating int
	nextBlock *block
}

func (b *block) truncate() {
	for {
		if b.nextBlock == nil || b.number != b.nextBlock.number {
			break
		} else {
			b.repeating += b.nextBlock.repeating
			b.nextBlock = b.nextBlock.nextBlock
		}
	}

	if b.nextBlock != nil {
		b.nextBlock.truncate()
	}
}

func (b *block) iterate() {
	if b.repeating == 0 {
		b.nextBlock.iterate()
	} else {
		newBlock := &block{
			number:    b.number,
			repeating: 1,
			nextBlock: b.nextBlock,
		}
		b.number = b.repeating
		b.nextBlock = newBlock
		b.repeating = 1

		if newBlock.nextBlock != nil {
			newBlock.nextBlock.iterate()
		}
	}
}

func (b *block) count() int {
	if b.nextBlock == nil {
		return b.repeating
	} else {
		return b.repeating + b.nextBlock.count()
	}
}

func part1(input string, iterations int) int {
	rootBlock := block{}
	lastBlock := &rootBlock
	for _, s := range input {
		num, _ := strconv.Atoi(string(s))
		newBlock := &block{
			number:    num,
			repeating: 1,
		}

		lastBlock.nextBlock = newBlock
		lastBlock = newBlock
	}

	rootBlock.truncate()
	for range iterations {
		rootBlock.iterate()
		rootBlock.truncate()
	}

	return rootBlock.count()
}
