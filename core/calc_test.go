package core

import (
	"testing"
	"trips/graph"
)

func TestDistance(t *testing.T) {
	route1 := make(graph.Route, 0)
	edge1 := graph.Edge{StartTown: "A", EndTown: "B", Weighting: 5}
	edge2 := graph.Edge{StartTown: "B", EndTown: "C", Weighting: 4}
	edge3 := graph.Edge{StartTown: "C", EndTown: "D", Weighting: 3}
	route1 = append(route1, &edge1, &edge2, &edge3)
	tests := []struct {
		name string
		args graph.Route
		want int
	}{
		{"1", route1, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.Distance(); got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
