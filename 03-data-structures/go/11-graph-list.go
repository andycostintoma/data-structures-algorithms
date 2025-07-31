package main

import (
	"fmt"
	"slices"
)

type set map[int]struct{}

func newSet() set {
	return make(set)
}

func (s set) add(val int) {
	s[val] = struct{}{}
}

func (s set) contains(val int) bool {
	_, exists := s[val]
	return exists
}

func (s set) remove(val int) {
	delete(s, val)
}

func (s set) toSlice() []int {
	out := make([]int, 0, len(s))
	for v := range s {
		out = append(out, v)
	}
	return out
}

type Graph struct {
	adj map[int]set
}

func NewGraph() *Graph {
	return &Graph{adj: make(map[int]set)}
}

func (g *Graph) AddNode(u int) {
	if _, exists := g.adj[u]; !exists {
		g.adj[u] = newSet()
	}
}

func (g *Graph) AddEdge(u, v int) {
	if _, exists := g.adj[u]; !exists {
		g.adj[u] = newSet()
	}
	if _, exists := g.adj[v]; !exists {
		g.adj[v] = newSet()
	}
	g.adj[u].add(v)
	g.adj[v].add(u)
}

func (g *Graph) EdgeExists(u, v int) bool {
	neighborsU, okU := g.adj[u]
	neighborsV, okV := g.adj[v]
	if !okU || !okV {
		return false
	}
	return neighborsU.contains(v) && neighborsV.contains(u)
}

func (g *Graph) AdjacentNodes(node int) []int {
	if neighbors, ok := g.adj[node]; ok {
		return neighbors.toSlice()
	}
	return nil
}

func (g *Graph) UnconnectedVertices() []int {
	var unconnected []int
	for node, neighbors := range g.adj {
		if len(neighbors) == 0 {
			unconnected = append(unconnected, node)
		}
	}
	return unconnected
}

func (g *Graph) Bfs(start int) []int {
	visited := newSet()
	queue := []int{start}
	result := []int{}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if visited.contains(curr) {
			continue
		}
		visited.add(curr)
		result = append(result, curr)

		for neightbour := range g.adj[curr] {
			if !visited.contains(neightbour) {
				queue = append(queue, neightbour)
			}
		}
	}
	return result
}

func (g *Graph) String() string {
	var result string
	for node, neighbors := range g.adj {
		adjList := neighbors.toSlice()
		slices.Sort(adjList)
		result += fmt.Sprintf("%d: %v\n", node, adjList)
	}
	return result
}

func main() {
	g := NewGraph()

	g.AddNode(1)
	g.AddNode(2)
	g.AddNode(3)
	g.AddNode(99) // unconnected node

	g.AddEdge(1, 2)
	g.AddEdge(2, 3)

	fmt.Println("Edge 1-2 exists?", g.EdgeExists(1, 2)) // true
	fmt.Println("Edge 1-3 exists?", g.EdgeExists(1, 3)) // false

	fmt.Println("Adjacent to 2:", g.AdjacentNodes(2))    // [1 3]
	fmt.Println("Unconnected:", g.UnconnectedVertices()) // [99]

	fmt.Println("Graph: ", g)
	fmt.Println("BFS: ", g.Bfs(1))
}
