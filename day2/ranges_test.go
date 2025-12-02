package day2_test

import (
	"aoc2025/day2"
	"os"
	"strings"
	"testing"
)

func TestSolveDay2Part1(t *testing.T) {
	line := getLine(t, "input.txt")
	actual := day2.SumInvalidIDHalves(day2.ParseRanges(line))
	t.Log(actual)
}

func TestSolveDay2Part2(t *testing.T) {
	line := getLine(t, "input.txt")
	actual := day2.SumInvalidIDs(day2.ParseRanges(line))
	t.Log(actual)
}

func getLine(t *testing.T, path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	line, _, _ := strings.Cut(string(data), "\n")
	return line
}

const testRanges = "11-22,95-115,998-1012,1188511880-1188511890," +
	"222220-222224,1698522-1698528,446443-446449,38593856-38593862," +
	"565653-565659,824824821-824824827,2121212118-2121212124"

func TestExample(t *testing.T) {
	actual := day2.SumInvalidIDHalves(day2.ParseRanges(testRanges))

	if actual != 1227775554 {
		t.Error("unexpected value")
	}
}

func TestExample2(t *testing.T) {
	actual := day2.SumInvalidIDs(day2.ParseRanges(testRanges))

	if actual != 4174379265 {
		t.Logf("unexpected value: %d", actual)
	}
}
