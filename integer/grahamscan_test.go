package integer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvexHullFalse(t *testing.T) {
	st := NewSliceTuple()
	st.Append(Tuple(0, 0))

	assert.Equal(t, 1, st.ConvexHull(false))

	st.Append(Tuple(0, 10))

	assert.Equal(t, 2, st.ConvexHull(false))

	st.Append(Tuple(10, 10))

	assert.Equal(t, 3, st.ConvexHull(false))

	st.Append(Tuple(10, 0))

	assert.Equal(t, 4, st.ConvexHull(false))

	st.Prepend(Tuple(5, 5))

	assert.Equal(t, 4, st.ConvexHull(false))
	assert.Equal(t, Tuple(0, 0), st.Get(0))
	assert.Equal(t, Tuple(10, 0), st.Get(1))
	assert.Equal(t, Tuple(10, 10), st.Get(2))
	assert.Equal(t, Tuple(0, 10), st.Get(3))

	st.Append(Tuple(10, 5))

	assert.Equal(t, 4, st.ConvexHull(false))
	assert.Equal(t, Tuple(0, 0), st.Get(0))
	assert.Equal(t, Tuple(10, 0), st.Get(1))
	assert.Equal(t, Tuple(10, 10), st.Get(2))
	assert.Equal(t, Tuple(0, 10), st.Get(3))

	st.Append(Tuple(5, -5))

	assert.Equal(t, 5, st.ConvexHull(false))
	assert.Equal(t, Tuple(5, -5), st.Get(0))
	assert.Equal(t, Tuple(10, 0), st.Get(1))
	assert.Equal(t, Tuple(10, 10), st.Get(2))
	assert.Equal(t, Tuple(0, 10), st.Get(3))
	assert.Equal(t, Tuple(0, 0), st.Get(4))

	st.Prepend(Tuple(4, -3))

	assert.Equal(t, 5, st.ConvexHull(false))
	assert.Equal(t, Tuple(5, -5), st.Get(0))
	assert.Equal(t, Tuple(10, 0), st.Get(1))
	assert.Equal(t, Tuple(10, 10), st.Get(2))
	assert.Equal(t, Tuple(0, 10), st.Get(3))
	assert.Equal(t, Tuple(0, 0), st.Get(4))

	st.Append(Tuple(1000, 900))
	assert.Equal(t, 4, st.ConvexHull(false))
	assert.Equal(t, Tuple(5, -5), st.Get(0))
	assert.Equal(t, Tuple(1000, 900), st.Get(1))
	assert.Equal(t, Tuple(0, 10), st.Get(2))
	assert.Equal(t, Tuple(0, 0), st.Get(3))
}

func TestConvexHullFalseProblematic1(t *testing.T) {
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

	assert.Equal(t, 4, st.ConvexHull(false))
}

func TestConvexHullFalseProblematic2(t *testing.T) {
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

	assert.Equal(t, 4, st.ConvexHull(false))
}

func TestConvexHullFalseProblematic3(t *testing.T) {
	st := NewSliceTuple()
	for i := 0; i < 10; i++ {
		st.Append(Tuple(i, 0))
	}

	assert.Equal(t, 2, st.ConvexHull(false))
	assert.Equal(t, Tuple(0, 0), st.Get(0))
	assert.Equal(t, Tuple(9, 0), st.Get(1))
}

func TestConvexHullFalseProblematic4(t *testing.T) {
	st := NewSliceTuple()
	for i := 0; i < 10; i++ {
		st.Append(Tuple(i, 0))
	}
	st.Append(Tuple(0, -1))

	assert.Equal(t, 3, st.ConvexHull(false))
	assert.Equal(t, Tuple(0, -1), st.Get(0))
	assert.Equal(t, Tuple(9, 0), st.Get(1))
	assert.Equal(t, Tuple(0, 0), st.Get(2))
}

