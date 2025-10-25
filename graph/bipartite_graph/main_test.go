package main

import "testing"

func TestIsBipartite(t *testing.T) {
	tests := []struct {
		name     string
		graph    [][]int
		expected bool
	}{
		{
			name: "Simple Bipartite Graph",
			graph: [][]int{
				{1, 3},
				{0, 2},
				{1, 3},
				{0, 2},
			},
			expected: true,
		},
		{
			name: "Non-Bipartite Graph (Odd Cycle)",
			graph: [][]int{
				{1, 2},
				{0, 2},
				{0, 1},
			},
			expected: false,
		},
		{
			name:     "6-node Bipartite Graph (Edge List)",
			graph:    [][]int{{0, 1}, {1, 2}, {2, 5}, {0, 3}, {1, 4}, {3, 4}, {4, 2}},
			expected: false,
		},
		//{
		//	name: "Disconnected Bipartite Graph",
		//	graph: [][]int{
		//		{1},
		//		{0},
		//		{3},
		//		{2},
		//	},
		//	expected: true,
		//},
		//{
		//	name: "Single Node",
		//	graph: [][]int{
		//		{},
		//	},
		//	expected: true,
		//},
		//{
		//	name: "Graph with Self-loop",
		//	graph: [][]int{
		//		{0},
		//	},
		//	expected: false,
		//},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := IsBipartite(tc.graph)
			if result != tc.expected {
				t.Errorf("Test %q failed: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}
