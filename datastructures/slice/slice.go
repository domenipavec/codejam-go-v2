package slice

import (
	"fmt"
	"sort"

	"github.com/cheekybits/genny/generic"
)

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "ValueType=byte,int8,int16,int32,int64,uint8,uint16,uint32,uint64,uint,int,float64"

// ValueType genny generic type for slice
type ValueType generic.Number

// SliceValueType type
type SliceValueType []ValueType

// NewValueType creates slice length n
func NewValueType(n int) SliceValueType {
	return make([]ValueType, n)
}

// Copy makes a new independent copy of slice
func (slice SliceValueType) Copy() SliceValueType {
	newSlice := NewValueType(len(slice))
	copy(newSlice, slice)
	return newSlice
}

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

func (slice SliceValueType) Len() int {
	return len(slice)
}

func (slice SliceValueType) Less(i, j int) bool {
	return slice[i] < slice[j]
}

func (slice SliceValueType) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// Append c to slice
func (slice *SliceValueType) Append(c ValueType) {
	*slice = append(*slice, c)
}

func (slice SliceValueType) String() string {
	output := ""
	for _, c := range slice {
		if output != "" {
			output += " "
		}
		output += fmt.Sprintf("%v", c)
	}
	return output
}
