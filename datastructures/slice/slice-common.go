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
