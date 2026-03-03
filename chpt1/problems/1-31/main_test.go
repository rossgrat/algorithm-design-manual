package main

import (
	"testing"
)

func TestGenerateCombinations(t *testing.T) {
	tests := []struct {
		name     string
		N, k     int
		expected int // number of combinations
	}{
		{"C(4,2)", 4, 2, 6},
		{"C(4,0)", 4, 0, 1},
		{"C(4,4)", 4, 4, 1},
		{"C(5,3)", 5, 3, 10},
		{"C(2,3) impossible", 2, 3, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateCombinations(tt.N, tt.k)
			if len(result) != tt.expected {
				t.Errorf("got %d combinations, want %d", len(result), tt.expected)
			}
		})
	}
}

func TestCheckTickets(t *testing.T) {
	tests := []struct {
		name     string
		tickets  []Ticket
		n, k, L  int
		expected bool
	}{
		{
			name:     "two tickets cover all C(4,2) with L=1",
			tickets:  []Ticket{{Numbers: []int{1, 2}}, {Numbers: []int{3, 4}}},
			n:        4,
			k:        2,
			L:        1,
			expected: true,
		},
		{
			name:     "single ticket cannot cover all C(4,2) with L=1",
			tickets:  []Ticket{{Numbers: []int{1, 2}}},
			n:        4,
			k:        2,
			L:        1,
			expected: false,
		},
		{
			name:     "all combinations purchased covers L=2",
			tickets:  []Ticket{{Numbers: []int{1, 2}}, {Numbers: []int{1, 3}}, {Numbers: []int{1, 4}}, {Numbers: []int{2, 3}}, {Numbers: []int{2, 4}}, {Numbers: []int{3, 4}}},
			n:        4,
			k:        2,
			L:        2,
			expected: true,
		},
		{
			name:     "single ticket not enough for L=2",
			tickets:  []Ticket{{Numbers: []int{1, 2}}, {Numbers: []int{3, 4}}},
			n:        4,
			k:        2,
			L:        2,
			expected: false,
		},
		{
			name:     "one ticket covers everything",
			tickets:  []Ticket{{Numbers: []int{1, 2, 3, 4, 5, 6}}},
			n:        6,
			k:        6,
			L:        1,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := checkTickets(tt.tickets, tt.n, tt.k, tt.L)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
