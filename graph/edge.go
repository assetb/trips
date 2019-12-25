package graph

import (
	"fmt"
	"strings"
)

//Edge represents a structure to describe direct route between two towns without in-between towns/stops - elementary route
type Edge struct {
	StartTown string
	EndTown   string
	Weighting int
}

//Print prints edge
func (edge Edge) Print() string {
	return fmt.Sprint(edge.StartTown, edge.EndTown, edge.Weighting)
}

//CompareE compares edge versus another edge
func (edge Edge) CompareE(comp Edge) bool {
	if strings.EqualFold(edge.StartTown, comp.StartTown) && strings.EqualFold(edge.EndTown, comp.EndTown) {
		return true
	}
	return false
}

//CompareS compares edge versus two towns
func (edge Edge) CompareS(start, end string) bool {
	if strings.EqualFold(edge.StartTown, start) && strings.EqualFold(edge.EndTown, end) {
		return true
	}
	return false
}
