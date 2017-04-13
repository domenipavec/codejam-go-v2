package slice

// PostfixIndex insert index of element as last element
func (slice SliceSliceInt) PostfixIndex() {
	for i := range slice {
		slice[i].Append(i)
	}
}

// PrefixIndex insert index of element as last element
func (slice SliceSliceInt) PrefixIndex() {
	for i := range slice {
		slice[i].Prepend(i)
	}
}

// PostfixIndex insert index of element as last element
func (slice SliceSliceFloat64) PostfixIndex() {
	for i := range slice {
		slice[i].Append(float64(i))
	}
}

// PrefixIndex insert index of element as last element
func (slice SliceSliceFloat64) PrefixIndex() {
	for i := range slice {
		slice[i].Prepend(float64(i))
	}
}
