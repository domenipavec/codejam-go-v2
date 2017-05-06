// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package llrb

import "github.com/matematik7/codejam-go-v2/datastructures/slice"

// Compare node to newNode
func (node *NodeSliceSliceInt) Compare(v slice.SliceInt) int {
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

// Compare node to newNode
func (node *NodeSliceSliceFloat64) Compare(v slice.SliceFloat64) int {
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

// Compare node to newNode
func (node *NodeSliceSliceString) Compare(v slice.SliceString) int {
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

// Compare node to newNode
func (node *NodeSliceSliceByte) Compare(v slice.SliceByte) int {
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