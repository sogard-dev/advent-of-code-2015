package day12

import (
	"encoding/json"

	"github.com/sogard-dev/advent-of-code-2015/utils"
)

type Foo struct {
	X map[string]interface{} `json:"-"`
}

func part1(input string) int {
	sum := 0
	for _, v := range utils.GetAllNumbers(input) {
		sum += v
	}
	return sum
}

func part2(input string) int {
	f := Foo{}
	if err := json.Unmarshal([]byte(input), &f.X); err != nil {
		panic(err)
	}

	sum := 0
	var visitArray func(node []interface{})
	var visitMap func(node map[string]interface{})

	visitArray = func(node []interface{}) {
		for _, v := range node {
			switch result := v.(type) {
			case map[string]interface{}:
				visitMap(result)
			case []interface{}:
				visitArray(result)
			case float64:
				sum += int(result)
			}
		}
	}

	visitMap = func(node map[string]interface{}) {
		for _, v := range node {
			switch result := v.(type) {
			case map[string]interface{}:
			case []interface{}:
			default:
				if result == "red" {
					return
				}
			}
		}

		for _, v := range node {
			switch result := v.(type) {
			case map[string]interface{}:
				visitMap(result)
			case []interface{}:
				visitArray(result)
			case float64:
				sum += int(result)
			}
		}
	}

	visitMap(f.X)

	return sum
}
