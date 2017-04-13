package slice

//go:generate genny -in=$GOFILE -out=../gen-$GOFILE gen "ValueType=SliceInt,SliceFloat64,SliceString,SliceByte"

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

func (slice SliceValueType) Less(i, j int) bool {
	for k := range slice[i] {
		if slice[i][k] < slice[j][k] {
			return true
		}
		if slice[i][k] > slice[j][k] {
			return false
		}
	}
	return false
}
