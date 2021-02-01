package Graph

import (
	"testing"
)

func TestGraph_BFS(t *testing.T) {
	graph := newGraph(7)
	graph.addEdge(1, 2)
	graph.addEdge(1, 3)
	graph.addEdge(2, 4)
	graph.addEdge(2, 5)
	graph.addEdge(4, 5)
	graph.addEdge(4, 6)
	graph.addEdge(5, 3)
	graph.addEdge(5, 6)
	graph.addEdge(3, 6)

	res := graph.BFSTraverse()
	t.Log(res)

}

func TestGraph_BFSMinDistance(t *testing.T) {
	graph := newGraph(7)
	graph.addEdge(1, 2)
	graph.addEdge(1, 3)
	graph.addEdge(2, 4)
	graph.addEdge(2, 5)
	graph.addEdge(4, 5)
	graph.addEdge(4, 6)
	graph.addEdge(5, 3)
	graph.addEdge(5, 6)
	graph.addEdge(3, 6)

	d := graph.BFSMinDistance(1)
	t.Log(d)
}

func TestGraph_DFSTraverse(t *testing.T) {
	graph := newGraph(7)
	graph.addEdge(1, 2)
	graph.addEdge(1, 3)
	graph.addEdge(2, 4)
	graph.addEdge(2, 5)
	graph.addEdge(4, 5)
	graph.addEdge(4, 6)
	graph.addEdge(5, 3)
	graph.addEdge(5, 6)
	graph.addEdge(3, 6)

	t.Log(graph.DFSTraverse())
}
