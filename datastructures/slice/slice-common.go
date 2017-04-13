package slice

import "container/heap"

// ValueType genny generic type for slice
type ValueType interface {
	LessThan(ValueType) bool
	GreaterThan(ValueType) bool
}

// NewValueType generic constructor
func NewValueType(dimensions ...int) ValueType { return nil }

type heapreverse struct {
	heap.Interface
}

func (r heapreverse) Less(i, j int) bool {
	return r.Interface.Less(j, i)
}

// HeapReverse same as sort.Reverse for heap
func HeapReverse(data heap.Interface) heap.Interface {
	return &heapreverse{data}
}

// Coordinate is used to describe I and J grid index
type Coordinate struct {
	I int
	J int
}

func spiralTopRight(data []Coordinate, x1, y1, x2, y2 int) []Coordinate {
	for j := x1; j <= x2; j++ {
		data = append(data, Coordinate{I: y1, J: j})
	}

	for i := y1 + 1; i <= y2; i++ {
		data = append(data, Coordinate{I: i, J: x2})
	}

	if x2-x1 > 0 {
		data = spiralBottomLeft(data, x1, y1+1, x2-1, y2)
	}

	return data
}

func spiralBottomLeft(data []Coordinate, x1, y1, x2, y2 int) []Coordinate {
	for j := x2; j >= x1; j-- {
		data = append(data, Coordinate{I: y2, J: j})
	}

	for i := y2 - 1; i >= y1; i-- {
		data = append(data, Coordinate{I: i, J: x1})
	}

	if x2-x1 > 0 {
		data = spiralTopRight(data, x1+1, y1, x2, y2-1)
	}

	return data
}
