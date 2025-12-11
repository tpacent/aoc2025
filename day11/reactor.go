package day11

import (
	"bufio"
	"io"
	"strings"
)

func WalkThrough(src string, dst string, data map[string][]string, midA, midB string) (total int) {
	total = 1
	if count := Walk(midA, midB, data); count > 0 {
		total *= count
		total *= Walk(src, midA, data)
		total *= Walk(midB, dst, data)
	} else {
		total *= Walk(src, midB, data)
		total *= Walk(midB, midA, data)
		total *= Walk(midA, dst, data)
	}
	return
}

func Walk(src string, dst string, data map[string][]string) (total int) {
	return walk(src, dst, data, &map[string]int{})
}

func walk(src string, dst string, data map[string][]string, cache *map[string]int) (total int) {
	if src == dst {
		return 1
	}

	if value, ok := (*cache)[src]; ok {
		return value
	}

	for _, next := range data[src] {
		total += walk(next, dst, data, cache)
	}

	(*cache)[src] = total
	return
}

func ParseInput(r io.Reader) map[string][]string {
	links := map[string][]string{}

	for scanner := bufio.NewScanner(r); scanner.Scan(); {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		src, dests, _ := strings.Cut(line, ":")
		for dst := range strings.FieldsSeq(dests) {
			links[src] = append(links[src], dst)
		}
	}

	return links
}
