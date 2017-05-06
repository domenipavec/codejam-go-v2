package slice

//go:generate genny -in=$GOFILE -out=../gen-$GOFILE gen "ValueType=SliceInt,SliceFloat64,SliceString,SliceByte"

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
