package day9

import (
	"aoc2025/lib"
	"bufio"
	"cmp"
	"io"
	"slices"
	"strings"
)

type Point struct {
	X int
	Y int
}

func GreatestArea(points []Point) (area int) {
	for indexA := range len(points) - 1 {
		for indexB := indexA + 1; indexB < len(points); indexB++ {
			rectarea := (lib.Abs(points[indexA].X-points[indexB].X) + 1) *
				(lib.Abs(points[indexA].Y-points[indexB].Y) + 1)
			area = max(area, rectarea)
		}
	}

	return
}

func GreatestAreaInsidePerimeter(points []Point) (area int) {
	perimeter := Perimeter(points)
	rowPoints := RowPoints(perimeter)
	rowRanges := RowRanges(rowPoints)

	for indexA := range len(points) - 1 {
		for indexB := indexA + 1; indexB < len(points); indexB++ {
			rectarea := (lib.Abs(points[indexA].X-points[indexB].X) + 1) *
				(lib.Abs(points[indexA].Y-points[indexB].Y) + 1)
			if rectarea > area && IsInside(points[indexA], points[indexB], rowRanges) {
				area = rectarea
			}
		}
	}

	return
}

type Bounds struct {
	XMin int
	XMax int
	YMin int
	YMax int
}

func GetBounds(points []Point) Bounds {
	b := Bounds{
		XMin: points[0].X,
		XMax: points[0].X,
		YMin: points[0].Y,
		YMax: points[0].Y,
	}

	for _, p := range points {
		b.XMin = min(b.XMin, p.X)
		b.XMax = max(b.XMax, p.X)
		b.YMin = min(b.YMin, p.Y)
		b.YMax = max(b.YMax, p.Y)
	}

	return b
}

type Range struct {
	From int
	Upto int
}

func RowRanges(rowPoints map[int][]Point) map[int][]Range {
	rr := make(map[int][]Range, len(rowPoints))
	for row, points := range rowPoints {

		for len(points) > 0 {
			rangeStart := points[0]
			rangeEnd := rangeStart
			index := 0

			for index = range len(points) {
				rangeEnd = points[index]
				if points[index].X-rangeEnd.X > 1 {
					break
				}
			}

			for {
				if index == len(points)-1 {
					break
				}
				index++
				if points[index].X-rangeEnd.X < 2 {
					rangeEnd = points[index]
				}
			}

			points = points[index+1:]
			rr[row] = append(rr[row], Range{rangeStart.X, rangeEnd.X})
		}

	}
	return rr
}

func RowPoints(points []Point) map[int][]Point {
	rp := make(map[int][]Point, len(points))

	for _, p := range points {
		rp[p.Y] = append(rp[p.Y], p)
	}

	for row := range rp {
		slices.SortFunc(rp[row], func(a, b Point) int { return cmp.Compare(a.X, b.X) })
	}

	return rp
}

func IsInside(p1, p2 Point, ranges map[int][]Range) bool {
	xRange := Range{From: min(p1.X, p2.X), Upto: max(p1.X, p2.X)}
	yMin := min(p1.Y, p2.Y)
	yMax := max(p1.Y, p2.Y)

	for row := yMin; row <= yMax; row++ {
		if !rangeInside(xRange, ranges[row]) {
			return false
		}
	}
	return true
}

func rangeInside(r Range, rs []Range) bool {
	for _, t := range rs {
		if r.From >= t.From && r.Upto <= t.Upto {
			return true
		}
	}
	return false
}

func Perimeter(points []Point) []Point {
	perim := make([]Point, 0, len(points))
	for index, point := range points[:len(points)-1] {
		perim = append(perim, pointsLine(point, points[index+1])...)
	}
	return append(perim, pointsLine(points[len(points)-1], points[0])...)
}

func pointsLine(a, b Point) (out []Point) {
	dx := cmp.Compare(b.X, a.X)
	dy := cmp.Compare(b.Y, a.Y)
	for {
		out = append(out, a)
		a.X += dx
		a.Y += dy
		if a == b {
			break
		}
	}
	return out
}

func ParseInput(r io.Reader) (points []Point) {
	for scanner := bufio.NewScanner(r); scanner.Scan(); {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		x, y, ok := strings.Cut(line, ",")
		if !ok {
			continue
		}

		points = append(points, Point{X: lib.MustAtoi(x), Y: lib.MustAtoi(y)})
	}

	return
}
