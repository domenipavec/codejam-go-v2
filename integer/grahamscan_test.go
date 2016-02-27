package integer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvexHull(t *testing.T) {
	st := NewSliceTuple()
	st.Append(Tuple(0, 0))

	assert.Equal(t, 1, st.ConvexHull())

	st.Append(Tuple(0, 10))

	assert.Equal(t, 2, st.ConvexHull())

	st.Append(Tuple(10, 10))

	assert.Equal(t, 3, st.ConvexHull())

	st.Append(Tuple(10, 0))

	assert.Equal(t, 4, st.ConvexHull())

	st.Prepend(Tuple(5, 5))

	assert.Equal(t, 4, st.ConvexHull())
	assert.Equal(t, Tuple(0, 0), st.Get(0))
	assert.Equal(t, Tuple(10, 0), st.Get(1))
	assert.Equal(t, Tuple(10, 10), st.Get(2))
	assert.Equal(t, Tuple(0, 10), st.Get(3))

	st.Append(Tuple(10, 5))

	assert.Equal(t, 4, st.ConvexHull())
	assert.Equal(t, Tuple(0, 0), st.Get(0))
	assert.Equal(t, Tuple(10, 0), st.Get(1))
	assert.Equal(t, Tuple(10, 10), st.Get(2))
	assert.Equal(t, Tuple(0, 10), st.Get(3))

	st.Append(Tuple(5, -5))

	assert.Equal(t, 5, st.ConvexHull())
	assert.Equal(t, Tuple(5, -5), st.Get(0))
	assert.Equal(t, Tuple(10, 0), st.Get(1))
	assert.Equal(t, Tuple(10, 10), st.Get(2))
	assert.Equal(t, Tuple(0, 10), st.Get(3))
	assert.Equal(t, Tuple(0, 0), st.Get(4))

	st.Prepend(Tuple(4, -3))

	assert.Equal(t, 5, st.ConvexHull())
	assert.Equal(t, Tuple(5, -5), st.Get(0))
	assert.Equal(t, Tuple(10, 0), st.Get(1))
	assert.Equal(t, Tuple(10, 10), st.Get(2))
	assert.Equal(t, Tuple(0, 10), st.Get(3))
	assert.Equal(t, Tuple(0, 0), st.Get(4))

	st.Append(Tuple(1000, 900))
	assert.Equal(t, 4, st.ConvexHull())
	assert.Equal(t, Tuple(5, -5), st.Get(0))
	assert.Equal(t, Tuple(1000, 900), st.Get(1))
	assert.Equal(t, Tuple(0, 10), st.Get(2))
	assert.Equal(t, Tuple(0, 0), st.Get(3))
}

func TestConvexHullProblematic1(t *testing.T) {
	st := NewSliceTuple()
	st.Append(Tuple(0, 0))
	st.Append(Tuple(5, 0))
	st.Append(Tuple(10, 0))
	st.Append(Tuple(0, 5))
	st.Append(Tuple(5, 5))
	st.Append(Tuple(10, 5))
	st.Append(Tuple(0, 10))
	st.Append(Tuple(5, 10))
	st.Append(Tuple(10, 10))

	assert.Equal(t, 4, st.ConvexHull())

	for _, t := range st {
		fmt.Println(*t)
	}
}

func TestConvexHullProblematic2(t *testing.T) {
	st := NewSliceTuple()
	st.Append(Tuple(0, 0))
	st.Append(Tuple(10, 0))
	st.Append(Tuple(5, 0))
	st.Append(Tuple(0, 5))
	st.Append(Tuple(5, 5))
	st.Append(Tuple(10, 5))
	st.Append(Tuple(0, 8))
	st.Append(Tuple(0, 10))
	st.Append(Tuple(5, 10))
	st.Append(Tuple(10, 10))

	assert.Equal(t, 4, st.ConvexHull())
}

func TestConvexHullProblematic3(t *testing.T) {
	st := NewSliceTuple()
	for i := 0; i < 10; i++ {
		st.Append(Tuple(i, 0))
	}

	assert.Equal(t, 2, st.ConvexHull())
	assert.Equal(t, Tuple(0, 0), st.Get(0))
	assert.Equal(t, Tuple(9, 0), st.Get(1))
}

func TestConvexHullProblematic4(t *testing.T) {
	st := NewSliceTuple()
	for i := 0; i < 10; i++ {
		st.Append(Tuple(i, 0))
	}
	st.Append(Tuple(0, -1))

	assert.Equal(t, 3, st.ConvexHull())
	assert.Equal(t, Tuple(0, -1), st.Get(0))
	assert.Equal(t, Tuple(9, 0), st.Get(1))
	assert.Equal(t, Tuple(0, 0), st.Get(2))
}
