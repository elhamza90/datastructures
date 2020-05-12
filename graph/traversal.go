package graph

// Use our own implemented queue
import (
	"datastructures/queue"
	"datastructures/stack"
);

func (g graph) BreadthFirstSearch(rootVertex int) (path []int){

	// 1. Create and Add root to queue
	q := queue.New(rootVertex);
	path = make([]int, 0, g.NbrOfVertices()); // result of traversal.
	path = append(path, rootVertex);
	visited := map[int]bool{rootVertex: true};

	// 2. For each element in queue
	for q.Length() != 0 {
		v, _ := q.Pop();
		for _, adj := range g[v] {
			if visited[adj] == false {
				q.Push(adj);
				path = append(path, adj);
				visited[adj] = true;
			}
		}
	}

	return path;
}


func (g graph) DepthFirstSearch(rootVertex int) (path []int) {

	// 1. Create and add root to stack
	s := stack.New(rootVertex);
	path = make([]int, 0, g.NbrOfVertices()); // result of traversal.
	path = append(path, rootVertex);
	visited := map[int]bool{rootVertex: true};

	// 2. Loop
	for s.Length() != 0 {
		v, _ := s.Pop();
		if visited[v] == false {
			path = append(path, v);
			visited[v] = true;
		}
		for _, adj := range g[v] {
			if visited[adj] == false {
				s.Push(adj);
			}
		}
	}
	return path;
}