package graph

import "strconv"

type graph [][]int;

func New(nbrVertices int) graph {
	return make(graph, nbrVertices, nbrVertices);
}

// AddEdge adds an edge between two given vertices.
func (g graph) AddEdge(v1 int, v2 int) {
	if v1 < len(g) && v2 < len(g) {
		g[v1] = append(g[v1], v2);
		g[v2] = append(g[v2], v1);
	}
}

func (g graph) NbrOfVertices() int {
	return len(g);
}

func (g graph) String() (s string) {
	s = "";
	for i, list := range g {
		s += strconv.Itoa(i) + " -> ";
		if len(list) == 0 {
			s += "[]\n";
		} else {
			s += "[ ";
			for _, v := range list {
				s += strconv.Itoa(v) + " ";
			}
			s += "]\n";
		}
	}
	return s;
}