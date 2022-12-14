package graph

import (
	"fmt"
	"testing"
)

func TestTopologicalSort(t *testing.T) {
	graphAdj := NewGraphAdj[string, int]()

	e3 := Edge[string, int]{
		U:         "shirt",
		V:         "tie",
		W:         0,
		NotDirect: false,
	}
	graphAdj.AddEdge(&e3)

	e5 := Edge[string, int]{
		U:         "tie",
		V:         "jacket",
		W:         0,
		NotDirect: false,
	}
	graphAdj.AddEdge(&e5)

	e4 := Edge[string, int]{
		U:         "shirt",
		V:         "belt",
		W:         0,
		NotDirect: false,
	}
	graphAdj.AddEdge(&e4)

	e8 := Edge[string, int]{
		U:         "belt",
		V:         "jacket",
		W:         0,
		NotDirect: false,
	}
	graphAdj.AddEdge(&e8)

	e1 := Edge[string, int]{
		U:         "panties",
		V:         "pants",
		W:         0,
		NotDirect: false,
	}
	graphAdj.AddEdge(&e1)

	e9 := Edge[string, int]{
		U:         "panties",
		V:         "shoe",
		W:         0,
		NotDirect: false,
	}
	graphAdj.AddEdge(&e9)

	e6 := Edge[string, int]{
		U:         "pants",
		V:         "belt",
		W:         0,
		NotDirect: false,
	}
	graphAdj.AddEdge(&e6)

	e7 := Edge[string, int]{
		U:         "pants",
		V:         "shoe",
		W:         0,
		NotDirect: false,
	}
	graphAdj.AddEdge(&e7)

	e2 := Edge[string, int]{
		U:         "socks",
		V:         "shoe",
		W:         0,
		NotDirect: false,
	}
	graphAdj.AddEdge(&e2)

	graphAdj.AddVertex("watch")

	fmt.Println(graphAdj)

	ret := graphAdj.TopologicalSort()

	fmt.Println(ret)

}
