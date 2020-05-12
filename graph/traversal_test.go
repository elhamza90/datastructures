package graph

import "testing"
func TestBreadthFirstSearch(t *testing.T) {
	g := New(10);
	g.AddEdge(3, 9);
	g.AddEdge(3, 7);
	g.AddEdge(7, 8);
	g.AddEdge(0, 3);
	g.AddEdge(0, 1);
	g.AddEdge(1, 6);
	g.AddEdge(1, 5);
	g.AddEdge(1, 2);
	res := g.BreadthFirstSearch(0);
	if res[0] != 0 || res[1] != 3 || res[2] != 1 || res[3] != 9 || res[4] != 7 {
		t.Error("Error BFS Graph traversal")
	}
	//fmt.Println("BFS", res);
}

func TestDepthFirstSearch(t *testing.T) {
	g := New(10);
	g.AddEdge(3, 9);
	g.AddEdge(3, 7);
	g.AddEdge(7, 8);
	g.AddEdge(0, 3);
	g.AddEdge(0, 1);
	g.AddEdge(1, 6);
	g.AddEdge(1, 5);
	g.AddEdge(1, 2);
	res := g.DepthFirstSearch(0);
	if res[0] != 0 || res[1] != 1 || res[2] != 2 || res[3] != 5 || res[4] != 6 || res[5] != 3 || res[6] != 7 || res[7] != 8 {
		t.Error("Error DFS Graph traversal")
	}
	//fmt.Println("DFS", res);
}
