package graph

//ObjEdgesType is a type of collection of elementary routes existing in the task run in object format - 1-st approach
type ObjEdgesType []Edge

//ObjEdges is a collection of edges (elementary toutes) existing in the task run in object format - 1-st approach
var ObjEdges ObjEdgesType

//EdgesMapType is a type for collection of elementary routes existing in the task in map format - 2-nd approach
type EdgesMapType map[string]map[string]int

//EdgesMap is a collection of elementary routes existing in the run in map format - 2-nd approach
var EdgesMap EdgesMapType

//FindEdge finds edge by towns
func (graph ObjEdgesType) FindEdge(start, end string) *Edge {
	for _, edge := range graph {
		if edge.Compare2(start, end) {
			return &edge
		}
	}
	return nil
}

//FindEdge finds edge by towns
func (graph EdgesMapType) FindEdge(start, end string) *Edge {
	_, ok := graph[start][end]
	if ok {
		edge := Edge{StartTown: start, EndTown: end, Weighting: graph[start][end]}
		return &edge
	}
	return nil
}
