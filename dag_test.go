package graph

import (
	"fmt"
	"testing"
)

func TestTopologicalSort(t *testing.T) {
	graphAdj := NewGraphAdj[string, int]()

	e3 := Edge[string, int]{
		u:         "shirt",
		v:         "tie",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e3)

	e5 := Edge[string, int]{
		u:         "tie",
		v:         "jacket",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e5)

	e4 := Edge[string, int]{
		u:         "shirt",
		v:         "belt",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e4)

	e8 := Edge[string, int]{
		u:         "belt",
		v:         "jacket",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e8)

	e1 := Edge[string, int]{
		u:         "panties",
		v:         "pants",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e1)

	e9 := Edge[string, int]{
		u:         "panties",
		v:         "shoe",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e9)

	e6 := Edge[string, int]{
		u:         "pants",
		v:         "belt",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e6)

	e7 := Edge[string, int]{
		u:         "pants",
		v:         "shoe",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e7)

	e2 := Edge[string, int]{
		u:         "socks",
		v:         "shoe",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e2)

	graphAdj.AddVertex("watch")

	fmt.Println(graphAdj)

	ret := graphAdj.TopologicalSort()

	fmt.Println(ret)

}
