package llrb

//go:generate genny -in=$GOFILE -out=../gen-$GOFILE gen "ValueType=int,float64,string,byte"

// Compare node to newNode
func (node *NodeValueType) Compare(v ValueType) int {
	if node.Value < v {
		return -1
	}
	if node.Value > v {
		return 1
	}
	return 0
}
