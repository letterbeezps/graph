package graph

import "testing"

func TestDraw(t *testing.T) {
	graphAdj := NewGraphAdj[string, int]()

	e1 := Edge[string, int]{
		u:         "r",
		v:         "v",
		w:         0,
		notDirect: true,
	}
	graphAdj.AddEdge(&e1)

	e2 := Edge[string, int]{
		u:         "r",
		v:         "s",
		w:         0,
		notDirect: true,
	}
	graphAdj.AddEdge(&e2)

	e3 := Edge[string, int]{
		u:         "s",
		v:         "w",
		w:         0,
		notDirect: true,
	}
	graphAdj.AddEdge(&e3)

	e4 := Edge[string, int]{
		u:         "t",
		v:         "w",
		w:         0,
		notDirect: true,
	}
	graphAdj.AddEdge(&e4)

	e5 := Edge[string, int]{
		u:         "t",
		v:         "u",
		w:         0,
		notDirect: true,
	}
	graphAdj.AddEdge(&e5)

	e6 := Edge[string, int]{
		u:         "t",
		v:         "x",
		w:         0,
		notDirect: true,
	}
	graphAdj.AddEdge(&e6)

	e7 := Edge[string, int]{
		u:         "w",
		v:         "x",
		w:         0,
		notDirect: true,
	}
	graphAdj.AddEdge(&e7)

	e8 := Edge[string, int]{
		u:         "u",
		v:         "x",
		w:         0,
		notDirect: true,
	}
	graphAdj.AddEdge(&e8)

	e9 := Edge[string, int]{
		u:         "y",
		v:         "x",
		w:         0,
		notDirect: true,
	}
	graphAdj.AddEdge(&e9)

	e10 := Edge[string, int]{
		u:         "y",
		v:         "u",
		w:         0,
		notDirect: true,
	}
	graphAdj.AddEdge(&e10)

	graphAdj.Draw(false)

}
