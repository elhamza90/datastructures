package graph

// Use our own implemented queue
import "datastructures/queue";

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