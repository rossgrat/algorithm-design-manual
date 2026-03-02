package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

func main() {
	pointSet := []Point{
		{
			X: 0,
			Y: 0,
		},
		{
			X: 1,
			Y: 0,
		},
		{
			X: 3,
			Y: 0,
		},
		{
			X: -5,
			Y: 0,
		},
		{
			X: 11,
			Y: 0,
		},
		{
			X: -21,
			Y: 0,
		},
		{
			X: -1,
			Y: 0,
		},
	}

	order := nearestNeighbor(pointSet)

	for _, p := range order {
		fmt.Println(p.X, p.Y)
	}
}

func calcDistance(p1 Point, p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.X-p1.X, 2) + math.Pow(p2.Y-p1.Y, 2))
}

func nearestNeighbor(pointSet []Point) []Point {
	visitedPoints := map[int]struct{}{}
	bestOrder := []Point{}

	pPrev := pointSet[0]
	bestOrder = append(bestOrder, pPrev)
	visitedPoints[0] = struct{}{}

	for len(visitedPoints) != len(pointSet) {

		shortestDistance := math.Inf(1)
		bestPIdx := 0
		for i, p := range pointSet {
			if _, ok := visitedPoints[i]; ok {
				continue
			}
			dist := calcDistance(pPrev, p)
			if dist < shortestDistance {
				shortestDistance = dist
				bestPIdx = i
			}
		}

		visitedPoints[bestPIdx] = struct{}{}
		bestOrder = append(bestOrder, pointSet[bestPIdx])

		pPrev = pointSet[bestPIdx]
	}

	return bestOrder
}
