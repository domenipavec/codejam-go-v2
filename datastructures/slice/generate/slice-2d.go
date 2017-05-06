package slice

//go:generate genny -in=$GOFILE -out=../gen-$GOFILE gen "ValueType=SliceInt,SliceFloat64,SliceString,SliceByte,SliceBool"

// NewSliceValueType creates slice length n
func NewSliceValueType(n, m int) SliceValueType {
	newSlice := make([]ValueType, n)
	for i := range newSlice {
		newSlice[i] = NewValueType(m)
	}
	return newSlice
}

// Copy makes a new independent copy of slice
func (slice SliceValueType) Copy() SliceValueType {
	newSlice := make([]ValueType, len(slice))
	for i := range newSlice {
		newSlice[i] = slice[i].Copy()
	}
	return newSlice
}

// String is for print
func (slice SliceValueType) String() string {
	return slice.Print("\n")
}

// SpiralIterator returns []Coordinate in spiral order
func (slice SliceValueType) SpiralIterator() []Coordinate {
	data := make([]Coordinate, len(slice)*len(slice[0]))
	return spiralTopRight(data, 0, 0, len(slice)-1, len(slice[0])-1)
}
