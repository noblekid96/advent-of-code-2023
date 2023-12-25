package main

import (
	"fmt"
	"math"
)

type Edge struct {
    to, weight int
}

type Graph struct {
    edges [][]Edge
}

func NewGraph(size int) *Graph {
    return &Graph{
        edges: make([][]Edge, size),
    }
}

func (g *Graph) AddEdge(from, to, weight int) {
    g.edges[from] = append(g.edges[from], Edge{to, weight})
    // For undirected graph, add edge in both directions
    g.edges[to] = append(g.edges[to], Edge{from, weight})
}

func (g *Graph) MinCut() int {
    n := len(g.edges)
    bestCut := math.MaxInt32

    for phase := 0; phase < n-1; phase++ {
        // Initialize arrays for the algorithm
        inA := make([]bool, n)
        w := make([]int, n)
        prev := -1

        for i := 0; i < n; i++ {
            maxW, maxV := -1, -1

            for j := 0; j < n; j++ {
                if !inA[j] && w[j] > maxW {
                    maxW = w[j]
                    maxV = j
                }
            }

            if maxV == -1 {
                break
            }

            // Add maxV to A
            inA[maxV] = true
            prev = maxV

            // Update the weights
            for _, e := range g.edges[maxV] {
                w[e.to] += e.weight
            }
        }

        // Update the best cut found so far
        if prev != -1 && w[prev] < bestCut {
            bestCut = w[prev]
        }

        // Contract the graph
        // Merge the two vertices that were added last
        // Update graph (not shown here for brevity)
    }

    return bestCut
}

func main() {
    graph := NewGraph(5)
    // Add edges to the graph
    graph.AddEdge(0, 1, 2)
    graph.AddEdge(1, 2, 3)
    // ... add other edges ...

    minCut := graph.MinCut()
    fmt.Println("Minimum Cut:", minCut)
}
