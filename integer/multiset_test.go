package integer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiSet(t *testing.T) {
	ms := NewMultiSet(Range(3)...)

	assert.True(t, ms.Contains(0))
	assert.True(t, ms.Contains(1))
	assert.True(t, ms.Contains(2))
	assert.False(t, ms.Contains(3))
	assert.False(t, ms.Contains(-1))

	assert.True(t, ms.ContainsAll(Range(3)...))
	assert.False(t, ms.ContainsAll(Range(4)...))

	assert.True(t, ms.ContainsAny(Range(3)...))
	assert.True(t, ms.ContainsAny(Range(4)...))
	assert.False(t, ms.ContainsAny(Range(3, 100)...))

	assert.Equal(t, 1, ms.Count(0))
	assert.Equal(t, 1, ms.Count(1))
	assert.Equal(t, 1, ms.Count(2))
	assert.Equal(t, 0, ms.Count(3))
	assert.Equal(t, 0, ms.Count(-1))

	assert.Equal(t, 3, ms.Len())

	msCopy := ms.Copy()
	ms.Insert(Range(2, 4)...)

	assert.True(t, ms.Contains(0))
	assert.True(t, ms.Contains(1))
	assert.True(t, ms.Contains(2))
	assert.True(t, ms.Contains(3))
	assert.False(t, ms.Contains(4))
	assert.False(t, ms.Contains(-1))

	assert.True(t, ms.ContainsAll(Range(4)...))
	assert.False(t, ms.ContainsAll(Range(5)...))

	assert.True(t, ms.ContainsAny(Range(4)...))
	assert.True(t, ms.ContainsAny(Range(5)...))
	assert.False(t, ms.ContainsAny(Range(4, 100)...))

	assert.Equal(t, 1, ms.Count(0))
	assert.Equal(t, 1, ms.Count(1))
	assert.Equal(t, 2, ms.Count(2))
	assert.Equal(t, 1, ms.Count(3))
	assert.Equal(t, 0, ms.Count(4))
	assert.Equal(t, 0, ms.Count(-1))

	assert.Equal(t, 5, ms.Len())

	assert.True(t, msCopy.Contains(0))
	assert.True(t, msCopy.Contains(1))
	assert.True(t, msCopy.Contains(2))
	assert.False(t, msCopy.Contains(3))
	assert.False(t, msCopy.Contains(-1))

	assert.True(t, msCopy.ContainsAll(Range(3)...))
	assert.False(t, msCopy.ContainsAll(Range(4)...))

	assert.True(t, msCopy.ContainsAny(Range(3)...))
	assert.True(t, msCopy.ContainsAny(Range(4)...))
	assert.False(t, msCopy.ContainsAny(Range(3, 100)...))

	assert.Equal(t, 1, msCopy.Count(0))
	assert.Equal(t, 1, msCopy.Count(1))
	assert.Equal(t, 1, msCopy.Count(2))
	assert.Equal(t, 0, msCopy.Count(3))
	assert.Equal(t, 0, msCopy.Count(-1))

	assert.Equal(t, 3, msCopy.Len())

	ms.InsertN(1, 50)

	assert.True(t, ms.Contains(0))
	assert.True(t, ms.Contains(1))
	assert.True(t, ms.Contains(2))
	assert.True(t, ms.Contains(3))
	assert.False(t, ms.Contains(4))
	assert.False(t, ms.Contains(-1))

	assert.True(t, ms.ContainsAll(Range(4)...))
	assert.False(t, ms.ContainsAll(Range(5)...))

	assert.True(t, ms.ContainsAny(Range(4)...))
	assert.True(t, ms.ContainsAny(Range(5)...))
	assert.False(t, ms.ContainsAny(Range(4, 100)...))

	assert.Equal(t, 1, ms.Count(0))
	assert.Equal(t, 51, ms.Count(1))
	assert.Equal(t, 2, ms.Count(2))
	assert.Equal(t, 1, ms.Count(3))
	assert.Equal(t, 0, ms.Count(4))
	assert.Equal(t, 0, ms.Count(-1))

	assert.Equal(t, 55, ms.Len())

	assert.True(t, msCopy.Contains(0))
	assert.True(t, msCopy.Contains(1))
	assert.True(t, msCopy.Contains(2))
	assert.False(t, msCopy.Contains(3))
	assert.False(t, msCopy.Contains(-1))

	assert.True(t, msCopy.ContainsAll(Range(3)...))
	assert.False(t, msCopy.ContainsAll(Range(4)...))

	assert.True(t, msCopy.ContainsAny(Range(3)...))
	assert.True(t, msCopy.ContainsAny(Range(4)...))
	assert.False(t, msCopy.ContainsAny(Range(3, 100)...))

	assert.Equal(t, 1, msCopy.Count(0))
	assert.Equal(t, 1, msCopy.Count(1))
	assert.Equal(t, 1, msCopy.Count(2))
	assert.Equal(t, 0, msCopy.Count(3))
	assert.Equal(t, 0, msCopy.Count(-1))

	assert.Equal(t, 3, msCopy.Len())

	ms.RemoveOne(Range(1, 4)...)

	assert.True(t, ms.Contains(0))
	assert.True(t, ms.Contains(1))
	assert.True(t, ms.Contains(2))
	assert.False(t, ms.Contains(3))
	assert.False(t, ms.Contains(4))
	assert.False(t, ms.Contains(-1))

	assert.True(t, ms.ContainsAll(Range(3)...))
	assert.False(t, ms.ContainsAll(Range(4)...))

	assert.True(t, ms.ContainsAny(Range(3)...))
	assert.True(t, ms.ContainsAny(Range(4)...))
	assert.False(t, ms.ContainsAny(Range(3, 100)...))

	assert.Equal(t, 1, ms.Count(0))
	assert.Equal(t, 50, ms.Count(1))
	assert.Equal(t, 1, ms.Count(2))
	assert.Equal(t, 0, ms.Count(3))
	assert.Equal(t, 0, ms.Count(4))
	assert.Equal(t, 0, ms.Count(-1))

	assert.Equal(t, 52, ms.Len())

	assert.True(t, msCopy.Contains(0))
	assert.True(t, msCopy.Contains(1))
	assert.True(t, msCopy.Contains(2))
	assert.False(t, msCopy.Contains(3))
	assert.False(t, msCopy.Contains(-1))

	assert.True(t, msCopy.ContainsAll(Range(3)...))
	assert.False(t, msCopy.ContainsAll(Range(4)...))

	assert.True(t, msCopy.ContainsAny(Range(3)...))
	assert.True(t, msCopy.ContainsAny(Range(4)...))
	assert.False(t, msCopy.ContainsAny(Range(3, 100)...))

	assert.Equal(t, 1, msCopy.Count(0))
	assert.Equal(t, 1, msCopy.Count(1))
	assert.Equal(t, 1, msCopy.Count(2))
	assert.Equal(t, 0, msCopy.Count(3))
	assert.Equal(t, 0, msCopy.Count(-1))

	assert.Equal(t, 3, msCopy.Len())

	ms.RemoveAll(1)

	assert.True(t, ms.Contains(0))
	assert.False(t, ms.Contains(1))
	assert.True(t, ms.Contains(2))
	assert.False(t, ms.Contains(3))
	assert.False(t, ms.Contains(4))
	assert.False(t, ms.Contains(-1))

	assert.False(t, ms.ContainsAll(Range(3)...))

	assert.True(t, ms.ContainsAny(Range(3)...))
	assert.True(t, ms.ContainsAny(Range(4)...))
	assert.False(t, ms.ContainsAny(Range(3, 100)...))

	assert.Equal(t, 1, ms.Count(0))
	assert.Equal(t, 0, ms.Count(1))
	assert.Equal(t, 1, ms.Count(2))
	assert.Equal(t, 0, ms.Count(3))
	assert.Equal(t, 0, ms.Count(4))
	assert.Equal(t, 0, ms.Count(-1))

	assert.Equal(t, 2, ms.Len())

	assert.True(t, msCopy.Contains(0))
	assert.True(t, msCopy.Contains(1))
	assert.True(t, msCopy.Contains(2))
	assert.False(t, msCopy.Contains(3))
	assert.False(t, msCopy.Contains(-1))

	assert.True(t, msCopy.ContainsAll(Range(3)...))
	assert.False(t, msCopy.ContainsAll(Range(4)...))

	assert.True(t, msCopy.ContainsAny(Range(3)...))
	assert.True(t, msCopy.ContainsAny(Range(4)...))
	assert.False(t, msCopy.ContainsAny(Range(3, 100)...))

	assert.Equal(t, 1, msCopy.Count(0))
	assert.Equal(t, 1, msCopy.Count(1))
	assert.Equal(t, 1, msCopy.Count(2))
	assert.Equal(t, 0, msCopy.Count(3))
	assert.Equal(t, 0, msCopy.Count(-1))

	assert.Equal(t, 3, msCopy.Len())

	ms.Clear()

	assert.False(t, ms.Contains(0))
	assert.False(t, ms.Contains(1))
	assert.False(t, ms.Contains(2))
	assert.False(t, ms.Contains(3))
	assert.False(t, ms.Contains(4))
	assert.False(t, ms.Contains(-1))

	assert.False(t, ms.ContainsAll(Range(3)...))
	assert.False(t, ms.ContainsAll(Range(4)...))

	assert.False(t, ms.ContainsAny(Range(3)...))
	assert.False(t, ms.ContainsAny(Range(4)...))
	assert.False(t, ms.ContainsAny(Range(3, 100)...))

	assert.Equal(t, 0, ms.Count(0))
	assert.Equal(t, 0, ms.Count(1))
	assert.Equal(t, 0, ms.Count(2))
	assert.Equal(t, 0, ms.Count(3))
	assert.Equal(t, 0, ms.Count(4))
	assert.Equal(t, 0, ms.Count(-1))

	assert.Equal(t, 0, ms.Len())

	assert.True(t, msCopy.Contains(0))
	assert.True(t, msCopy.Contains(1))
	assert.True(t, msCopy.Contains(2))
	assert.False(t, msCopy.Contains(3))
	assert.False(t, msCopy.Contains(-1))

	assert.True(t, msCopy.ContainsAll(Range(3)...))
	assert.False(t, msCopy.ContainsAll(Range(4)...))

	assert.True(t, msCopy.ContainsAny(Range(3)...))
	assert.True(t, msCopy.ContainsAny(Range(4)...))
	assert.False(t, msCopy.ContainsAny(Range(3, 100)...))

	assert.Equal(t, 1, msCopy.Count(0))
	assert.Equal(t, 1, msCopy.Count(1))
	assert.Equal(t, 1, msCopy.Count(2))
	assert.Equal(t, 0, msCopy.Count(3))
	assert.Equal(t, 0, msCopy.Count(-1))

	assert.Equal(t, 3, msCopy.Len())
}
