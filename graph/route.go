package graph

import (
	"strings"
)

//Town is decribibg town of problems
type Town struct {
	Index int
	Name  string
}

//Route is the route represented as a collection of edges
type Route []*Edge

//Append edge to route
//we provide separate function for adding to make control of adding
func (route Route) Append(edge *Edge) (Route, bool) {
	if len(route) != 0 && !strings.EqualFold(route[len(route)-1].EndTown, edge.StartTown) {
		return route, false
	}
	route = append(route, edge)
	return route, true
}

//Distance calculates the length of route
func (route Route) Distance() int {
	var sum int
	sum = 0
	for _, edge := range route {
		sum = sum + edge.Weighting
	}
	return sum
}

//GetOutTownEdges gets all edges starting from given town
func (town Town) GetOutTownEdges() []Edge {
	edges := make([]Edge, 0)
	for _, edge := range ObjEdges {
		if strings.EqualFold(edge.StartTown, town.Name) {
			edges = append(edges, edge)
		}
	}
	return edges
}

//GetInTownEdges gets all edges ending in given town
func (town Town) GetInTownEdges() []Edge {
	edges := make([]Edge, 0)
	for _, edge := range ObjEdges {
		if strings.EqualFold(edge.EndTown, town.Name) {
			edges = append(edges, edge)
		}
	}
	return edges
}
