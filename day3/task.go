package day3

type pos struct {
	x int
	y int
}

func part1(input string) int {
	visited := map[pos]bool{}
	visited[pos{}] = true

	santaPos := pos{}
	for _, c := range input {
		switch c {
		case '>':
			santaPos = pos{santaPos.x + 1, santaPos.y}
		case '<':
			santaPos = pos{santaPos.x - 1, santaPos.y}
		case '^':
			santaPos = pos{santaPos.x, santaPos.y - 1}
		case 'v':
			santaPos = pos{santaPos.x, santaPos.y + 1}
		}
		visited[santaPos] = true
	}

	return len(visited)
}

func part2(input string) int {
	visited := map[pos]bool{}
	visited[pos{}] = true

	santaPos := pos{}
	robotSantaPos := pos{}
	for i, c := range input {
		moving := &santaPos
		if i%2 == 1 {
			moving = &robotSantaPos
		}

		switch c {
		case '>':
			moving.x += 1
		case '<':
			moving.x -= 1
		case '^':
			moving.y -= 1
		case 'v':
			moving.y += 1
		}
		visited[*moving] = true
	}

	return len(visited)
}
