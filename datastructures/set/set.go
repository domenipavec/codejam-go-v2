package set

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "ValueType=BUILTINS"

// ValueType genny generic type for set
type ValueType generic.Type

// SetValueType type
type SetValueType map[ValueType]struct{}

// NewValueType creates new set
func NewValueType() SetValueType {
	return make(map[ValueType]struct{})
}

// Add adds elements
func (s SetValueType) Add(elements ...ValueType) {
	for _, element := range elements {
		s[element] = struct{}{}
	}
}

// Remove elements
func (s SetValueType) Remove(elements ...ValueType) {
	for _, element := range elements {
		delete(s, element)
	}
}

// Len get set length
func (s SetValueType) Len() int {
	return len(s)
}

// Contains check if c in set
func (s SetValueType) Contains(element ValueType) bool {
	_, ok := s[element]
	return ok
}

// Equal are sets equal
func (s SetValueType) Equal(other SetValueType) bool {
	if s.Len() != other.Len() {
		return false
	}
	for element := range s {
		if !other.Contains(element) {
			return false
		}
	}
	return true
}
