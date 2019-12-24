package graph

import (
	"reflect"
	"testing"
)

func TestRoute_Append(t *testing.T) {
	route := make(Route, 0)
	edge1 := &Edge{StartTown: "A", EndTown: "B", Weighting: 5}
	edge2 := &Edge{StartTown: "B", EndTown: "C", Weighting: 4}
	edge3 := &Edge{StartTown: "C", EndTown: "D", Weighting: 3}

	tests := []struct {
		name string
		edge *Edge
		want bool
	}{
		{"1", edge1, true},
		{"2", edge3, false},
		{"3", edge2, true},
		{"4", edge2, false},
		{"5", edge3, true},
	}
	for _, tt := range tests {
		var got bool
		if route, got = route.Append(tt.edge); got != tt.want {
			t.Errorf("name = %v, Route.Append() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestGraph_FindEdge(t *testing.T) {
	edge1 := Edge{StartTown: "A", EndTown: "B", Weighting: 1}
	edge2 := Edge{StartTown: "C", EndTown: "D", Weighting: 1}
	var g Graph
	g = append(g, edge1)
	g = append(g, edge2)
	type args struct {
		start string
		end   string
	}
	tests := []struct {
		name string
		g    Graph
		args args
		want *Edge
	}{
		{"1", g, args{start: "C", end: "D"}, &edge2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.FindEdge(tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Graph.FindEdge() = %v, want %v", got, tt.want)
			}
		})
	}
}
