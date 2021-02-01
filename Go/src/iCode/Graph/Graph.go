package Graph

import (
	"container/list"
)

/**
@Description: 无向图，使用邻接表存储
@Date: 2/1/2021
@Author: lichang
*/

type Graph struct {
	adj []*list.List
	v   int // 节点数目
}

// 初始化图
func newGraph(v int) *Graph {
	graph := &Graph{
		adj: make([]*list.List, v),
		v:   v,
	}
	for i := range graph.adj {
		graph.adj[i] = list.New()
	}
	return graph
}

// 插入一条边，一条边存两次
func (g *Graph) addEdge(s int, t int) {
	g.adj[s].PushBack(t)
	g.adj[t].PushBack(s)
}

// 广度优先遍历
func (g *Graph) BFSTraverse() []int {
	visited := make([]bool, g.v+1) // 访问标志
	var res []int
	for i := 0; i < g.v; i++ {
		if !visited[i] { // 对每个连通分量做一次BFS
			r := g.BFS(i, visited) // 从第i个节点出发
			res = append(res, r...)
		}
	}
	return res
}

func (g *Graph) BFS(i int, visited []bool) []int {
	var res []int      // 遍历结果
	var queue intQueue // 辅助队列

	queue.push(i)     // 入队
	visited[i] = true // 置访问位

	for !queue.isEmpty() {
		cur := queue.pop() // pop
		l := g.adj[cur]
		res = append(res, cur)

		// 访问相邻节点
		p := l.Front()
		for p != nil {
			if va, ok := p.Value.(int); ok && visited[va] != true {
				queue.push(va)
				visited[va] = true
			}
			p = p.Next()
		}

	}
	return res
}

// BFS求解单源最短路径
func (g *Graph) BFSMinDistance(u int) []int {
	var d []int = make([]int, g.v) // d[i]表示从u到i节点的最短路径
	for i := range d {
		d[i] = -1
	}

	visited := make([]bool, g.v)
	visited[u] = true
	d[u] = 0
	var queue intQueue

	queue.push(u)

	for !queue.isEmpty() {
		u := queue.pop()
		for w := g.firstNeighbor(u); w >= 0; w = g.nextNeighbor(u, w) {
			if !visited[w] {
				visited[w] = true
				d[w] = d[u] + 1
				queue.push(w)
			}
		}

	}
	return d
}

// 求第u个节点的第一个邻居节点
func (g *Graph) firstNeighbor(u int) int {
	if g.adj[u].Len() <= 0 {
		return -1
	}
	p := g.adj[u].Front()
	if v, ok := p.Value.(int); ok {
		return v
	}
	return -1
}

// 求第u个节点邻居节点中第w个节点后面的节点
func (g *Graph) nextNeighbor(u int, w int) int {
	if g.adj[u].Len() <= 1 {
		return -1
	}
	p := g.adj[u].Front()
	for p != nil {
		if v, ok := p.Value.(int); ok {
			if v == w && p.Next() != nil {
				q := p.Next()
				return q.Value.(int)
			}
		}
		p = p.Next()
	}
	return -1
}

// 深度优先遍历
func (g *Graph) DFSTraverse() []int {
	var visited []bool = make([]bool, g.v)
	var res []int
	for i := 0; i < g.v; i++ {
		if !visited[i] {
			r := g.DFS(i, visited)
			res = append(res, r...)
		}
	}
	return res
}

func (g *Graph) DFS(u int, visited []bool) []int {
	var res []int
	res = append(res, u)
	visited[u] = true

	for w := g.firstNeighbor(u); w >= 0; w = g.nextNeighbor(u, w) {
		if !visited[w] {
			r := g.DFS(w, visited)
			res = append(res, r...)
		}
	}
	return res
}
