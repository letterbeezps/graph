package graph

import "testing"

func TestDraw(t *testing.T) {
	graphAdj := NewGraphAdj[string, int]()

	e1 := Edge[string, int]{
		U:         "r",
		V:         "v",
		W:         0,
		notDirect: true,
	}
	graphAdj.AddEdge(&e1)

	e2 := Edge[string, int]{
		U:         "r",
		V:         "s",
		W:         0,
		notDirect: true,
	}
	graphAdj.AddEdge(&e2)

	e3 := Edge[string, int]{
		U:         "s",
		V:         "w",
		W:         0,
		notDirect: true,
	}
	graphAdj.AddEdge(&e3)

	e4 := Edge[string, int]{
		U:         "t",
		V:         "w",
		W:         0,
		notDirect: true,
	}
	graphAdj.AddEdge(&e4)

	e5 := Edge[string, int]{
		U:         "t",
		V:         "u",
		W:         0,
		notDirect: true,
	}
	graphAdj.AddEdge(&e5)

	e6 := Edge[string, int]{
		U:         "t",
		V:         "x",
		W:         0,
		notDirect: true,
	}
	graphAdj.AddEdge(&e6)

	e7 := Edge[string, int]{
		U:         "w",
		V:         "x",
		W:         0,
		notDirect: true,
	}
	graphAdj.AddEdge(&e7)

	e8 := Edge[string, int]{
		U:         "u",
		V:         "x",
		W:         0,
		notDirect: true,
	}
	graphAdj.AddEdge(&e8)

	e9 := Edge[string, int]{
		U:         "y",
		V:         "x",
		W:         0,
		notDirect: true,
	}
	graphAdj.AddEdge(&e9)

	e10 := Edge[string, int]{
		U:         "y",
		V:         "u",
		W:         0,
		notDirect: true,
	}
	graphAdj.AddEdge(&e10)

	graphAdj.Draw(false)

}
