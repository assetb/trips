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

//SearchShortestRoute searches first route between start town and end town
//the code uses recursion to go through graph
//it stops when end reached
//it starts to collect edges to route when end town found
//collecting edges needed to inverse collection
func SearchShortestRoute(col map[string]int, start string, end string) graph.Route {
	next := make(map[string]int)
	for k := range col {
		if k == end {
			edge := graph.Edge{StartTown: start, EndTown: end, Weighting: graph.AllGraph[start][end]}
			route := graph.Route{&edge}
			return route
		}
		for h := range graph.AllGraph[k] {
			next[h] = graph.AllGraph[k][h]
		}
	}
	return SearchShortestRoute(next, start, end)
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
