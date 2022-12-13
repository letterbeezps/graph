package graph

import (
	"errors"
	"log"
	"strconv"
)

type ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type vertex interface {
	string | int
}

func VertexToString(x interface{}) (string, error) {
	switch typex := x.(type) {
	case string:
		return typex, nil
	case int:
		return strconv.Itoa(typex), nil
	}
	return "", errors.New("only allowed string or int")
}

type Edge[Vertex vertex, Weight ordered] struct {
	U         Vertex
	V         Vertex
	W         Weight
	notDirect bool
}

type E[Vertex vertex, Weight ordered] struct {
	V Vertex
	W Weight
}

type GraphAdj[Vertex vertex, Weight ordered] struct {
	data     map[Vertex][]*E[Vertex, Weight]
	edges    []*Edge[Vertex, Weight]
	HasCycle bool
	Logger   *log.Logger
}

type DfsRet[Vertex vertex] struct {
	V          Vertex
	findTime   int
	finishTime int
}

func NewGraphAdj[Vertex vertex, Weight ordered]() GraphAdj[Vertex, Weight] {
	return GraphAdj[Vertex, Weight]{
		data:   map[Vertex][]*E[Vertex, Weight]{},
		edges:  []*Edge[Vertex, Weight]{},
		Logger: log.Default(),
	}
}

func (g GraphAdj[Vertex, Weight]) AddEdge(e *Edge[Vertex, Weight]) {
	g.data[e.U] = append(g.data[e.U], &E[Vertex, Weight]{
		V: e.V,
		W: e.W,
	})

	if e.notDirect {
		g.data[e.V] = append(g.data[e.V], &E[Vertex, Weight]{
			V: e.U,
			W: e.W,
		})
	}
}

func (g GraphAdj[Vertex, Weight]) AddVertex(u Vertex) {
	g.data[u] = []*E[Vertex, Weight]{}
}

// reverse the edge in Graph g and return a new Graph
func (g GraphAdj[Vertex, Weight]) Transpose() GraphAdj[Vertex, Weight] {
	ret := NewGraphAdj[Vertex, Weight]()

	for u, edges := range g.data {
		if len(edges) > 0 {
			for _, e := range edges {
				ret.data[e.V] = append(ret.data[e.V], &E[Vertex, Weight]{
					V: u,
					W: e.W,
				})
			}
		} else {
			ret.AddVertex(u)
		}

	}

	return ret
}
