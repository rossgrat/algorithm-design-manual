package main

var pointSets = [][]Point{
	// Straight line example from book
	// Min distance: 32
	{
		{X: 0, Y: 0},
		{X: 11, Y: 0},
		{X: 1, Y: 0},
		{X: 3, Y: 0},
		{X: -21, Y: 0},
		{X: -2, Y: 0},
		{X: -5, Y: 0},
	},
	// Square: (0,0) -> (1,0) -> (1,1) -> (0,1)
	// Min distance: 4 (perimeter of unit square)
	{
		{X: 0, Y: 0},
		{X: 1, Y: 0},
		{X: 1, Y: 1},
		{X: 0, Y: 1},
	},
	// Triangle with a point in the middle
	// Optimal path: (-1,0) -> (0,0) -> (0,1) -> (1,0)
	// Min distance: 1 + 1 + sqrt(2) ≈ 3.414
	{
		{X: 0, Y: 1},
		{X: 1, Y: 0},
		{X: -1, Y: 0},
		{X: 0, Y: 0},
	},
	// 3 collinear points
	// Optimal path: (0,0) -> (5,0) -> (10,0)
	// Min distance: 10
	{
		{X: 0, Y: 0},
		{X: 10, Y: 0},
		{X: 5, Y: 0},
	},
	// Two pairs of close points far apart
	// Optimal path: (0,0) -> (1,0) -> (10,0) -> (11,0)
	// Min distance: 12
	{
		{X: 0, Y: 0},
		{X: 11, Y: 0},
		{X: 1, Y: 0},
		{X: 10, Y: 0},
	},
}
