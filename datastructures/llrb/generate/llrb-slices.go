package llrb

//go:generate genny -in=$GOFILE -out=../gen-$GOFILE gen "ValueType=slice.SliceInt,slice.SliceFloat64,slice.SliceString,slice.SliceByte"

// Compare node to newNode
func (node *NodeValueType) Compare(v ValueType) int {
	for k := range node.Value {
		if node.Value[k] < v[k] {
			return -1
		}
		if node.Value[k] > v[k] {
			return 1
		}
	}
	return 0
}
