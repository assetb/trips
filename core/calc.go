package core

import (
	"strings"
	"trainsandtowns/graph"
)

//Distance calculates the length of route
func Distance(route graph.Route) int {
	var sum int
	sum = 0
	for _, edge := range route {
		sum = sum + edge.Weighting
	}
	return sum
}

//GetOutTownEdges gets all edges starting from given town
func GetOutTownEdges(town string) []graph.Edge {
	edges := make([]graph.Edge, 0)
	for _, edge := range graph.AllEdges {
		if strings.EqualFold(edge.StartTown, town) {
			edges = append(edges, edge)
		}
	}
	return edges
}

//GetInTownEdges gets all edges ending in given town
func GetInTownEdges(town string) []graph.Edge {
	edges := make([]graph.Edge, 0)
	for _, edge := range graph.AllEdges {
		if strings.EqualFold(edge.EndTown, town) {
			edges = append(edges, edge)
		}
	}
	return edges
}

//MakeRouteFromTownSeq makes route over given towns
func MakeRouteFromTownSeq(seq []string) graph.Route {
	route := make(graph.Route, 0)
	for i := 1; i < len(seq); i++ {
		edge := graph.AllEdges.FindEdge(seq[i-1], seq[i])
		if edge == nil {
			return nil
		}
		//route.Append(edge)
		route = append(route, edge)
	}
	return route
}

//Search is a Universal search
func Search(start string, end string, condition func(string, string) bool) {
	for k := range graph.AllGraph[start] {
		if condition(start, end) {

		}
		Search(k, end, condition)
	}
}

//SearchMaxStops searches with maximum stops
//the code uses recursion to go through graph
//it stops when number of stops achieved
//it starts to collect edges to route when end town found
//collecting edges needed to inverse collection
func SearchMaxStops(start string, end string, nstops int) []graph.Route {
	routes := make([]graph.Route, 0)
	for k := range graph.AllGraph[start] {
		if k == end {
			edge := graph.Edge{StartTown: start, EndTown: end, Weighting: graph.AllGraph[start][end]}
			route := graph.Route{&edge}
			routes = append(routes, route)
		}
	}
	if nstops == 0 {
		return routes
	}
	for k := range graph.AllGraph[start] {
		interroutes := SearchMaxStops(k, end, nstops-1)
		if interroutes != nil {
			edge := graph.Edge{StartTown: start, EndTown: k, Weighting: graph.AllGraph[start][k]}
			for _, r := range interroutes {
				rr := graph.Route{&edge}
				for _, e := range r {
					rr = append(rr, e)
				}
				routes = append(routes, rr)
			}
		}
	}
	return routes
}

//SearchRoutesWithDistance searches with maximum stops
//the code uses recursion to go through graph
//it stops when number of stops achieved
//it starts to collect edges to route when end town found
//collecting edges needed to inverse collection
func SearchRoutesWithDistance(start string, end string, nstops int) []graph.Route {
	routes := make([]graph.Route, 0)
	for k := range graph.AllGraph[start] {
		if k == end {
			edge := graph.Edge{StartTown: start, EndTown: end, Weighting: graph.AllGraph[start][end]}
			route := graph.Route{&edge}
			routes = append(routes, route)
		}
	}
	if nstops == 0 {
		return routes
	}
	for k := range graph.AllGraph[start] {
		interroutes := SearchMaxStops(k, end, nstops-1)
		if interroutes != nil {
			edge := graph.Edge{StartTown: start, EndTown: k, Weighting: graph.AllGraph[start][k]}
			for _, r := range interroutes {
				rr := graph.Route{&edge}
				for _, e := range r {
					rr = append(rr, e)
				}
				routes = append(routes, rr)
			}
		}
	}
	return routes
}

//SearchFirstRouteIter searches first route between start town and end town
//the code uses recursion to go through graph
//it stops when end reached
//it starts to collect edges to route when end town found
//collecting edges needed to inverse collection
func SearchFirstRouteIter(start string, end string) graph.Route {
	route := make(graph.Route, 0)
	for k := range graph.AllGraph[start] {
		if k == end {
			edge := graph.Edge{StartTown: start, EndTown: end, Weighting: graph.AllGraph[start][end]}
			route = graph.Route{&edge}
			return route
		}
	}
	for k := range graph.AllGraph[start] {
		interroute := SearchFirstRouteIter(k, end)
		if interroute != nil {
			edge := graph.Edge{StartTown: start, EndTown: k, Weighting: graph.AllGraph[start][k]}
			route = graph.Route{&edge}
			for _, e := range interroute {
				route = append(route, e)
				return route
			}
		}
	}
	return nil
}

//found checks if route if found
var found bool

var distance int
var foundroute graph.Route

//Capacity controls a depth of search
var Capacity int

//interCapacity is the variable of capacity
var interCapacity int

//var next map[string]int

//SearchShortestDistance searches and calculates the shortest distance between start town and end town
//the code goes sequently from start town over all possible directions in each step calculationing distance using map structure
// each step is a recursionally called function
//it stops when capacity (the number steps) reached
//start lived for future use
func SearchShortestDistance(col map[string]int, start string, end string) int {
	next := make(map[string]int)
	for k := range col {
		if k == end {
			if !found || distance > col[k] {
				found = true
				distance = col[k]
			}
		}
		for h := range graph.AllGraph[k] {
			_, ok := next[h]
			if !ok || next[h] > col[k]+graph.AllGraph[k][h] {
				next[h] = col[k] + graph.AllGraph[k][h]
			}
		}
	}
	interCapacity++
	if interCapacity >= Capacity {
		return distance
	}
	return SearchShortestDistance(next, start, end)
}

//SearchShortestDistanceInit inits SearchShortestRoute
func SearchShortestDistanceInit(capacity int) {
	found = false
	distance = 0
	interCapacity = 0
	Capacity = capacity
	//next = make(map[string]int)
}

//SearchExactStops searches with maximum stops
//the code uses recursion to go through graph
//it stops when number of stops achieved
//it starts to collect edges to route when end town found and number of stops
//collecting edges needed to inverse collection
func SearchExactStops(start string, end string, nstops int) []graph.Route {
	routes := make([]graph.Route, 0)
	//	found := false
	for k := range graph.AllGraph[start] {
		if nstops == 0 && k == end {
			edge := graph.Edge{StartTown: start, EndTown: end, Weighting: graph.AllGraph[start][end]}
			route := graph.Route{&edge}
			routes = append(routes, route)
			//			found = true
		}
	}
	if nstops == 0 {
		//		if found {
		return routes
		//		}
		//		return nil
	}
	for k := range graph.AllGraph[start] {
		interroutes := SearchExactStops(k, end, nstops-1)
		//collecting to routes only if route exists
		if interroutes != nil {
			edge := graph.Edge{StartTown: start, EndTown: k, Weighting: graph.AllGraph[start][k]}
			for _, r := range interroutes {
				rr := graph.Route{&edge}
				for _, e := range r {
					rr = append(rr, e)
				}
				routes = append(routes, rr)
			}
		}
	}
	//	if found {
	return routes
	//	}
	//	return nil
}
