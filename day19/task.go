package day19

import (
	"math"
	"strconv"
	"strings"
)

type rule struct {
	in  string
	out string
}

func parse(rulesStr string) []rule {
	rules := make([]rule, 0)
	for _, line := range strings.Split(rulesStr, "\n") {
		spl := strings.Split(line, " => ")
		rules = append(rules, rule{
			in:  spl[0],
			out: spl[1],
		})
	}
	return rules
}

func part1(rulesStr, molecule string) int {
	rules := parse(rulesStr)
	return options(molecule, rules)
}

func options(molecule string, rules []rule) int {
	seenMolecules := map[string]bool{}
	for _, r := range rules {
		for _, s := range getManipulatedMolecules(molecule, r, false) {
			seenMolecules[s] = true
		}
	}
	return len(seenMolecules)
}

func getManipulatedMolecules(molecule string, r rule, part2 bool) []string {
	seenMolecules := map[string]bool{}
	toReplace := r.in
	toReplaceWith := r.out
	if part2 {
		toReplace = r.out
		toReplaceWith = r.in
	}
	for i := 0; i <= len(molecule)-len(toReplace); i++ {
		matches := true
		for k := range len(toReplace) {
			matches = matches && toReplace[k] == molecule[i+k]
		}
		if matches {
			pref := molecule[0:i]
			suff := molecule[i+len(toReplace):]
			newStr := pref + toReplaceWith + suff
			seenMolecules[newStr] = true
		}
	}
	options := []string{}
	for k := range seenMolecules {
		options = append(options, k)
	}
	return options
}

func part2(rulesStr, molecule string) int {
	rules := parse(rulesStr)

	return recurse(rules, molecule)
}

func recurse(rules []rule, molecule string) int {
	ret := math.MaxInt
	seenMolecules := map[string]bool{}

	var inner func(m string, manipulations int)
	inner = func(m string, manipulations int) {
		if m == "e" {
			ret = min(ret, manipulations)
			println("Found " + strconv.Itoa(manipulations))
			return
		}
		_, ok := seenMolecules[m]
		if m == "" || manipulations >= ret || ok {
			return
		}
		seenMolecules[m] = true

		for _, r := range rules {
			for _, s := range getManipulatedMolecules(m, r, true) {
				inner(s, manipulations+1)
			}
		}
	}

	inner(molecule, 0)

	return ret
}
