package graph


import "testing"

func TestNew(t *testing.T) {
	g := New(5);
	if len(g) != 5 {
		t.Error("Error New graph length");
	}
}

func TestAddEdge(t *testing.T) {
	g := make(graph, 3, 3);

	// 1. Case both vertices exist. Should work.
	g.AddEdge(0, 2);
	if g[0][0] != 2 || g[2][0] != 0 {
		t.Error("Error Adding edge to graph");
	}

	// 2. Case a vertex does not exist.
	g.AddEdge(0, 3);
	if len(g[0]) != 1 {
		t.Error("Error adding edge non existing");
	}
}

func TestString(t *testing.T) {
	g := make(graph, 3, 3);
	g.AddEdge(0, 2);
	g.AddEdge(1, 2);
	if g.String() != "0 -> [ 2 ]\n1 -> [ 2 ]\n2 -> [ 0 1 ]\n" {
		t.Error("Error Graph to String conversion")
	}
}