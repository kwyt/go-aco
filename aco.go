package main

import (
	"fmt"
)

type Aco struct {
	alpha           float64
	beta            float64
	evaporationRate float64
	it              float64
	q               float64
	res             []int
	cost            float64
}

func NewAco() *Aco {
	return &Aco{
		alpha:           1,
		beta:            1,
		evaporationRate: 0.4,
		it:              0,
		q:               1,
		cost:            0,
	}
}

func (aco *Aco) run(graph *Graph, tryCount int, agentNum int, export bool) {

	pl := &Plot{}

	for i := 0; i < tryCount; i++ {

		an := agentNum

		for an > 0 {
			if an == 0 {
				break
			}

			agent := NewAgent()
			graph.resetEdgeEnabled()
			graph.resetPheromone()

			for {
				graph.capableEdges(agent)
				prob := graph.probability(agent, aco.alpha, aco.beta)
				next := agent.nextEdge(graph, prob)
				agent.addPath(next)

				if agent.done(graph) {
					break
				}
			}

			graph.evaporate(aco.evaporationRate)
			graph.updatePheromones(aco.q, agent)
			tw := agent.totalWeight()

			if aco.cost == 0 {
				aco.cost = tw
				aco.res = agent.paths
			} else {
				if aco.cost > tw {
					aco.cost = tw
					aco.res = agent.paths
				}
			}

			an--
		}

	}

	fmt.Println("Route: ", aco.res)
	fmt.Println("Optimization: ", aco.cost)

	if export {
		pl.initialize()
		pl.drawNode(graph.nodes, aco.res)
		pl.export("result")
	}
}

func (aco *Aco) result() ([]int, float64) {
	return aco.res, aco.cost
}
