package main

import "fmt"

type GraphMatrix struct {
	adjMatrix [][]bool
}

// NewGraphMatrix initializes a graph with the given number of vertices
func NewGraphMatrix(numVertices int) *GraphMatrix {
	matrix := make([][]bool, numVertices)
	for i := range matrix {
		matrix[i] = make([]bool, numVertices)
	}
	return &GraphMatrix{adjMatrix: matrix}
}

// AddEdge adds an undirected edge between u and v
func (g *GraphMatrix) AddEdge(u, v int) {
	if u < 0 || v < 0 || u >= len(g.adjMatrix) || v >= len(g.adjMatrix) {
		return // invalid indices
	}
	g.adjMatrix[u][v] = true
	g.adjMatrix[v][u] = true
}

// EdgeExists checks whether an edge exists between u and v
func (g *GraphMatrix) EdgeExists(u, v int) bool {
	if u < 0 || v < 0 || u >= len(g.adjMatrix) || v >= len(g.adjMatrix[u]) {
		return false
	}
	return g.adjMatrix[u][v]
}

// Example usage
func main() {
	g := NewGraphMatrix(5)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)

	fmt.Println(g.EdgeExists(0, 1)) // true
	fmt.Println(g.EdgeExists(1, 3)) // false
}
