package core

import (
	"fmt"
	"trainsandtowns/graph"
)

//MakeRouteFromStrSeq makes route from the string of given towns
func MakeRouteFromStrSeq(seq []string) graph.Route {
	route := make(graph.Route, 0)
	for i := 1; i < len(seq); i++ {
		edge := graph.ObjEdges.FindEdge(seq[i-1], seq[i])
		if edge == nil {
			return nil
		}
		route = append(route, edge)
	}
	return route
}

//SearchMaxStops searches routes between start town and end town with maximum n stops
//the code uses recursion to go through graph
//it stops when number of stops achieved
//it starts to collect edges to route when end town found
//collecting edges needed to inverse collection
func SearchMaxStops(start string, end string, nstops int) []graph.Route {
	routes := make([]graph.Route, 0)
	for k := range graph.EdgesMap[start] {
		if k == end {
			edge := graph.Edge{StartTown: start, EndTown: end, Weighting: graph.EdgesMap[start][end]}
			route := graph.Route{&edge}
			routes = append(routes, route)
		}
	}
	if nstops == 0 {
		return routes
	}
	for k := range graph.EdgesMap[start] {
		interroutes := SearchMaxStops(k, end, nstops-1)
		if interroutes != nil {
			edge := graph.Edge{StartTown: start, EndTown: k, Weighting: graph.EdgesMap[start][k]}
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

//SearchShortestDistance searches and calculates the shortest distance between start town and end town
//code description see at SearchShortestDistanceImpl function comments
func SearchShortestDistance(start string, end string) (int, string) {
	distance = -1
	interCapacity = 0
	if Capacity <= 0 {
		Capacity = 20
	}
	d, r := searchShortestDistanceImpl(graph.EdgesMap[start], end)
	r = fmt.Sprint(start, r)
	return d, r
}

//SearchRoutesWithDistanceLessThen searches and calculates the shortest distance between start town and end town
//code description see at SearchRoutesWithDistanceLessThen function comments
func SearchRoutesWithDistanceLessThen(start string, end string, ifdist int) (int, []string) {
	nroutewithlessdistance = 0
	interCapacity = 0
	if Capacity <= 0 {
		Capacity = 20
	}
	routes = make([]string, 0)
	d, rr := searchRoutesWithDistanceLessThenImpl(graph.EdgesMap[start], end, ifdist)
	rrr := make([]string, 0)
	for _, r := range rr {
		r = fmt.Sprint(start, r)
		rrr = append(rrr, r)
	}
	return d, rrr
}

//Capacity controls a depth of search
var Capacity int

//SearchCapacityInit inits search capacity
func SearchCapacityInit(capacity int) {
	Capacity = capacity
}

var distance int
var route string
var interCapacity int

//searchShortestDistanceImpl searches and calculates the shortest distance between start town and end town
//the code goes sequently from start town over all possible directions in each step calculationing distance using map structure
// each step is a recursionally called function
//it stops when capacity (the number steps) reached
//start lived for future use
func searchShortestDistanceImpl(col map[string]int, end string) (int, string) {
	next := make(map[string]int)
	for k := range col {
		l := k[len(k)-1:]
		if l == end {
			if distance < 0 || distance > col[k] {
				distance = col[k]
				route = k
			}
		}
		for h := range graph.EdgesMap[l] {
			next[k+h] = col[k] + graph.EdgesMap[l][h]
		}
	}
	interCapacity++
	if interCapacity >= Capacity {
		return distance, route
	}
	return searchShortestDistanceImpl(next, end)
}

var nroutewithlessdistance int
var routes []string

//searchRoutesWithDistanceLessThenImpl searches and calculates the shortest distance between start town and end town
//the code goes sequently from start town over all possible directions in each step calculationing distance using map structure
// each step is a recursionally called function
//it stops when capacity (the number steps) reached
//start lived for future use
func searchRoutesWithDistanceLessThenImpl(col map[string]int, end string, ifdist int) (int, []string) {
	next := make(map[string]int)
	for k := range col {
		l := k[len(k)-1:]
		if l == end {
			if col[k] < ifdist {
				nroutewithlessdistance++
				routes = append(routes, k)
			}
		}
		for h := range graph.EdgesMap[l] {
			next[k+h] = col[k] + graph.EdgesMap[l][h]
		}
	}
	interCapacity++
	if interCapacity >= Capacity {
		return nroutewithlessdistance, routes
	}
	searchRoutesWithDistanceLessThenImpl(next, end, ifdist)
	return nroutewithlessdistance, routes
}

//SearchExactStops searches with maximum stops
//the code uses recursion to go through graph
//it stops when number of stops achieved
//it starts to collect edges to route when end town found and number of stops
//collecting edges needed to inverse collection
func SearchExactStops(start string, end string, nstops int) []graph.Route {
	routes := make([]graph.Route, 0)
	//	found := false
	for k := range graph.EdgesMap[start] {
		if nstops == 0 && k == end {
			edge := graph.Edge{StartTown: start, EndTown: end, Weighting: graph.EdgesMap[start][end]}
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
	for k := range graph.EdgesMap[start] {
		interroutes := SearchExactStops(k, end, nstops-1)
		//collecting to routes only if route exists
		if interroutes != nil {
			edge := graph.Edge{StartTown: start, EndTown: k, Weighting: graph.EdgesMap[start][k]}
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

//Search is a Universal search
func Search(start string, end string, condition func(string, string) bool) {
	for k := range graph.EdgesMap[start] {
		if condition(start, end) {

		}
		Search(k, end, condition)
	}
}
