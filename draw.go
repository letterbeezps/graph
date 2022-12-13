package graph

import (
	"log"

	"github.com/goccy/go-graphviz"
)

func (g GraphAdj[Vertex, Weight]) Draw(drawWeight bool) {
	graphv := graphviz.New()
	graph, err := graphv.Graph()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := graph.Close(); err != nil {
			log.Fatal(err)
		}
		graphv.Close()
	}()

	for k, v := range g.data {
		var err error
		u, err := VertexToString(k)
		if err != nil {
			log.Fatal(err)
		}

		node_u, err := graph.CreateNode(u)
		if err != nil {
			log.Fatal(err)
		}

		for _, e := range v {
			v, err := VertexToString(e.v)
			if err != nil {
				log.Fatal(err)
			}
			node_v, err := graph.CreateNode(v)
			if err != nil {
				log.Fatal(err)
			}
			if drawWeight {
				node_uv, err := graph.CreateEdge(u+v, node_u, node_v)
				if err != nil {
					log.Fatal(err)
				}
				label, err := VertexToString(e.w)
				if err != nil {
					log.Fatal(err)
				}
				node_uv.SetLabel(label)
			} else {
				_, err := graph.CreateEdge(u+v, node_u, node_v)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}

	graphv.SetLayout(graphviz.SFDP)

	if err := graphv.RenderFilename(graph, graphviz.PNG, "./graph.png"); err != nil {
		log.Fatal(err)
	}

}
