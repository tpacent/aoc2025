package day2

import (
	"aoc2025/lib"
	"fmt"
	"maps"
	"slices"
	"strconv"
	"strings"
)

func SumInvalidIDs(ranges []Range) (total int) {
	for _, r := range ranges {
		for _, id := range FindInvalidIDs(r.Lo, r.Hi) {
			total += id
		}
	}

	return
}

func FindInvalidIDs(lower, upper int) []int {
	upperstr := strconv.Itoa(upper)
	ids := map[int]struct{}{}

	for chunklen := 1; chunklen <= len(upperstr)/2; chunklen++ {
		for _, id := range FindInvalidIDLen(lower, upper, chunklen) {
			ids[id] = struct{}{}
		}
	}

	return slices.Collect(maps.Keys(ids))
}

func FindInvalidIDLen(lower, upper, chunklen int) (ids []int) {
	lower = max(10, lower)

	for lower <= upper {
		fromstr := strconv.Itoa(lower)

		if len(fromstr)%chunklen != 0 {
			lower = lib.Pow(10, len(fromstr))
			continue
		}

		chunk := fromstr[:chunklen]
		upperBound := lib.Pow(10, len(chunk))

		for n := lib.MustAtoi(chunk); n < upperBound; n++ {
			nstr := strconv.Itoa(n)
			idstr := strings.Repeat(nstr, len(fromstr)/chunklen)
			id := lib.MustAtoi(idstr)

			if id > upper {
				break
			}

			if id >= lower {
				ids = append(ids, id)
			}
		}

		lower = lib.Pow(10, len(fromstr))
	}

	return ids
}

func SumInvalidIDHalves(ranges []Range) (total int) {
	for _, r := range ranges {
		for _, id := range FindInvalidHalves(r.Lo, r.Hi) {
			total += id
		}
	}
	return
}

func FindInvalidHalves(lower, upper int) (ids []int) {
	for lower <= upper {
		if lowerstr := strconv.Itoa(lower); len(lowerstr)%2 == 0 {
			upto := min(upper, lib.Pow(10, len(lowerstr))-1)
			subrange := FindInvalidIDLen(lower, upto, len(lowerstr)/2)
			ids = append(ids, subrange...)
		}

		lower = lib.Pow(10, len(strconv.Itoa(lower)))
	}

	return
}

type Range struct {
	Lo int
	Hi int
}

func ParseRanges(line string) (ranges []Range) {
	var from, upto int

	for s := range strings.SplitSeq(line, ",") {
		_, err := fmt.Sscanf(s, "%d-%d", &from, &upto)
		if err != nil {
			panic(err)
		}
		ranges = append(ranges, Range{Lo: from, Hi: upto})
	}

	return
}
