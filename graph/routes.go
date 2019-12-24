package graph

import (
	"fmt"
	"strings"
)

//Town is decribibg town of problems
type Town struct {
	Index int
	Name  string
}

//Edge represents a structure to describe route between two towns
type Edge struct {
	StartTown string
	EndTown   string
	Weighting int
}

//Route is the route represented as a collection of edges
type Route []*Edge

//Graph is a collections of elementary routes existing in the task
type Graph []Edge

//AllEdges is a collection of edges in the task
var AllEdges Graph

//Graph2 is a collections of elementary routes existing in the task
type Graph2 map[string]map[string]int

//AllGraph is a collection of elementary routes
var AllGraph Graph2

//Append edge to route
//we provide separate function for adding to make control of adding
func (route Route) Append(edge *Edge) (Route, bool) {
	if len(route) != 0 && !strings.EqualFold(route[len(route)-1].EndTown, edge.StartTown) {
		return route, false
	}
	route = append(route, edge)
	return route, true
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

//FindEdge finds edge by towns
func (graph Graph) FindEdge(start, end string) *Edge {
	for _, edge := range graph {
		if edge.Compare2(start, end) {
			return &edge
		}
	}
	return nil
}

//FindEdge finds edge by towns
func (graph Graph2) FindEdge(start, end string) *Edge {
	_, ok := graph[start][end]
	if ok {
		edge := Edge{StartTown: start, EndTown: end, Weighting: graph[start][end]}
		return &edge
	}
	return nil
}
