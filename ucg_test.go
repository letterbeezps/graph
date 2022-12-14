package graph

import "testing"

func Test_undirectedGraph(t *testing.T) {
	g := NewGraphAdj[string, int]()
	g.NotDirect = true
	e1 := Edge[string, int]{
		U:         "a",
		V:         "b",
		W:         4,
		NotDirect: false,
	}
	g.AddEdge(&e1)
}
