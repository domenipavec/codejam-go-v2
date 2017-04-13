package slice

import (
	"container/heap"
	"fmt"
	"sort"
)

//go:generate genny -in=$GOFILE -out=../gen-$GOFILE gen "ValueType=int,SliceInt,float64,SliceFloat64,string,SliceString,byte,SliceByte"

// SliceValueType type
type SliceValueType []ValueType

// Set sets all element to c
func (slice SliceValueType) Set(c ValueType) {
	for i := 0; i < len(slice); i++ {
		slice[i] = c
	}
}

// SortAsc sort ascending
func (slice SliceValueType) SortAsc() {
	sort.Sort(slice)
}

// SortDesc sort descending
func (slice SliceValueType) SortDesc() {
	sort.Sort(sort.Reverse(slice))
}

// Len length
func (slice SliceValueType) Len() int {
	return len(slice)
}

// Get i-th element
func (slice SliceValueType) Get(i int) ValueType {
	if i < 0 {
		i += len(slice)
	}

	return slice[i]
}

// Swap two elements
func (slice SliceValueType) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// Push element
func (slice *SliceValueType) Push(c interface{}) {
	*slice = append(*slice, c.(ValueType))
}

// Pop element
func (slice *SliceValueType) Pop() interface{} {
	element := slice.Get(-1)
	*slice = (*slice)[:len(*slice)-1]
	return element
}

// Append values
func (slice *SliceValueType) Append(values ...ValueType) {
	*slice = append(*slice, values...)
}

// Prepend values
func (slice *SliceValueType) Prepend(values ...ValueType) {
	*slice = append(values, (*slice)...)
}

// Print prints using separator
func (slice SliceValueType) Print(sep string) string {
	output := ""
	for _, c := range slice {
		if output != "" {
			output += sep
		}
		output += fmt.Sprintf("%v", c)
	}
	return output
}

// MinHeapSliceValueType struct for min heap
type MinHeapSliceValueType struct {
	Slice *SliceValueType
}

// MinHeap returns struct with min heap functionality based on SliceValueType
func (slice *SliceValueType) MinHeap() MinHeapSliceValueType {
	hp := MinHeapSliceValueType{
		Slice: slice,
	}
	heap.Init(hp.Slice)
	return hp
}

// Min returns min element
func (hp MinHeapSliceValueType) Min() ValueType {
	return (*hp.Slice)[0]
}

// Fix re-establishes heap ordering after i has changed value
func (hp MinHeapSliceValueType) Fix(i int) {
	heap.Fix(hp.Slice, i)
}

// Pop removes the minimum element
func (hp MinHeapSliceValueType) Pop() ValueType {
	return heap.Pop(hp.Slice).(ValueType)
}

// Push v to heap
func (hp MinHeapSliceValueType) Push(v ValueType) {
	heap.Push(hp.Slice, v)
}

// Remove i-th element
func (hp MinHeapSliceValueType) Remove(i int) ValueType {
	return heap.Remove(hp.Slice, i).(ValueType)
}

// MaxHeapSliceValueType struct for max heap
type MaxHeapSliceValueType struct {
	Slice *SliceValueType
}

// MaxHeap returns struct with max heap functionality based on SliceValueType
func (slice *SliceValueType) MaxHeap() MaxHeapSliceValueType {
	hp := MaxHeapSliceValueType{
		Slice: slice,
	}
	heap.Init(HeapReverse(hp.Slice))
	return hp
}

// Max returns max element
func (hp MaxHeapSliceValueType) Max() ValueType {
	return (*hp.Slice)[0]
}

// Fix re-establishes heap ordering after i has changed value
func (hp MaxHeapSliceValueType) Fix(i int) {
	heap.Fix(HeapReverse(hp.Slice), i)
}

// Pop removes the minimum element
func (hp MaxHeapSliceValueType) Pop() ValueType {
	return heap.Pop(HeapReverse(hp.Slice)).(ValueType)
}

// Push v to heap
func (hp MaxHeapSliceValueType) Push(v ValueType) {
	heap.Push(HeapReverse(hp.Slice), v)
}

// Remove i-th element
func (hp MaxHeapSliceValueType) Remove(i int) ValueType {
	return heap.Remove(HeapReverse(hp.Slice), i).(ValueType)
}
