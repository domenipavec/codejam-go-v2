// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package slice

// NewSliceInt creates slice length n
func NewSliceInt(n int) SliceInt {
	return make([]int, n)
}

// Copy makes a new independent copy of slice
func (slice SliceInt) Copy() SliceInt {
	newSlice := NewSliceInt(len(slice))
	copy(newSlice, slice)
	return newSlice
}

func (slice SliceInt) Less(i, j int) bool {
	return slice[i] < slice[j]
}

// NewSliceFloat64 creates slice length n
func NewSliceFloat64(n int) SliceFloat64 {
	return make([]float64, n)
}

// Copy makes a new independent copy of slice
func (slice SliceFloat64) Copy() SliceFloat64 {
	newSlice := NewSliceFloat64(len(slice))
	copy(newSlice, slice)
	return newSlice
}

func (slice SliceFloat64) Less(i, j int) bool {
	return slice[i] < slice[j]
}

// NewSliceString creates slice length n
func NewSliceString(n int) SliceString {
	return make([]string, n)
}

// Copy makes a new independent copy of slice
func (slice SliceString) Copy() SliceString {
	newSlice := NewSliceString(len(slice))
	copy(newSlice, slice)
	return newSlice
}

func (slice SliceString) Less(i, j int) bool {
	return slice[i] < slice[j]
}

// NewSliceByte creates slice length n
func NewSliceByte(n int) SliceByte {
	return make([]byte, n)
}

// Copy makes a new independent copy of slice
func (slice SliceByte) Copy() SliceByte {
	newSlice := NewSliceByte(len(slice))
	copy(newSlice, slice)
	return newSlice
}

func (slice SliceByte) Less(i, j int) bool {
	return slice[i] < slice[j]
}
