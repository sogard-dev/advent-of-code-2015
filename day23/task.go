package day23

import (
	"strings"

	"github.com/sogard-dev/advent-of-code-2015/utils"
)

type machine struct {
	ins       int
	registers []int
}

type instruction interface {
	execute(*machine)
}

type hlf struct {
	register int
}

type tpl struct {
	register int
}

type inc struct {
	register int
}

type jmp struct {
	offset int
}

type jie struct {
	register int
	offset   int
}

type jio struct {
	register int
	offset   int
}

func (s hlf) execute(m *machine) {
	m.registers[s.register] /= 2
	m.ins++
}

func (s tpl) execute(m *machine) {
	m.registers[s.register] *= 3
	m.ins++
}

func (s inc) execute(m *machine) {
	m.registers[s.register] += 1
	m.ins++
}

func (s jmp) execute(m *machine) {
	m.ins += s.offset
}

func (s jie) execute(m *machine) {
	if m.registers[s.register]%2 == 0 {
		m.ins += s.offset
	} else {
		m.ins++
	}
}

func (s jio) execute(m *machine) {
	if m.registers[s.register] == 1 {
		m.ins += s.offset
	} else {
		m.ins++
	}
}

func registerToInt(inp byte) int {
	return int(inp) - 'a'
}

func solve(input string, registers []int) {
	m := machine{
		registers: registers,
	}

	instructions := []instruction{}
	for _, line := range strings.Split(input, "\n") {
		args := strings.Split(strings.ReplaceAll(line, ",", ""), " ")
		switch args[0] {
		case "inc":
			instructions = append(instructions, inc{register: registerToInt(args[1][0])})
		case "hlf":
			instructions = append(instructions, hlf{register: registerToInt(args[1][0])})
		case "tpl":
			instructions = append(instructions, tpl{register: registerToInt(args[1][0])})
		case "jmp":
			instructions = append(instructions, jmp{offset: utils.GetAllNumbers(line)[0]})
		case "jie":
			instructions = append(instructions, jie{
				offset:   utils.GetAllNumbers(line)[0],
				register: registerToInt(args[1][0]),
			})
		case "jio":
			instructions = append(instructions, jio{
				offset:   utils.GetAllNumbers(line)[0],
				register: registerToInt(args[1][0]),
			})
		default:
			panic("Could not handle " + args[0])
		}
	}

	for m.ins < len(instructions) {
		instructions[m.ins].execute(&m)
	}
}

func part1(input string) int {
	r := []int{0, 0}
	solve(input, r)
	return r[1]
}

func part2(input string) int {
	r := []int{1, 0}
	solve(input, r)
	return r[1]
}