func TestConvexHullTrue(t *testing.T) {
	st := NewSliceTuple()
	st.Append(Tuple(0, 0))

	assert.Equal(t, 1, st.ConvexHull(true))

	st.Append(Tuple(0, 10))

	assert.Equal(t, 2, st.ConvexHull(true))

	st.Append(Tuple(10, 10))

	assert.Equal(t, 3, st.ConvexHull(true))

	st.Append(Tuple(10, 0))

	assert.Equal(t, 4, st.ConvexHull(true))

	st.Prepend(Tuple(5, 5))

	assert.Equal(t, 4, st.ConvexHull(true))
	assert.Equal(t, Tuple(0, 0), st.Get(0))
	assert.Equal(t, Tuple(10, 0), st.Get(1))
	assert.Equal(t, Tuple(10, 10), st.Get(2))
	assert.Equal(t, Tuple(0, 10), st.Get(3))

	st.Append(Tuple(10, 5))

	assert.Equal(t, 5, st.ConvexHull(true))
	assert.Equal(t, Tuple(0, 0), st.Get(0))
	assert.Equal(t, Tuple(10, 0), st.Get(1))
	assert.Equal(t, Tuple(10, 5), st.Get(2))
	assert.Equal(t, Tuple(10, 10), st.Get(3))
	assert.Equal(t, Tuple(0, 10), st.Get(4))

	st.Append(Tuple(5, -5))

	assert.Equal(t, 6, st.ConvexHull(true))
	assert.Equal(t, Tuple(5, -5), st.Get(0))
	assert.Equal(t, Tuple(10, 0), st.Get(1))
	assert.Equal(t, Tuple(10, 5), st.Get(2))
	assert.Equal(t, Tuple(10, 10), st.Get(3))
	assert.Equal(t, Tuple(0, 10), st.Get(4))
	assert.Equal(t, Tuple(0, 0), st.Get(5))

	st.Prepend(Tuple(4, -3))

	assert.Equal(t, 6, st.ConvexHull(true))
	assert.Equal(t, Tuple(5, -5), st.Get(0))
	assert.Equal(t, Tuple(10, 0), st.Get(1))
	assert.Equal(t, Tuple(10, 5), st.Get(2))
	assert.Equal(t, Tuple(10, 10), st.Get(3))
	assert.Equal(t, Tuple(0, 10), st.Get(4))
	assert.Equal(t, Tuple(0, 0), st.Get(5))

	st.Append(Tuple(1000, 900))
	assert.Equal(t, 4, st.ConvexHull(true))
	assert.Equal(t, Tuple(5, -5), st.Get(0))
	assert.Equal(t, Tuple(1000, 900), st.Get(1))
	assert.Equal(t, Tuple(0, 10), st.Get(2))
	assert.Equal(t, Tuple(0, 0), st.Get(3))
}

func TestConvexHullTrueProblematic1(t *testing.T) {
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

	assert.Equal(t, 8, st.ConvexHull(true))

	assert.Equal(t, Tuple(0, 0), st.Get(0))
	assert.Equal(t, Tuple(5, 0), st.Get(1))
	assert.Equal(t, Tuple(10, 0), st.Get(2))
	assert.Equal(t, Tuple(10, 5), st.Get(3))
	assert.Equal(t, Tuple(10, 10), st.Get(4))
	assert.Equal(t, Tuple(5, 10), st.Get(5))
	assert.Equal(t, Tuple(0, 10), st.Get(6))
	assert.Equal(t, Tuple(0, 5), st.Get(7))
}

func TestConvexHullTrueProblematic2(t *testing.T) {
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

	assert.Equal(t, 9, st.ConvexHull(true))
}

func TestConvexHullTrueProblematic3(t *testing.T) {
	st := NewSliceTuple()
	for i := 0; i < 10; i++ {
		st.Append(Tuple(i, 0))
	}

	assert.Equal(t, 10, st.ConvexHull(true))
	assert.Equal(t, Tuple(0, 0), st.Get(0))
	assert.Equal(t, Tuple(9, 0), st.Get(9))
}

func TestConvexHullTrueProblematic4(t *testing.T) {
	st := NewSliceTuple()
	for i := 0; i < 10; i++ {
		st.Append(Tuple(i, 0))
	}
	st.Append(Tuple(0, -1))

	assert.Equal(t, 3, st.ConvexHull(false))
	assert.Equal(t, Tuple(0, -1), st.Get(0))
	assert.Equal(t, Tuple(9, 0), st.Get(1))
	assert.Equal(t, Tuple(0, 0), st.Get(2))
}
