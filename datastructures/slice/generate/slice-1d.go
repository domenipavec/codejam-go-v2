package slice

//go:generate genny -in=$GOFILE -out=../gen-$GOFILE gen "ValueType=int,float64,string,byte"

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

func (slice SliceValueType) Less(i, j int) bool {
	return slice[i] < slice[j]
}
