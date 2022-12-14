package graph

import (
	"fmt"
	"testing"
)

func TestDepthFirstForest(t *testing.T) {
	graphAdj := NewGraphAdj[string, int]()

	e3 := Edge[string, int]{
		U:         "x",
		V:         "v",
		W:         0,
		NotDirect: false,
	}
	graphAdj.AddEdge(&e3)

	e5 := Edge[string, int]{
		U:         "y",
		V:         "x",
		W:         0,
		NotDirect: false,
	}
	graphAdj.AddEdge(&e5)

	e4 := Edge[string, int]{
		U:         "v",
		V:         "y",
		W:         0,
		NotDirect: false,
	}
	graphAdj.AddEdge(&e4)

	e8 := Edge[string, int]{
		U:         "z",
		V:         "z",
		W:         0,
		NotDirect: false,
	}
	graphAdj.AddEdge(&e8)

	e1 := Edge[string, int]{
		U:         "u",
		V:         "x",
		W:         0,
		NotDirect: false,
	}
	graphAdj.AddEdge(&e1)

	e6 := Edge[string, int]{
		U:         "w",
		V:         "y",
		W:         0,
		NotDirect: false,
	}
	graphAdj.AddEdge(&e6)

	e7 := Edge[string, int]{
		U:         "w",
		V:         "z",
		W:         0,
		NotDirect: false,
	}
	graphAdj.AddEdge(&e7)

	e2 := Edge[string, int]{
		U:         "u",
		V:         "v",
		W:         0,
		NotDirect: false,
	}
	graphAdj.AddEdge(&e2)

	fmt.Println(graphAdj)

	ans := graphAdj.DepthFirstForest()

	fmt.Println(ans)

}

func TestStrongConnectedComponents(t *testing.T) {
	graphAdj := NewGraphAdj[string, int]()

	e3 := Edge[string, int]{
		U:         "a",
		V:         "b",
		W:         0,
		NotDirect: false,
	}
	graphAdj.AddEdge(&e3)

	e5 := Edge[string, int]{
		U:         "b",
		V:         "c",
		W:         0,
		NotDirect: false,
	}
	graphAdj.AddEdge(&e5)

	e1 := Edge[string, int]{
		U:         "b",
		V:         "e",
		W:         0,
		NotDirect: false,
	}
	graphAdj.AddEdge(&e1)

	e6 := Edge[string, int]{
		U:         "b",
		V:         "f",
		W:         0,
		NotDirect: false,
	}
	graphAdj.AddEdge(&e6)

	e4 := Edge[string, int]{
		U:         "c",
		V:         "d",
		W:         0,
		NotDirect: true,
	}
	graphAdj.AddEdge(&e4)

	e8 := Edge[string, int]{
		U:         "c",
		V:         "g",
		W:         0,
		NotDirect: false,
	}
	graphAdj.AddEdge(&e8)

	e7 := Edge[string, int]{
		U:         "d",
		V:         "h",
		W:         0,
		NotDirect: false,
	}
	graphAdj.AddEdge(&e7)

	e2 := Edge[string, int]{
		U:         "h",
		V:         "h",
		W:         0,
		NotDirect: false,
	}
	graphAdj.AddEdge(&e2)

	e11 := Edge[string, int]{
		U:         "g",
		V:         "h",
		W:         0,
		NotDirect: false,
	}
	graphAdj.AddEdge(&e11)

	e12 := Edge[string, int]{
		U:         "g",
		V:         "f",
		W:         0,
		NotDirect: true,
	}
	graphAdj.AddEdge(&e12)

	e13 := Edge[string, int]{
		U:         "e",
		V:         "f",
		W:         0,
		NotDirect: false,
	}
	graphAdj.AddEdge(&e13)

	e14 := Edge[string, int]{
		U:         "e",
		V:         "a",
		W:         0,
		NotDirect: false,
	}
	graphAdj.AddEdge(&e14)

	fmt.Println(graphAdj)

	ans := graphAdj.StrongConnectedComponents()

	fmt.Println(ans)

	graphAdj.Draw(false)

}
