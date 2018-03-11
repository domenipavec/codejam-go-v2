package graph

import (
	"errors"
	"fmt"

	"github.com/matematik7/codejam-go-v2/datastructures/queue"
	"github.com/matematik7/codejam-go-v2/datastructures/slice"
)

type Graph struct {
	N        int
	OutEdges [][]int
	InEdges  [][]int
}

func New(N int) Graph {
	return Graph{
		N:        N,
		OutEdges: make([][]int, N),
		InEdges:  make([][]int, N),
	}
}

func (g Graph) String() string {
	return g.StringOutEdges()
}

func (g Graph) StringOutEdges() string {
	output := ""
	for i := range g.OutEdges {
		if i != 0 {
			output += "\n"
		}
		output += fmt.Sprintf("%d:", i)
		for _, neighbor := range g.OutEdges[i] {
			output += fmt.Sprintf(" %d", neighbor)
		}
	}
	return output
}

func (g Graph) StringInEdges() string {
	output := ""
	for i := range g.InEdges {
		if i != 0 {
			output += "\n"
		}
		output += fmt.Sprintf("%d:", i)
		for _, neighbor := range g.InEdges[i] {
			output += fmt.Sprintf(" %d", neighbor)
		}
	}
	return output
}

func (g Graph) AddEdge(u, v int) {
	g.OutEdges[u] = append(g.OutEdges[u], v)
	g.InEdges[v] = append(g.InEdges[v], u)
}

func (g Graph) AddBiEdge(u, v int) {
	g.AddEdge(u, v)
	g.AddEdge(v, u)
}

func (g Graph) TopologicalSort() ([]int, error) {
	order := make([]int, 0, g.N)
	q := queue.NewInt()
	visited := 0

	inDegrees := make([]int, g.N)
	for i := range inDegrees {
		inDegrees[i] = len(g.InEdges[i])
		if inDegrees[i] == 0 {
			q.Push(i)
		}
	}

	for q.Len() > 0 {
		v := q.Pop()
		visited += 1
		order = append(order, v)

		for _, neighbor := range g.OutEdges[v] {
			inDegrees[neighbor]--
			if inDegrees[neighbor] == 0 {
				q.Push(neighbor)
			}
		}
	}

	if visited != g.N {
		return nil, errors.New("Topological sort not possible")
	}

	return order, nil
}

func (g Graph) MinTopologicalSort() ([]int, error) {
	order := make([]int, 0, g.N)
	s := slice.NewSliceInt(0)
	q := s.MinHeap()
	visited := 0

	inDegrees := make([]int, g.N)
	for i := range inDegrees {
		inDegrees[i] = len(g.InEdges[i])
		if inDegrees[i] == 0 {
			q.Push(i)
		}
	}

	for q.Slice.Len() > 0 {
		v := q.Pop()
		visited += 1
		order = append(order, v)

		for _, neighbor := range g.OutEdges[v] {
			inDegrees[neighbor]--
			if inDegrees[neighbor] == 0 {
				q.Push(neighbor)
			}
		}
	}

	if visited != g.N {
		return nil, errors.New("Topological sort not possible")
	}

	return order, nil
}
