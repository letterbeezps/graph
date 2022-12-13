package graph

import (
	"fmt"
	"testing"
)

func TestDepthFirstForest(t *testing.T) {
	graphAdj := NewGraphAdj[string, int]()

	e3 := Edge[string, int]{
		u:         "x",
		v:         "v",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e3)

	e5 := Edge[string, int]{
		u:         "y",
		v:         "x",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e5)

	e4 := Edge[string, int]{
		u:         "v",
		v:         "y",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e4)

	e8 := Edge[string, int]{
		u:         "z",
		v:         "z",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e8)

	e1 := Edge[string, int]{
		u:         "u",
		v:         "x",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e1)

	e6 := Edge[string, int]{
		u:         "w",
		v:         "y",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e6)

	e7 := Edge[string, int]{
		u:         "w",
		v:         "z",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e7)

	e2 := Edge[string, int]{
		u:         "u",
		v:         "v",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e2)

	fmt.Println(graphAdj)

	ans := graphAdj.DepthFirstForest()

	fmt.Println(ans)

}

func TestStrongConnectedComponents(t *testing.T) {
	graphAdj := NewGraphAdj[string, int]()

	e3 := Edge[string, int]{
		u:         "a",
		v:         "b",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e3)

	e5 := Edge[string, int]{
		u:         "b",
		v:         "c",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e5)

	e1 := Edge[string, int]{
		u:         "b",
		v:         "e",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e1)

	e6 := Edge[string, int]{
		u:         "b",
		v:         "f",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e6)

	e4 := Edge[string, int]{
		u:         "c",
		v:         "d",
		w:         0,
		notDirect: true,
	}
	graphAdj.AddEdge(&e4)

	e8 := Edge[string, int]{
		u:         "c",
		v:         "g",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e8)

	e7 := Edge[string, int]{
		u:         "d",
		v:         "h",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e7)

	e2 := Edge[string, int]{
		u:         "h",
		v:         "h",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e2)

	e11 := Edge[string, int]{
		u:         "g",
		v:         "h",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e11)

	e12 := Edge[string, int]{
		u:         "g",
		v:         "f",
		w:         0,
		notDirect: true,
	}
	graphAdj.AddEdge(&e12)

	e13 := Edge[string, int]{
		u:         "e",
		v:         "f",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e13)

	e14 := Edge[string, int]{
		u:         "e",
		v:         "a",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e14)

	fmt.Println(graphAdj)

	ans := graphAdj.StrongConnectedComponents()

	fmt.Println(ans)

	graphAdj.Draw(false)

}
