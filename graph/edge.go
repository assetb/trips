package graph

import (
	"fmt"
	"strings"
)

//Edge represents a structure to describe route between two towns
type Edge struct {
	StartTown string
	EndTown   string
	Weighting int
}

//Print prints edge
func (edge Edge) Print() string {
	return fmt.Sprint(edge.StartTown, edge.EndTown, edge.Weighting)
}

//Compare edges
func (edge Edge) Compare(comp Edge) bool {
	if strings.EqualFold(edge.StartTown, comp.StartTown) && strings.EqualFold(edge.EndTown, comp.EndTown) {
		return true
	}
	return false
}

//Compare2 compares edge versus towns
func (edge Edge) Compare2(start, end string) bool {
	if strings.EqualFold(edge.StartTown, start) && strings.EqualFold(edge.EndTown, end) {
		return true
	}
	return false
}
