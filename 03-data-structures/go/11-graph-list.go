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
	queue := NewQueue[int]()
	queue.Enqueue(start)
	result := []int{}

	for queue.Size() > 0 {
		curr, _ := queue.Dequeue()

		if visited.contains(curr) {
			continue
		}
		visited.add(curr)
		result = append(result, curr)

		if neighbors, ok := g.adj[curr]; ok {
			for neighbor := range neighbors {
				if !visited.contains(neighbor) {
					queue.Enqueue(neighbor)
				}
			}
		}
	}
	return result
}

func (g *Graph) Dfs(start int) []int {
	visited := newSet()
	stack := NewStack[int]()
	stack.Push(start)
	result := []int{}

	for stack.Size() > 0 {
		curr, _ := stack.Pop()

		if visited.contains(curr) {
			continue
		}
		visited.add(curr)
		result = append(result, curr)

		if neighbors, ok := g.adj[curr]; ok {
			for neighbor := range neighbors {
				if !visited.contains(neighbor) {
					stack.Push(neighbor)
				}
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
