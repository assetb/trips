package db

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"trainsandtowns/graph"
	"trainsandtowns/middleware"
)

//DataRead reads data from a input file
func DataRead(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		middleware.Error("Error with open file ", filename, " with error ", err)
		os.Exit(1)
	}
	defer file.Close()

	graph.EdgesMap = make(map[string]map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ",")
		for _, s := range split {
			s = strings.ToUpper(strings.TrimSpace(s))
			var edge graph.Edge
			edge.StartTown = s[0:1]
			edge.EndTown = s[1:2]
			edge.Weighting, err = strconv.Atoi(s[2:3])
			if err != nil {
				middleware.Error("Error with assignment values to edgw ", err)
			}
			graph.ObjEdges = append(graph.ObjEdges, edge)
			g, ok := graph.EdgesMap[edge.StartTown]
			if !ok {
				g = make(map[string]int)
				graph.EdgesMap[edge.StartTown] = g
			}
			g[edge.EndTown] = edge.Weighting
		}
	}
	if err = scanner.Err(); err != nil {
		middleware.Error("Error with scaning file ", err)
	}

	// for {
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			break
	// 		} else {
	// 			middleware.Error("Error with reading file ", filename, " with error ", err)
	// 		}
	// 	}
	// }
}
