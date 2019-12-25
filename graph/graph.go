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
func (edges ObjEdgesType) FindEdge(start, end string) *Edge {
	for _, edge := range edges {
		if edge.CompareS(start, end) {
			return &edge
		}
	}
	return nil
}

//FindEdge finds edge by towns
func (edges EdgesMapType) FindEdge(start, end string) *Edge {
	_, ok := edges[start][end]
	if ok {
		edge := Edge{StartTown: start, EndTown: end, Weighting: edges[start][end]}
		return &edge
	}
	return nil
}
