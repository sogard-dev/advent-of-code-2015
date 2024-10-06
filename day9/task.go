package day9

import (
	"strconv"
	"strings"
)

func part1(input string) int {
	graph := createGraph(input)
	ret := 9999999
	for _, n := range graph.nodes {
		ret = min(dfs(graph, n, func(a, b int) int {
			return min(a, b)
		}), ret)
	}
	return ret
}

func part2(input string) int {
	graph := createGraph(input)
	ret := 0
	for _, n := range graph.nodes {
		ret = max(dfs(graph, n, func(a, b int) int {
			return max(a, b)
		}), ret)
	}
	return ret
}

func dfs(g graph, n string, compareFunc func(int, int) int) int {
	neighbours := map[string][]string{}
	for _, n := range g.nodes {
		neighbours[n] = g.neighbours(n)
	}

	visited := map[string]bool{}
	var recurse func(visited map[string]bool, visiting string) int
	recurse = func(visited map[string]bool, visiting string) int {
		visited[visiting] = true
		weight := 0

		for _, n := range neighbours[visiting] {
			_, ok := visited[n]
			if !ok {
				weightThisEdge := g.getEdge(visiting, n).weight + recurse(visited, n)
				if weight == 0 {
					weight = weightThisEdge
				} else {
					weight = compareFunc(weight, weightThisEdge)
				}
			}
		}

		delete(visited, visiting)
		return weight
	}

	return recurse(visited, n)
}

func createGraph(input string) graph {
	edges := []edge{}
	for _, s := range strings.Split(input, "\n") {
		spl := strings.Split(s, " ")
		num, _ := strconv.Atoi(spl[4])
		edges = append(edges,
			edge{
				from:   spl[0],
				to:     spl[2],
				weight: num,
			},
			edge{
				from:   spl[2],
				to:     spl[0],
				weight: num,
			})
	}
	return createDirectedGraphFromEdges(edges)
}

func createDirectedGraphFromEdges(edges []edge) graph {
	nodesMap := map[string]bool{}
	for _, e := range edges {
		nodesMap[e.from] = true
		nodesMap[e.to] = true
	}
	nodes := []string{}
	for k := range nodesMap {
		nodes = append(nodes, k)
	}

	return graph{edges: edges, nodes: nodes}
}

type graph struct {
	edges []edge
	nodes []string
}

func (g graph) getEdge(from, to string) edge {
	for _, e := range g.edges {
		if e.from == from && e.to == to {
			return e
		}
	}
	panic("Did not find edge")
}

func (g graph) neighbours(n string) []string {
	connected := []string{}

	for _, e := range g.edges {
		if e.from == n {
			connected = append(connected, e.to)
		}
	}

	return connected
}

type edge struct {
	weight   int
	from, to string
}
