package graph

import "testing"
import "fmt"
func TestBreadthFirstSearch(t *testing.T) {
	g := New(7);
	g.AddEdge(0, 4);
	g.AddEdge(0, 3);
	g.AddEdge(0, 1);
	g.AddEdge(1, 6);
	g.AddEdge(1, 5);
	g.AddEdge(1, 2);
	res := g.BreadthFirstSearch(0);
	fmt.Println(res);
}

