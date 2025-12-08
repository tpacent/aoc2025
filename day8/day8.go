package day8

import (
	"aoc2025/lib"
	"bufio"
	"container/heap"
	"io"
	"strings"
)

func CalcLargestCircuits(points []Point, connections int, limit int) (total int) {
	cb := NewCircuitBuilder()

	pairs := GetClosestPoints(points)
	for range connections {
		pair := heap.Pop(&pairs).(PointPairDist)
		cb.AddPair(pair.PointPair)
	}

	total = 1
	stats := cb.CircuitStats()
	for range limit {
		total *= heap.Pop(&stats).(int)
	}
	return total
}

func CalcLastJunctions(points []Point) int {
	cb := NewCircuitBuilder()
	pairs := GetClosestPoints(points)
	for {
		pair := heap.Pop(&pairs).(PointPairDist)
		pts, cts := cb.AddPair(pair.PointPair)
		if cts == 1 && pts == len(points) {
			return int(points[pair.IdxA].X) * int(points[pair.IdxB].X)
		}
	}
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
	IdxA int
	IdxB int
}

type PointPairDist struct {
	PointPair
	Dist int
}

func GetClosestPoints(points []Point) PointHeap {
	pairs := make(PointHeap, 0, lib.Pow(len(points), 2)/2)
	heap.Init(&pairs)

	for indexA := 0; indexA < len(points)-1; indexA++ {
		for indexB := indexA + 1; indexB < len(points); indexB++ {
			heap.Push(&pairs, PointPairDist{
				PointPair: PointPair{
					IdxA: indexA,
					IdxB: indexB,
				},
				Dist: DistSquared(points[indexA], points[indexB]),
			})
		}
	}

	return pairs
}

func NewCircuitBuilder() *CircuitBuilder {
	return &CircuitBuilder{
		pointRegistry:   make(map[int]int),
		circuitRegistry: make(map[int]map[int]struct{}),
	}
}

type CircuitBuilder struct {
	circuitNum      int
	pointRegistry   map[int]int
	circuitRegistry map[int]map[int]struct{}
}

func (cb *CircuitBuilder) AddPair(pair PointPair) (int, int) {
	circuitA, hasA := cb.pointRegistry[pair.IdxA]
	circuitB, hasB := cb.pointRegistry[pair.IdxB]
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
		cb.pointRegistry[pair.IdxB] = circuitA
		cb.circuitRegistry[circuitA][pair.IdxB] = struct{}{}
	case !hasA && hasB:
		cb.pointRegistry[pair.IdxA] = circuitB
		cb.circuitRegistry[circuitB][pair.IdxA] = struct{}{}
	case !hasA && !hasB:
		circuit := cb.AllocCircuit()
		cb.pointRegistry[pair.IdxA] = circuit
		cb.pointRegistry[pair.IdxB] = circuit
		cb.circuitRegistry[circuit][pair.IdxA] = struct{}{}
		cb.circuitRegistry[circuit][pair.IdxB] = struct{}{}
	}

	return len(cb.pointRegistry), len(cb.circuitRegistry)
}

func (cb *CircuitBuilder) AllocCircuit() int {
	circuit := cb.circuitNum + 1
	cb.circuitNum = circuit
	cb.circuitRegistry[circuit] = make(map[int]struct{})
	return circuit
}

func (cb *CircuitBuilder) CircuitStats() SizeHeap {
	stats := make(SizeHeap, 0, len(cb.pointRegistry))
	heap.Init(&stats)

	for _, points := range cb.circuitRegistry {
		heap.Push(&stats, len(points))
	}

	return stats
}

type CircuitStat struct {
	ID   int
	Size int
}

type PointHeap []PointPairDist

func (ph PointHeap) Len() int           { return len(ph) }
func (ph PointHeap) Less(i, j int) bool { return ph[i].Dist < ph[j].Dist }
func (ph PointHeap) Swap(i, j int)      { ph[i], ph[j] = ph[j], ph[i] }
func (ph *PointHeap) Push(x any) {
	*ph = append(*ph, x.(PointPairDist))
}
func (ph *PointHeap) Pop() any {
	old := *ph
	n := len(old)
	x := old[n-1]
	*ph = old[0 : n-1]
	return x
}

type SizeHeap []int

func (ih SizeHeap) Len() int           { return len(ih) }
func (ih SizeHeap) Less(i, j int) bool { return ih[i] > ih[j] }
func (ih SizeHeap) Swap(i, j int)      { ih[i], ih[j] = ih[j], ih[i] }
func (ih *SizeHeap) Push(x any) {
	*ih = append(*ih, x.(int))
}
func (ih *SizeHeap) Pop() any {
	old := *ih
	n := len(old)
	x := old[n-1]
	*ih = old[0 : n-1]
	return x
}
