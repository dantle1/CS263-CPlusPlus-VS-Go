package main

import (
    "container/heap"
    "fmt"
)

const INF = int(^uint(0) >> 1) // Max int

type Graph struct {
    V   int
    Adj map[int][]Edge
}

type Edge struct {
    To, Weight int
}

func NewGraph(V int) *Graph {
    return &Graph{
        V:   V,
        Adj: make(map[int][]Edge),
    }
}

func (g *Graph) AddEdge(u, v, w int) {
    g.Adj[u] = append(g.Adj[u], Edge{v, w})
    g.Adj[v] = append(g.Adj[v], Edge{u, w})
}

func (g *Graph) ShortestPath(src int) []int {
    dist := make([]int, g.V)
    for i := range dist {
        dist[i] = INF
    }
    dist[src] = 0

    pq := make(PriorityQueue, 0)
    heap.Init(&pq)
    heap.Push(&pq, &Item{value: src, priority: 0})

    for pq.Len() > 0 {
        u := heap.Pop(&pq).(*Item).value
        for _, e := range g.Adj[u] {
            if dist[e.To] > dist[u]+e.Weight {
                dist[e.To] = dist[u] + e.Weight
                heap.Push(&pq, &Item{value: e.To, priority: dist[e.To]})
            }
        }
    }
    return dist
}

type Item struct {
    value    int // The value of the item; arbitrary.
    priority int // The priority of the item in the queue.
    // The index is needed by update and is maintained by the heap.Interface methods.
    index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
    // We want Pop to give us the highest, not lowest, priority so we use greater than here.
    return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
    pq[i].index = i
    pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
    n := len(*pq)
    item := x.(*Item)
    item.index = n
    *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    item.index = -1 // for safety
    *pq = old[0 : n-1]
    return item
}

func main() {
    V := 9
    g := NewGraph(V)

    // making above shown graph
    g.AddEdge(0, 1, 4)
    g.AddEdge(0, 7, 8)
    g.AddEdge(1, 2, 8)
    g.AddEdge(1, 7, 11)
    g.AddEdge(2, 3, 7)
    g.AddEdge(2, 8, 2)
    g.AddEdge(2, 5, 4)
    g.AddEdge(3, 4, 9)
    g.AddEdge(3, 5, 14)
    g.AddEdge(4, 5, 10)
    g.AddEdge(5, 6, 2)
    g.AddEdge(6, 7, 1)
    g.AddEdge(6, 8, 6)
    g.AddEdge(7, 8, 7)

    // Function call
    dist := g.ShortestPath(0)
    
    fmt.Println("Vertex Distance from Source:")
    for i, d := range dist {
        fmt.Printf("%d \t\t %d\n", i, d)
    }
}