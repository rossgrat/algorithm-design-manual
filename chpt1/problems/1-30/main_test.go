package main

import (
	"math"
	"testing"
)

var testCases = []struct {
	name        string
	pointSet    []Point
	optimalDist float64
}{
	{"straight line", pointSets[0], 32},
	{"unit square", pointSets[1], 4},
	{"triangle with center", pointSets[2], 1 + 1 + math.Sqrt(2)},
	{"3 collinear", pointSets[3], 10},
	{"two clusters", pointSets[4], 12},
}

func TestNearestNeighbor(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			order, dist := NearestNeighbor(tc.pointSet)
			t.Logf("distance: %.3f (optimal: %.3f)", dist, tc.optimalDist)
			t.Logf("order: %v", order)
			if len(order) != len(tc.pointSet)+1 {
				t.Errorf("expected %d points, got %d", len(tc.pointSet)+1, len(order))
			}
		})
	}
}

func TestClosestPair(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			order, dist := ClosestPair(tc.pointSet)
			t.Logf("distance: %.3f (optimal: %.3f)", dist, tc.optimalDist)
			t.Logf("order: %v", order)
			if len(order) != len(tc.pointSet)+1 {
				t.Errorf("expected %d points, got %d", len(tc.pointSet)+1, len(order))
			}
		})
	}
}
