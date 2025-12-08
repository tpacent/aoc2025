package day8

import (
	"aoc2025/lib"
	"bufio"
	"cmp"
	"io"
	"slices"
	"strings"
)

func CalcLargestCircuits(points []Point, connections int, limit int) (total int) {
	cb := NewCircuitBuilder()
	for _, pair := range GetClosestPoints(points)[:connections] {
		cb.AddPair(pair.PA, pair.PB)
	}

	total = 1
	for _, stat := range cb.CircuitStats()[:limit] {
		total *= stat.Size
	}

	return total
}

func CalcLastJunctions(points []Point) int {
	cb := NewCircuitBuilder()

	for _, pair := range GetClosestPoints(points) {
		pts, cts := cb.AddPair(pair.PA, pair.PB)
		if cts == 1 && pts == len(points) {
			return int(pair.PA.X) * int(pair.PB.X)
		}
	}

	return -1
}

type Point struct {
	X, Y, Z int32
}

func DistSquared(a, b Point) int {
	x := int(a.X - b.X)
	y := int(a.Y - b.Y)
	z := int(a.Z - b.Z)
	return x*x + y*y + z*z
}

func ParseInput(r io.Reader) (points []Point) {
	for scanner := bufio.NewScanner(r); scanner.Scan(); {
		if coords := strings.Split(scanner.Text(), ","); len(coords) == 3 {
			points = append(points, Point{
				X: int32(lib.MustAtoi(coords[0])),
				Y: int32(lib.MustAtoi(coords[1])),
				Z: int32(lib.MustAtoi(coords[2])),
			})
		}
	}

	return
}

type PointPair struct {
	Dist int
	PA   Point
	PB   Point
}

func GetClosestPoints(points []Point) []PointPair {
	pairs := make([]PointPair, 0, len(points)*len(points))

	for indexA := 0; indexA < len(points)-1; indexA++ {
		for indexB := indexA + 1; indexB < len(points); indexB++ {
			pairs = append(pairs, PointPair{
				PA:   points[indexA],
				PB:   points[indexB],
				Dist: DistSquared(points[indexA], points[indexB]),
			})
		}
	}

	slices.SortFunc(pairs, func(a, b PointPair) int {
		switch {
		case a.Dist < b.Dist:
			return -1
		case a.Dist > b.Dist:
			return 1
		default:
			return 0
		}
	})

	return pairs
}

func NewCircuitBuilder() *CircuitBuilder {
	return &CircuitBuilder{
		pointRegistry:   make(map[Point]int),
		circuitRegistry: make(map[int]map[Point]struct{}),
	}
}

type CircuitBuilder struct {
	circuitNum      int
	pointRegistry   map[Point]int
	circuitRegistry map[int]map[Point]struct{}
}

func (cb *CircuitBuilder) AddPair(a, b Point) (int, int) {
	circuitA, hasA := cb.pointRegistry[a]
	circuitB, hasB := cb.pointRegistry[b]
	switch {
	case hasA && hasB:
		if circuitA != circuitB {
			for point := range cb.circuitRegistry[circuitB] {
				cb.pointRegistry[point] = circuitA
				cb.circuitRegistry[circuitA][point] = struct{}{}
			}
			delete(cb.circuitRegistry, circuitB)
		}
	case hasA && !hasB:
		cb.pointRegistry[b] = circuitA
		cb.circuitRegistry[circuitA][b] = struct{}{}
	case !hasA && hasB:
		cb.pointRegistry[a] = circuitB
		cb.circuitRegistry[circuitB][a] = struct{}{}
	case !hasA && !hasB:
		circuit := cb.AllocCircuit()
		cb.pointRegistry[a] = circuit
		cb.pointRegistry[b] = circuit
		cb.circuitRegistry[circuit][a] = struct{}{}
		cb.circuitRegistry[circuit][b] = struct{}{}
	}

	return len(cb.pointRegistry), len(cb.circuitRegistry)
}

func (cb *CircuitBuilder) AllocCircuit() int {
	circuit := cb.circuitNum + 1
	cb.circuitNum = circuit
	cb.circuitRegistry[circuit] = make(map[Point]struct{})
	return circuit
}

func (cb *CircuitBuilder) CircuitStats() []CircuitStat {
	stats := make([]CircuitStat, 0, len(cb.circuitRegistry))

	for id, points := range cb.circuitRegistry {
		stats = append(stats, CircuitStat{ID: id, Size: len(points)})
	}

	slices.SortFunc(stats, func(a, b CircuitStat) int {
		return cmp.Compare(b.Size, a.Size)
	})

	return stats
}

type CircuitStat struct {
	ID   int
	Size int
}
