package day5

import (
	"aoc2025/lib"
	"bufio"
	"cmp"
	"fmt"
	"io"
	"slices"
)

func TotalFresh(ranges [][2]int) (count int) {
	for _, r := range ranges {
		count += r[1] - r[0] + 1
	}
	return
}

func CountFresh(ingrs []int, ranges [][2]int) (count int) {
	currRange := ranges[0]
	for _, ingr := range ingrs {
		for currRange[1] < ingr && len(ranges) > 1 {
			ranges = ranges[1:]
			currRange = ranges[0]
		}

		if ingr >= currRange[0] && ingr <= currRange[1] {
			count++
		}
	}
	return
}

func ParseInput(r io.Reader) ([]int, [][2]int) {
	ranges := [][2]int{}

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			if len(ranges) > 0 {
				break
			} else {
				continue
			}
		}

		var from, upto int
		if _, err := fmt.Sscanf(line, "%d-%d", &from, &upto); err != nil {
			panic(err)
		}

		ranges = append(ranges, [2]int{from, upto})
	}

	ingredients := []int{}
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		ingredients = append(ingredients, lib.MustAtoi(line))
	}

	slices.Sort(ingredients)
	slices.SortFunc(ranges, func(a, b [2]int) int {
		return cmp.Compare(a[0], b[0])
	})

	return ingredients, mergeRanges(ranges)
}

func mergeRanges(ranges [][2]int) (merged [][2]int) {
	acc := ranges[0]

	for {
		r := ranges[0]

		if r[0] > acc[1]+1 {
			merged = append(merged, acc)
			acc = r
		}

		acc[1] = max(acc[1], r[1])

		if ranges = ranges[1:]; len(ranges) == 0 {
			merged = append(merged, acc)
			break
		}
	}

	return
}
