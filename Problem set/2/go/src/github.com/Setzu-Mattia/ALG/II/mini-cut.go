package main

import "fmt"
import "math/rand"

var (
	g     *Graph
	nodes []Node
	edges []Edge
)

const nodesCount int = 10
const maxEdgesPerNode int = 7

func setUp() {
	g = EmptyGraph()

	for i := 0; i < nodesCount; i++ {
		node := NewNode(i)
		nodes = append(nodes, *node)

		g.addNode(*node)
	}

	for i := 0; i < nodesCount-1; i++ {
		// Random number of edges per node, ranging
		// from 1 to maxEdgesPerNode + 1
		edgesCount := rand.Intn(maxEdgesPerNode) + 1
		edge := NewEdge(nodes[i], nodes[i+1])

		for j := 0; j < edgesCount; j++ {
			edges = append(edges, *edge)
			g.addEdge(*edge)
		}

	}
}

func (h *Graph) printGraph() {
	nCount := len((*h).nodes)
	eCount := len((*h).edges)

	fmt.Println("*------------------*")
	for i := 0; i < nCount; i++ {
		fmt.Println("Node:", (*h).nodes[i].val)
	}
	for i := 0; i < eCount; i++ {
		fmt.Println("Edge:", (*h).edges[i].m.val, (*h).edges[i].n.val)
	}
	fmt.Println("*------------------*")
}

func main() {
	setUp()

	nodes := []Node{}
	for i := 0; i < nodesCount; i++ {
		nodes = append(nodes, (*g).nodes[i])
	}

	g.printGraph()
}
