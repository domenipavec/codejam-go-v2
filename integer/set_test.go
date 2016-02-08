package integer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	s := NewSet(Range(3)...)

	assert.True(t, s.Contains(0))
	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(2))
	assert.False(t, s.Contains(3))
	assert.False(t, s.Contains(-1))

	assert.True(t, s.ContainsAll(Range(3)...))
	assert.False(t, s.ContainsAll(Range(4)...))

	assert.True(t, s.ContainsAny(Range(3)...))
	assert.True(t, s.ContainsAny(Range(4)...))
	assert.False(t, s.ContainsAny(Range(3, 100)...))

	assert.Equal(t, 3, s.Len())

	sCopy := s.Copy()
	s.Insert(Range(2, 4)...)

	assert.True(t, s.Contains(0))
	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(2))
	assert.True(t, s.Contains(3))
	assert.False(t, s.Contains(4))
	assert.False(t, s.Contains(-1))

	assert.True(t, s.ContainsAll(Range(4)...))
	assert.False(t, s.ContainsAll(Range(5)...))

	assert.True(t, s.ContainsAny(Range(4)...))
	assert.True(t, s.ContainsAny(Range(5)...))
	assert.False(t, s.ContainsAny(Range(4, 100)...))

	assert.Equal(t, 4, s.Len())

	assert.True(t, sCopy.Contains(0))
	assert.True(t, sCopy.Contains(1))
	assert.True(t, sCopy.Contains(2))
	assert.False(t, sCopy.Contains(3))
	assert.False(t, sCopy.Contains(-1))

	assert.True(t, sCopy.ContainsAll(Range(3)...))
	assert.False(t, sCopy.ContainsAll(Range(4)...))

	assert.True(t, sCopy.ContainsAny(Range(3)...))
	assert.True(t, sCopy.ContainsAny(Range(4)...))
	assert.False(t, sCopy.ContainsAny(Range(3, 100)...))

	assert.Equal(t, 3, sCopy.Len())

	s.Remove(Range(1, 4)...)

	assert.True(t, s.Contains(0))
	assert.False(t, s.Contains(1))
	assert.False(t, s.Contains(2))
	assert.False(t, s.Contains(3))
	assert.False(t, s.Contains(4))
	assert.False(t, s.Contains(-1))

	assert.True(t, s.ContainsAll(0))
	assert.False(t, s.ContainsAll(Range(4)...))

	assert.True(t, s.ContainsAny(Range(3)...))
	assert.True(t, s.ContainsAny(Range(4)...))
	assert.False(t, s.ContainsAny(Range(3, 100)...))

	assert.Equal(t, 1, s.Len())

	assert.True(t, sCopy.Contains(0))
	assert.True(t, sCopy.Contains(1))
	assert.True(t, sCopy.Contains(2))
	assert.False(t, sCopy.Contains(3))
	assert.False(t, sCopy.Contains(-1))

	assert.True(t, sCopy.ContainsAll(Range(3)...))
	assert.False(t, sCopy.ContainsAll(Range(4)...))

	assert.True(t, sCopy.ContainsAny(Range(3)...))
	assert.True(t, sCopy.ContainsAny(Range(4)...))
	assert.False(t, sCopy.ContainsAny(Range(3, 100)...))

	assert.Equal(t, 3, sCopy.Len())

	s.Clear()

	assert.False(t, s.Contains(0))
	assert.False(t, s.Contains(1))
	assert.False(t, s.Contains(2))
	assert.False(t, s.Contains(3))
	assert.False(t, s.Contains(4))
	assert.False(t, s.Contains(-1))

	assert.False(t, s.ContainsAll(Range(3)...))
	assert.False(t, s.ContainsAll(Range(4)...))

	assert.False(t, s.ContainsAny(Range(3)...))
	assert.False(t, s.ContainsAny(Range(4)...))
	assert.False(t, s.ContainsAny(Range(3, 100)...))

	assert.Equal(t, 0, s.Len())

	assert.True(t, sCopy.Contains(0))
	assert.True(t, sCopy.Contains(1))
	assert.True(t, sCopy.Contains(2))
	assert.False(t, sCopy.Contains(3))
	assert.False(t, sCopy.Contains(-1))

	assert.True(t, sCopy.ContainsAll(Range(3)...))
	assert.False(t, sCopy.ContainsAll(Range(4)...))

	assert.True(t, sCopy.ContainsAny(Range(3)...))
	assert.True(t, sCopy.ContainsAny(Range(4)...))
	assert.False(t, sCopy.ContainsAny(Range(3, 100)...))

	assert.Equal(t, 3, sCopy.Len())
}
