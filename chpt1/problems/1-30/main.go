package main

import (
	"fmt"
	"math"
	"sort"
)

type Point struct {
	X float64
	Y float64
}

func main() {

	for _, pointSet := range pointSets {
		order, distance := ClosestPair(pointSet)
		fmt.Println("Distance", distance)
		fmt.Println("Order:")
		for _, p := range order {
			fmt.Println(p.X, p.Y)
		}
	}
}

func calcDistance(p1 Point, p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.X-p1.X, 2) + math.Pow(p2.Y-p1.Y, 2))
}

func NearestNeighbor(pointSet []Point) ([]Point, float64) {
	visitPointIndex := map[int]struct{}{}
	bestOrder := []Point{}
	totalDistance := 0.0

	pPrev := pointSet[0]
	bestOrder = append(bestOrder, pPrev)
	visitPointIndex[0] = struct{}{}

	for len(visitPointIndex) != len(pointSet) {

		shortestDistance := math.Inf(1)
		bestPointIndex := 0
		for i, p := range pointSet {
			if _, ok := visitPointIndex[i]; ok {
				continue
			}
			dist := calcDistance(pPrev, p)
			if dist < shortestDistance {
				shortestDistance = dist
				bestPointIndex = i
			}
		}
		totalDistance = totalDistance + shortestDistance
		pPrev = pointSet[bestPointIndex]
		bestOrder = append(bestOrder, pointSet[bestPointIndex])
		visitPointIndex[bestPointIndex] = struct{}{}
	}

	totalDistance += calcDistance(pPrev, bestOrder[0])
	bestOrder = append(bestOrder, bestOrder[0])
	return bestOrder, totalDistance
}

type PointNode struct {
	Point Point
	Edges [2]*PointNode
}

func (n *PointNode) EndOfChain() bool {
	return n.Edges[0] == nil || n.Edges[1] == nil
}

func (n *PointNode) OtherNeighbor(from *PointNode) *PointNode {
	if n.Edges[0] == from {
		return n.Edges[1]
	}
	return n.Edges[0]
}

func (n *PointNode) IsConnectedTo(target *PointNode) bool {
	prev := (*PointNode)(nil)
	current := n

	for current != nil {
		if current == target {
			return true
		}

		next := current.OtherNeighbor(prev)
		prev = current
		current = next
	}

	return false
}

func (n *PointNode) ConnectTo(other *PointNode) {
	if n.Edges[0] == nil {
		n.Edges[0] = other
	} else {
		n.Edges[1] = other
	}

	if other.Edges[0] == nil {
		other.Edges[0] = n
	} else {
		other.Edges[1] = n
	}
}

func (n *PointNode) FindEnd() *PointNode {
	prev := (*PointNode)(nil)
	current := n

	for !current.EndOfChain() {
		next := current.OtherNeighbor(prev)
		prev = current
		current = next
	}
	return current
}

func (n *PointNode) PrintPoints() []Point {
	if !n.EndOfChain() {
		return []Point{}
	}

	points := []Point{}

	prev := (*PointNode)(nil)
	current := n

	for current != nil {
		points = append(points, current.Point)
		next := current.OtherNeighbor(prev)
		prev = current
		current = next
	}
	return points
}

type Edge struct {
	I, J     int
	Distance float64
}

func ClosestPair(pointSet []Point) ([]Point, float64) {
	totalDistance := 0.0

	pointNodes := make([]PointNode, len(pointSet))
	for i, p := range pointSet {
		pointNodes[i] = PointNode{
			Point: p,
		}
	}

	edges := []Edge{}
	for i := 0; i < len(pointNodes); i++ {
		for j := i + 1; j < len(pointNodes); j++ {
			edges = append(edges, Edge{
				I:        i,
				J:        j,
				Distance: calcDistance(pointNodes[i].Point, pointNodes[j].Point),
			})
		}
	}

	sort.Slice(edges, func(a, b int) bool {
		return edges[a].Distance < edges[b].Distance
	})

	for _, e := range edges {
		p1 := &pointNodes[e.I]
		p2 := &pointNodes[e.J]

		if !p1.EndOfChain() || !p2.EndOfChain() {
			continue
		}
		if p1.IsConnectedTo(p2) {
			continue
		}

		p1.ConnectTo(p2)
		totalDistance += e.Distance
	}

	end := pointNodes[0].FindEnd()
	points := end.PrintPoints()
	totalDistance += calcDistance(points[0], points[len(points)-1])
	points = append(points, points[0])
	return points, totalDistance
}
