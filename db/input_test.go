package db

import (
	"fmt"
	"testing"
	"trainsandtowns/graph"
)

func TestTextRead(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     bool
	}{
		{"1", "../db/input1.dat", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DataRead(tt.filename)
			TestEdges := []graph.Edge{
				{StartTown: "A", EndTown: "B", Weighting: 5},
				{StartTown: "B", EndTown: "C", Weighting: 4},
				//				{StartTown: "B", EndTown: "C", Weighting: 4},
			}
			for i := 0; i < len(TestEdges); i++ {
				fmt.Println("Edges: ", graph.AllEdges[i])
				fmt.Println("Test: ", TestEdges[i])
				if got := graph.AllEdges[i].Compare(TestEdges[i]); !got {
					t.Errorf("name = %v, TextRead() = %v, want %v", tt.name, got, tt.want)
				}
			}
		})
	}
}
