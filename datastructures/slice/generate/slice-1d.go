package slice

//go:generate genny -in=$GOFILE -out=../gen-$GOFILE gen "ValueType=int,float64,string,byte,bool"

// NewSliceValueType creates slice length n
func NewSliceValueType(n int) SliceValueType {
	return make([]ValueType, n)
}

// Copy makes a new independent copy of slice
func (slice SliceValueType) Copy() SliceValueType {
	newSlice := NewSliceValueType(len(slice))
	copy(newSlice, slice)
	return newSlice
}

// String is for print
func (slice SliceValueType) String() string {
	return slice.Print(" ")
}
