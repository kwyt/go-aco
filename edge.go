package main

import (
	"math"
)

type Edge struct {
	src       int
	dst       int
	weight    float64
	pheromone float64
	enabled   bool
}

func NewEdge(src int, dst int, weight float64) *Edge {
	return &Edge{src: src, dst: dst, weight: weight, pheromone: 0.0000001, enabled: true}
}

func (e *Edge) evalution(alpha float64, beta float64) float64 {
	return math.Pow(e.pheromone, alpha) * math.Pow(1.0/e.weight, beta)
}
