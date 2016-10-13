package main

import (
	"math/rand"
)

type Agent struct {
	paths           []int
	selectedWeights []float64
}

func NewAgent() *Agent {
	return &Agent{paths: []int{1}, selectedWeights: []float64{}}
}

func (a *Agent) now() int {
	return a.paths[len(a.paths)-1]
}

func (a *Agent) addPath(n int) {
	a.paths = append(a.paths, n)
}

func (a *Agent) nextEdge(g *Graph, prob map[int]float64) int {
	r := rand.Float64()

	sum := 0.0

	for i, p := range prob {

		sum += p

		if r < sum {
			nextEdge := g.nodes[a.now()].edges[i]
			a.selectedWeights = append(a.selectedWeights, nextEdge.weight)
			return nextEdge.dst
		}

	}
	return 0
}

func (a *Agent) totalWeight() float64 {
	w := 0.0

	for i := 0; i < len(a.selectedWeights); i++ {
		w += a.selectedWeights[i]
	}

	return w
}

func (a *Agent) done(g *Graph) bool {
	if len(g.nodes) == len(a.paths) {
		return true
	} else {
		return false
	}
}
