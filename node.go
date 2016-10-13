package main

type Node struct {
	n_id  int
	x     float64
	y     float64
	edges map[int]*Edge
}

func NewNode(n_id int, x float64, y float64) Node {
	return Node{n_id: n_id, x: x, y: y, edges: map[int]*Edge{}}
}

func (n *Node) addEdge(dst int, edge *Edge) {
	n.edges[dst] = edge
}
