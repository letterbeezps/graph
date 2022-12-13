package graph

import (
	"sort"
)

func (g GraphAdj[Vertex, Weight]) Bfs(s Vertex, visitFunc func(Vertex) bool) []Vertex {
	visited := make(map[Vertex]struct{})
	ret, stack := []Vertex{}, []Vertex{}
	visited[s] = struct{}{}
	ret = append(ret, s)

	// do something
	if visitFunc != nil {
		visitFunc(s)
	}

	stack = append(stack, s)

	for len(stack) > 0 {
		u := stack[0]
		stack = stack[1:]

		for _, e := range g.data[u] {
			if _, ok := visited[e.V]; !ok {
				visited[e.V] = struct{}{}
				ret = append(ret, e.V)
				stack = append(stack, e.V)
				if visitFunc != nil {
					visitFunc(e.V)
				}
			} else {
				g.HasCycle = true
			}
		}
	}

	return ret
}

func (g GraphAdj[Vertex, Weight]) Dfs(visitFunc func(Vertex) bool) []DfsRet[Vertex] {
	visited := make(map[Vertex]struct{})
	ret := []DfsRet[Vertex]{}
	time := 0

	for u := range g.data {
		if _, ok := visited[u]; !ok {
			g.dfsVisit(u, visited, &ret, &time, visitFunc)
		} else {
			g.HasCycle = true
		}

	}
	return ret
}

func (g GraphAdj[Vertex, Weight]) DfsWithStart(
	u Vertex,
	visitFunc func(Vertex) bool,
) []DfsRet[Vertex] {
	visited := make(map[Vertex]struct{})
	ret := []DfsRet[Vertex]{}

	time := 0

	g.dfsVisit(u, visited, &ret, &time, visitFunc)

	return ret
}

func (g GraphAdj[Vertex, Weight]) dfsVisit(
	u Vertex,
	visited map[Vertex]struct{},
	ret *[]DfsRet[Vertex],
	time *int,
	visitFunc func(Vertex) bool,
) {
	visited[u] = struct{}{}

	*time += 1
	node := DfsRet[Vertex]{
		V:        u,
		findTime: *time,
	}

	// do something
	if visitFunc != nil {
		visitFunc(u)
	}

	for _, e := range g.data[u] {
		if _, ok := visited[e.V]; !ok {
			g.dfsVisit(e.V, visited, ret, time, visitFunc)
		} else {
			g.HasCycle = true
		}
	}

	*time += 1
	node.finishTime = *time
	*ret = append(*ret, node)
}

func (g GraphAdj[Vertex, Weight]) DepthFirstForest() [][]DfsRet[Vertex] {
	dfsret := g.Dfs(nil)

	sort.Slice(dfsret, func(i, j int) bool {
		return dfsret[i].findTime < dfsret[j].findTime
	})

	ret := [][]DfsRet[Vertex]{}

	tmp := []DfsRet[Vertex]{}
	for _, v := range dfsret {

		tmp = append(tmp, v)
		if v.finishTime == v.findTime+1 {
			ret = append(ret, tmp)
			tmp = []DfsRet[Vertex]{}
		}
	}
	return ret
}

func (g GraphAdj[Vertex, Weight]) StrongConnectedComponents() [][]Vertex {

	dfsRet := g.Dfs(nil)

	sort.Slice(dfsRet, func(i, j int) bool {
		return dfsRet[i].finishTime > dfsRet[j].finishTime
	})

	// fmt.Println(dfsRet)

	gt := g.Transpose()

	visited := make(map[Vertex]struct{})
	gtDfsret := []DfsRet[Vertex]{}
	time := 0

	for _, ret := range dfsRet {
		if _, ok := visited[ret.V]; !ok {
			gt.dfsVisit(ret.V, visited, &gtDfsret, &time, nil)
		}
	}

	sort.Slice(gtDfsret, func(i, j int) bool {
		return gtDfsret[i].findTime < gtDfsret[j].findTime
	})

	// fmt.Println(gtDfsret)

	ret := [][]Vertex{}

	tmp := []Vertex{}
	for _, v := range gtDfsret {

		tmp = append(tmp, v.V)
		if v.finishTime == v.findTime+1 {
			ret = append(ret, tmp)
			tmp = []Vertex{}
		}
	}
	return ret
}
