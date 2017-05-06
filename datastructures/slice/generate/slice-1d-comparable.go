package slice

//go:generate genny -in=$GOFILE -out=../gen-$GOFILE gen "ValueType=int,float64,string,byte"

func (slice SliceValueType) Less(i, j int) bool {
	return slice[i] < slice[j]
}
