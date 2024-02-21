package main

import (
	"fmt"
	"sort"
)

// DSU data structure
type DSU struct {
	parent []int
	rank   []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = -1
		rank[i] = 1
	}
	return &DSU{parent, rank}
}

// Find function with path compression
func (d *DSU) find(i int) int {
	if d.parent[i] == -1 {
		return i
	}
	d.parent[i] = d.find(d.parent[i])
	return d.parent[i]
}

// Union function with rank by union
func (d *DSU) unite(x, y int) {
	s1 := d.find(x)
	s2 := d.find(y)

	if s1 != s2 {
		if d.rank[s1] < d.rank[s2] {
			d.parent[s1] = s2
		} else if d.rank[s1] > d.rank[s2] {
			d.parent[s2] = s1
		} else {
			d.parent[s2] = s1
			d.rank[s1]++
		}
	}
}

// Edge struct to represent graph edges
type Edge struct {
	x, y, w int // from, to, weight
}

type Graph struct {
	edgeList []Edge
	V        int
}

func NewGraph(V int) *Graph {
	return &Graph{
		edgeList: make([]Edge, 0),
		V:        V,
	}
}

func (g *Graph) addEdge(x, y, w int) {
	g.edgeList = append(g.edgeList, Edge{x, y, w})
}

func (g *Graph) kruskalsMST() {
	sort.Slice(g.edgeList, func(i, j int) bool {
		return g.edgeList[i].w < g.edgeList[j].w
	})

	ds := NewDSU(g.V)
	fmt.Println("Following are the edges in the constructed MST:")
	var ans int
	for _, edge := range g.edgeList {
		w, x, y := edge.w, edge.x, edge.y
		if ds.find(x) != ds.find(y) {
			ds.unite(x, y)
			ans += w
			fmt.Printf("%d -- %d == %d\n", x, y, w)
		}
	}
	fmt.Println("Minimum Cost Spanning Tree:", ans)
}

func main() {
	g := NewGraph(4)
	g.addEdge(0, 1, 10)
	g.addEdge(1, 3, 15)
	g.addEdge(2, 3, 4)
	g.addEdge(2, 0, 6)
	g.addEdge(0, 3, 5)

	g.kruskalsMST()
}
