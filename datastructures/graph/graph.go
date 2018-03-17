package graph

import (
	"errors"
	"fmt"

	"github.com/matematik7/codejam-go-v2/datastructures/intset"
	"github.com/matematik7/codejam-go-v2/datastructures/queue"
	"github.com/matematik7/codejam-go-v2/datastructures/slice"
	"github.com/matematik7/codejam-go-v2/integer"
)

type Graph struct {
	N        int
	OutEdges [][]Edge
	InEdges  [][]Edge
}

type Edge struct {
	From   int
	To     int
	Weight int
}

func New(N int) Graph {
	return Graph{
		N:        N,
		OutEdges: make([][]Edge, N),
		InEdges:  make([][]Edge, N),
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
		for _, edge := range g.OutEdges[i] {
			output += fmt.Sprintf(" %d", edge.To)
			if edge.Weight != 0 {
				output += fmt.Sprintf("(%d)", edge.Weight)
			}
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
		for _, edge := range g.InEdges[i] {
			output += fmt.Sprintf(" %d", edge.From)
		}
	}
	return output
}

func (g Graph) AddEdge(u, v int, w ...int) {
	weight := 0
	if len(w) > 0 {
		weight = w[0]
	}
	edge := Edge{
		From:   u,
		To:     v,
		Weight: weight,
	}
	g.OutEdges[u] = append(g.OutEdges[u], edge)
	g.InEdges[v] = append(g.InEdges[v], edge)
}

func (g Graph) AddBiEdge(u, v int, w ...int) {
	g.AddEdge(u, v, w...)
	g.AddEdge(v, u, w...)
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

		for _, edge := range g.OutEdges[v] {
			inDegrees[edge.To]--
			if inDegrees[edge.To] == 0 {
				q.Push(edge.To)
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

		for _, edge := range g.OutEdges[v] {
			inDegrees[edge.To]--
			if inDegrees[edge.To] == 0 {
				q.Push(edge.To)
			}
		}
	}

	if visited != g.N {
		return nil, errors.New("Topological sort not possible")
	}

	return order, nil
}

// DjikstraArray uses array for queue and is O(N^2)
func (g Graph) DjikstraArray(source int) slice.SliceInt {
	visited := intset.New(g.N)

	distances := slice.NewSliceInt(g.N)
	for v := range distances {
		if v == source {
			continue
		}
		distances[v] = integer.MAX
	}

	current := source
	for {
		for _, edge := range g.OutEdges[current] {
			if visited.Contains(edge.To) {
				continue
			}
			distance := distances[current] + edge.Weight
			if distances[edge.To] > distance {
				distances[edge.To] = distance
			}
		}
		visited.Add(current)

		imin := -1
		for i := range distances {
			if visited.Contains(i) {
				continue
			}
			if imin == -1 || distances[i] < distances[imin] {
				imin = i
			}
		}

		if imin == -1 || distances[imin] == integer.MAX {
			break
		}

		current = imin
	}

	return distances
}
