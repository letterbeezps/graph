package graph

import (
	"fmt"
	"testing"
)

func TestNewGraph(t *testing.T) {

	graphAdj := NewGraphAdj[string, int]()
	fmt.Println(graphAdj)

	e1 := Edge[string, int]{
		U: "s",
		V: "w",
		W: 100,
	}

	graphAdj.AddEdge(&e1)
	fmt.Println(graphAdj)

}

func TestBFS(t *testing.T) {
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

	visitFunc := func(u string) bool {
		fmt.Println("BFS-TEST: ", u)
		return true
	}

	ans := graphAdj.Bfs("s", visitFunc)

	fmt.Println(ans)

}

func TestDFS(t *testing.T) {
	graphAdj := NewGraphAdj[string, int]()

	e3 := Edge[string, int]{
		U:         "x",
		V:         "v",
		W:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e3)

	e4 := Edge[string, int]{
		U:         "v",
		V:         "y",
		W:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e4)

	e5 := Edge[string, int]{
		U:         "y",
		V:         "x",
		W:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e5)

	e1 := Edge[string, int]{
		U:         "u",
		V:         "x",
		W:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e1)

	e6 := Edge[string, int]{
		U:         "w",
		V:         "y",
		W:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e6)

	e7 := Edge[string, int]{
		U:         "w",
		V:         "z",
		W:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e7)

	e2 := Edge[string, int]{
		U:         "u",
		V:         "v",
		W:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e2)

	e8 := Edge[string, int]{
		U:         "z",
		V:         "z",
		W:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e8)

	fmt.Println(graphAdj)

	ans := graphAdj.Dfs(nil)

	fmt.Println(ans)

	dfsForest := graphAdj.DepthFirstForest()
	fmt.Println(dfsForest)
}

func TestDfsWithStart(t *testing.T) {
	graphAdj := NewGraphAdj[string, int]()

	e1 := Edge[string, int]{
		U:         "u",
		V:         "x",
		W:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e1)

	e2 := Edge[string, int]{
		U:         "u",
		V:         "v",
		W:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e2)

	e3 := Edge[string, int]{
		U:         "x",
		V:         "v",
		W:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e3)

	e4 := Edge[string, int]{
		U:         "v",
		V:         "y",
		W:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e4)

	e5 := Edge[string, int]{
		U:         "y",
		V:         "x",
		W:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e5)

	e6 := Edge[string, int]{
		U:         "w",
		V:         "y",
		W:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e6)

	e7 := Edge[string, int]{
		U:         "w",
		V:         "z",
		W:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e7)

	e8 := Edge[string, int]{
		U:         "z",
		V:         "z",
		W:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e8)

	fmt.Println(graphAdj)

	visitFunc := func(u string) bool {
		fmt.Println("DFS-TEST: ", u)
		return true
	}

	ans := graphAdj.DfsWithStart("w", visitFunc)

	fmt.Println(ans)

}

func TestTranspose(t *testing.T) {
	graphAdj := NewGraphAdj[string, int]()

	e1 := Edge[string, int]{
		U:         "r",
		V:         "v",
		W:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e1)

	e2 := Edge[string, int]{
		U:         "r",
		V:         "s",
		W:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e2)

	e3 := Edge[string, int]{
		U:         "s",
		V:         "w",
		W:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e3)

	e4 := Edge[string, int]{
		U:         "t",
		V:         "w",
		W:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e4)

	e5 := Edge[string, int]{
		U:         "t",
		V:         "u",
		W:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e5)

	e6 := Edge[string, int]{
		U:         "t",
		V:         "x",
		W:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e6)

	e7 := Edge[string, int]{
		U:         "w",
		V:         "x",
		W:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e7)

	e8 := Edge[string, int]{
		U:         "u",
		V:         "x",
		W:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e8)

	e9 := Edge[string, int]{
		U:         "y",
		V:         "x",
		W:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e9)

	e10 := Edge[string, int]{
		U:         "y",
		V:         "u",
		W:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e10)

	graphAdj.AddVertex("letter")

	fmt.Println(graphAdj)

	graphAdjT := graphAdj.Transpose()

	fmt.Println(graphAdjT)

}
