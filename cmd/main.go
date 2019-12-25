package main

import (
	"fmt"
	"strings"
	"trainsandtowns/core"
	"trainsandtowns/db"
	"trainsandtowns/graph"
)

func main() {
	db.DataRead("../db/input1.dat")

	//========= 1. The distance of the route A-B-C 		=====
	fmt.Print("#1. ", RouteDistance("A-B-C"), "\n")

	//========= 2. The distance of the route A-D 		=====
	fmt.Print("#2. ", RouteDistance("A-D"), "\n")

	//========= 3. The distance of the route A-D-C 		=====
	fmt.Print("#3. ", RouteDistance("A-D-C"), "\n")

	//========= 4. The distance of the route A-E-B-C-D 	=====
	fmt.Print("#4. ", RouteDistance("A-E-B-C-D"), "\n")

	//========= 5. The distance of the route A-E-D 		=====
	fmt.Print("#5. ", RouteDistance("A-E-D"), "\n")

	//========= 6. The number of trips starting at C and ending at C with a maximum of 3 stops.
	fmt.Print("#6. ")
	NTripsWithMaxStops("C", "C", 3)

	//==========7. The number of trips starting at A and ending at C with exactly 4 stops.
	fmt.Print("#7. ")
	NTripsWithExactStops("A", "C", 4)

	//==========8. The length of the shortest route (in terms of distance to travel) from A to C.
	fmt.Print("#8. ", ShortestRouteDistance("A", "C"), "\n")

	//==========9. The length of the shortest route (in terms of distance to travel) from B to B.
	fmt.Print("#9. ", ShortestRouteDistance("B", "B"), "\n")

	//==========10.The number of different routes from C to C with a distance of less than 30.
	fmt.Print("#10. ")
	NRoutesWithLessDistance("C", "C", 30)

	var delay string
	fmt.Scanln(&delay)
}

//RouteDistance prints the distance of given route
func RouteDistance(sample string) string {
	split := strings.Split(sample, "-")
	route := core.MakeRouteFromTownSeq(split)
	if route == nil {
		return fmt.Sprintf("NO SUCH ROUTE")
	}
	d := core.Distance(route)
	return fmt.Sprintf("The distance of the route %v is %v", sample, d)
}

//NTripsWithMaxStops calcs a number of trips from start to end with max n stops
func NTripsWithMaxStops(start string, end string, nstops int) {
	routes := core.SearchMaxStops(start, end, nstops-1)
	fmt.Print("The number of trips starting at ", start, " and ending  ", end, " with a maximum of ", nstops, " stops = ", len(routes), " They are:\n")
	for i, route := range routes {
		fmt.Print("     route ", i, ": ")
		for _, edge := range route {
			fmt.Print(edge.Print(), " ")
		}
		fmt.Print("\n")
	}
}

//NTripsWithExactStops calcs a number of trips from start to end with exactly n stops
func NTripsWithExactStops(start string, end string, nstops int) {
	interroutes := core.SearchMaxStops(start, end, nstops-1)
	routes := make([]graph.Route, 0)
	for _, r := range interroutes {
		if len(r) == nstops {
			routes = append(routes, r)
		}
	}
	fmt.Print("The number of trips starting at ", start, " and ending  ", end, " with exactly ", nstops, " stops = ", len(routes), " They are:\n")
	for i, route := range routes {
		fmt.Print("     route ", i, ": ")
		for _, edge := range route {
			fmt.Print(edge.Print(), " ")
		}
		fmt.Print("\n")
	}
}

//ShortestRouteDistance calcs length of shortest route
func ShortestRouteDistance(start string, end string) string {
	core.SearchShortestDistanceInit(14)
	d := core.SearchShortestDistance(graph.AllGraph[start], start, end)
	return fmt.Sprint("The length of the shortest route (in terms of distance to travel) from ", start, " to ", end, " = ", d)
}

//NRoutesWithLessDistance calcs length of shortest route
func NRoutesWithLessDistance(start string, end string, ifdist int) {
	core.SearchRouteWithLessDistanceInit(14)
	d, routes := core.SearchRouteWithLessDistance(graph.AllGraph[start], start, end, ifdist)
	fmt.Print("The number of different routes from ", start, " to ", end, " with a distance of less than ", ifdist, " = ", d, " They are:\n")
	for i, route := range routes {
		fmt.Print("     route ", i, ": ", route)
		fmt.Print("\n")
	}
}
