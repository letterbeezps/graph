package graph

import (
	"sort"
)

func (g GraphAdj[Vertex, Weight]) TopologicalSort() []DfsRet[Vertex] {
	dfsret := g.Dfs(nil)

	if g.HasCycle {
		g.Logger.Fatal("Has cycle : only DAG can call method TopologicalSort")
	}

	sort.Slice(dfsret, func(i, j int) bool {
		return dfsret[i].finishTime > dfsret[j].finishTime
	})

	return dfsret
}
