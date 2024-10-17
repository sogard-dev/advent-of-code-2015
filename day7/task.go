package day7

import (
	"fmt"
	"strconv"
	"strings"
)

type node interface {
	value(map[string]node, map[string]uint16) uint16
}

type signal struct{ val string }
type not struct{ n string }
type and struct{ left, right string }
type or struct{ left, right string }
type rshift struct {
	n  string
	by uint16
}
type lshift struct {
	n  string
	by uint16
}

func valueOfNode(nodes map[string]node, memory map[string]uint16, n string) uint16 {
	mem, ok := memory[n]
	if ok {
		return mem
	}

	node, ok := nodes[n]
	if !ok {
		num, ok := getNum(n)
		if !ok {
			panic("Could not find node and is not number" + n)
		}
		memory[n] = num
		return num
	}
	ret := node.value(nodes, memory)
	memory[n] = ret
	return ret
}

func (n signal) value(nodes map[string]node, memory map[string]uint16) uint16 {
	return valueOfNode(nodes, memory, n.val)
}

func (n and) value(nodes map[string]node, memory map[string]uint16) uint16 {
	return valueOfNode(nodes, memory, n.left) & valueOfNode(nodes, memory, n.right)
}

func (n or) value(nodes map[string]node, memory map[string]uint16) uint16 {
	return valueOfNode(nodes, memory, n.left) | valueOfNode(nodes, memory, n.right)
}

func (n not) value(nodes map[string]node, memory map[string]uint16) uint16 {
	return ^valueOfNode(nodes, memory, n.n)
}

func (n rshift) value(nodes map[string]node, memory map[string]uint16) uint16 {
	return valueOfNode(nodes, memory, n.n) >> n.by
}

func (n lshift) value(nodes map[string]node, memory map[string]uint16) uint16 {
	return valueOfNode(nodes, memory, n.n) << n.by
}

func getNum(str string) (uint16, bool) {
	v, err := strconv.ParseInt(str, 10, 16)
	if err != nil {
		return 0, false
	}
	return uint16(v), true
}

func part1(input string, endWire string) int {
	nodes := map[string]node{}
	memory := map[string]uint16{}

	for _, line := range strings.Split(input, "\n") {
		node, outputWire := createNode(line)
		nodes[outputWire] = node
	}

	return int(nodes[endWire].value(nodes, memory))
}

func createNode(line string) (node, string) {
	sides := strings.Split(line, " -> ")
	lside := sides[0]
	outputWire := sides[1]

	lspl := strings.Split(lside, " ")

	if len(lspl) == 1 {
		return signal{lspl[0]}, outputWire
	} else if len(lspl) == 2 {
		operator := lspl[0]
		wire := lspl[1]
		if operator == "NOT" {
			return not{
				n: wire,
			}, outputWire
		}
	} else {
		operator := lspl[1]
		arg1 := lspl[0]
		arg2 := lspl[2]

		switch operator {
		case "OR":
			return or{
				left:  arg1,
				right: arg2,
			}, outputWire
		case "AND":
			return and{
				left:  arg1,
				right: arg2,
			}, outputWire
		case "LSHIFT":
			num, _ := getNum(arg2)
			return lshift{
				n:  arg1,
				by: num,
			}, outputWire
		case "RSHIFT":
			num, _ := getNum(arg2)
			return rshift{
				n:  arg1,
				by: num,
			}, outputWire
		}
	}
	panic("Unknown operator " + line)
}

func part2(input string) int {
	wireA := part1(input, "a")
	newInput := strings.ReplaceAll(input, "14146 -> b", fmt.Sprintf("%d -> b", wireA))
	newWireA := part1(newInput, "a")
	return int(newWireA)
}
