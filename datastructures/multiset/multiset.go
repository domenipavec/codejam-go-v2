package multiset

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "ValueType=BUILTINS"

// ValueType genny generic type for multiset
type ValueType generic.Type

// MultiSetValueType type
type MultiSetValueType map[ValueType]int

// NewValueType creates new set
func NewValueType() MultiSetValueType {
	return make(map[ValueType]int)
}

// Add adds elements
func (s MultiSetValueType) Add(elements ...ValueType) {
	for _, element := range elements {
		s[element]++
	}
}

// AddN adds n elements
func (s MultiSetValueType) AddN(n int, elements ...ValueType) {
	for _, element := range elements {
		s[element] += n
	}
}

// Remove elements
func (s MultiSetValueType) Remove(elements ...ValueType) {
	for _, element := range elements {
		if s[element] >= 1 {
			s[element]--
			if s[element] <= 0 {
				delete(s, element)
			}
		}
	}
}

// RemoveN removes n elements
func (s MultiSetValueType) RemoveN(n int, elements ...ValueType) {
	for _, element := range elements {
		if s[element] >= n {
			s[element] -= n
			if s[element] <= 0 {
				delete(s, element)
			}
		}
	}
}

// Len get set length
func (s MultiSetValueType) Len() int {
	total := 0
	for _, count := range s {
		total += count
	}
	return total
}

// Count number of element in set
func (s MultiSetValueType) Count(element ValueType) int {
	return s[element]
}

// Equal are multisets equal
func (s MultiSetValueType) Equal(other MultiSetValueType) bool {
	if s.Len() != other.Len() {
		return false
	}
	for element, count := range s {
		if other.Count(element) != count {
			return false
		}
	}
	return true
}
