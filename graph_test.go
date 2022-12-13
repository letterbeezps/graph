package graph

import (
	"fmt"
	"testing"
)

func TestNewGraph(t *testing.T) {

	graphAdj := NewGraphAdj[string, int]()
	fmt.Println(graphAdj)

	e1 := Edge[string, int]{
		u: "s",
		v: "w",
		w: 100,
	}

	graphAdj.AddEdge(&e1)
	fmt.Println(graphAdj)

}

func TestBFS(t *testing.T) {
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
		u:         "x",
		v:         "v",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e3)

	e4 := Edge[string, int]{
		u:         "v",
		v:         "y",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e4)

	e5 := Edge[string, int]{
		u:         "y",
		v:         "x",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e5)

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

	e8 := Edge[string, int]{
		u:         "z",
		v:         "z",
		w:         0,
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
		u:         "u",
		v:         "x",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e1)

	e2 := Edge[string, int]{
		u:         "u",
		v:         "v",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e2)

	e3 := Edge[string, int]{
		u:         "x",
		v:         "v",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e3)

	e4 := Edge[string, int]{
		u:         "v",
		v:         "y",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e4)

	e5 := Edge[string, int]{
		u:         "y",
		v:         "x",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e5)

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

	e8 := Edge[string, int]{
		u:         "z",
		v:         "z",
		w:         0,
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
		u:         "r",
		v:         "v",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e1)

	e2 := Edge[string, int]{
		u:         "r",
		v:         "s",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e2)

	e3 := Edge[string, int]{
		u:         "s",
		v:         "w",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e3)

	e4 := Edge[string, int]{
		u:         "t",
		v:         "w",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e4)

	e5 := Edge[string, int]{
		u:         "t",
		v:         "u",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e5)

	e6 := Edge[string, int]{
		u:         "t",
		v:         "x",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e6)

	e7 := Edge[string, int]{
		u:         "w",
		v:         "x",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e7)

	e8 := Edge[string, int]{
		u:         "u",
		v:         "x",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e8)

	e9 := Edge[string, int]{
		u:         "y",
		v:         "x",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e9)

	e10 := Edge[string, int]{
		u:         "y",
		v:         "u",
		w:         0,
		notDirect: false,
	}
	graphAdj.AddEdge(&e10)

	graphAdj.AddVertex("letter")

	fmt.Println(graphAdj)

	graphAdjT := graphAdj.Transpose()

	fmt.Println(graphAdjT)

}
