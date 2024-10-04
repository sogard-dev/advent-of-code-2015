package day4

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
)

func part1(input string) int {
	return solve(input, "00000")
}

func part2(input string) int {
	return solve(input, "000000")
}

func solve(input string, prefix string) int {
	var i int
	for {
		if strings.HasPrefix(getMD5Hash(input+strconv.Itoa(i)), prefix) {
			return i
		}
		i++
	}
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
