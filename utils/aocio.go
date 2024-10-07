package utils

import (
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"testing"
)

func GetInput(t *testing.T, file string) string {
	goldenfile := filepath.Join("testdata", file)
	want, err := os.ReadFile(goldenfile)
	if err != nil {
		t.Fatal("error reading golden file:", err)
	}

	return string(want)
}

func GetAllNumbers(input string) []int {
	re := regexp.MustCompile("[-0-9]+")
	numbers := re.FindAllString(input, -1)
	ret := []int{}
	for _, v := range numbers {
		v, err := strconv.Atoi(v)
		if err != nil {
			panic("shit")
		}
		ret = append(ret, v)
	}
	return ret
}
