package main

import (
	"math"
)

type Graph struct {
	nodes map[int]Node
}

func NewGraph() *Graph {
	return &Graph{nodes: map[int]Node{}}
}

func (g *Graph) initialize(coords [][]float64) {

	for i := 0; i < len(coords); i++ {
		n_ida := int(coords[i][0])
		xa := coords[i][1]
		ya := coords[i][2]

		g.addNode(n_ida, xa, ya)

		for j := 0; j < len(coords); j++ {
			n_idb := int(coords[j][0])
			xb := coords[j][1]
			yb := coords[j][2]

			g.addNode(n_idb, xb, yb)

			weight := calcWeight(xa, xb, ya, yb)
			edge := NewEdge(n_ida, n_idb, weight)

			g.addEdge(n_ida, n_idb, edge)
		}
	}
}

func (g *Graph) addNode(n_id int, x float64, y float64) {
	if !g.isExists(n_id) {
		g.nodes[n_id] = NewNode(n_id, x, y)
	}
}

func (g *Graph) addEdge(src int, dst int, edge *Edge) {
	if n, ok := g.nodes[src]; ok {
		n.addEdge(dst, edge)
	}
}

func (g *Graph) isExists(n_id int) bool {
	if _, ok := g.nodes[n_id]; ok {
		return true
	}
	return false
}

func (g *Graph) capableEdges(agent *Agent) {
	for _, edge := range g.nodes[agent.now()].edges {
		for i := 0; i < len(agent.paths); i++ {
			if agent.paths[i] == edge.dst {
				edge.enabled = false
			}
		}
	}
}

func (g *Graph) resetEdgeEnabled() {
	for _, node := range g.nodes {
		for _, edge := range node.edges {
			edge.enabled = true
		}
	}
}

func (g *Graph) resetPheromone() {
	for _, node := range g.nodes {
		for _, edge := range node.edges {
			edge.pheromone = 0.0000001
		}
	}
}

func (g *Graph) capbaleEdgesEvalutionSum(agent *Agent, alpha float64, beta float64) float64 {
	sum := 0.0
	for _, edge := range g.nodes[agent.now()].edges {
		if edge.enabled {
			sum += edge.evalution(alpha, beta)
		}
	}

	return sum
}

func (g *Graph) probability(agent *Agent, alpha float64, beta float64) map[int]float64 {
	probabilities := map[int]float64{}

	for _, edge := range g.nodes[agent.now()].edges {
		if edge.enabled {
			probabilities[edge.dst] = edge.evalution(alpha, beta) / g.capbaleEdgesEvalutionSum(agent, alpha, beta)
		}
	}
	return probabilities
}

func (g *Graph) updatePheromones(q float64, agent *Agent) {
	w := agent.totalWeight()
	val := q / w

	for i := 0; i < len(agent.paths); i++ {
		g.nodes[agent.paths[i]].edges[agent.paths[i]].pheromone += val
	}

}

func (g *Graph) evaporate(evaporationRate float64) {
	for _, node := range g.nodes {
		for _, edge := range node.edges {
			edge.pheromone *= 1 - evaporationRate
		}
	}
}

func calcWeight(xa float64, xb float64, ya float64, yb float64) float64 {
	xd := xa - xb
	yd := ya - yb

	return math.Sqrt(xd*xd + yd*yd)
}
