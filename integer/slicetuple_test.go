package integer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceTuple(t *testing.T) {
	st := NewSliceTuple(2, 3)

	assert.Equal(t, 2, len(st))
	assert.Equal(t, Tuple(2), st.Get(0))
	assert.Equal(t, Tuple(3), st.Get(1))

	st.Prepend(Tuple(4))

	assert.Equal(t, 3, len(st))
	assert.Equal(t, Tuple(4), st.Get(0))
	assert.Equal(t, Tuple(2), st.Get(1))
	assert.Equal(t, Tuple(3), st.Get(2))

	st.Append(Tuple(5))

	assert.Equal(t, 4, len(st))
	assert.Equal(t, Tuple(4), st.Get(0))
	assert.Equal(t, Tuple(2), st.Get(1))
	assert.Equal(t, Tuple(3), st.Get(2))
	assert.Equal(t, Tuple(5), st.Get(3))

	st.Insert(2, Tuple(1))

	assert.Equal(t, 5, len(st))
	assert.Equal(t, Tuple(4), st.Get(0))
	assert.Equal(t, Tuple(2), st.Get(1))
	assert.Equal(t, Tuple(1), st.Get(2))
	assert.Equal(t, Tuple(3), st.Get(3))
	assert.Equal(t, Tuple(5), st.Get(4))

	st.PrefixConst(-1)

	assert.Equal(t, 5, len(st))
	assert.Equal(t, Tuple(-1, 4), st.Get(0))
	assert.Equal(t, Tuple(-1, 2), st.Get(1))
	assert.Equal(t, Tuple(-1, 1), st.Get(2))
	assert.Equal(t, Tuple(-1, 3), st.Get(3))
	assert.Equal(t, Tuple(-1, 5), st.Get(4))

	st.PrefixIndex()

	assert.Equal(t, 5, len(st))
	assert.Equal(t, Tuple(0, -1, 4), st.Get(0))
	assert.Equal(t, Tuple(1, -1, 2), st.Get(1))
	assert.Equal(t, Tuple(2, -1, 1), st.Get(2))
	assert.Equal(t, Tuple(3, -1, 3), st.Get(3))
	assert.Equal(t, Tuple(4, -1, 5), st.Get(4))

	st.PostfixConst(1)

	assert.Equal(t, 5, len(st))
	assert.Equal(t, Tuple(0, -1, 4, 1), st.Get(0))
	assert.Equal(t, Tuple(1, -1, 2, 1), st.Get(1))
	assert.Equal(t, Tuple(2, -1, 1, 1), st.Get(2))
	assert.Equal(t, Tuple(3, -1, 3, 1), st.Get(3))
	assert.Equal(t, Tuple(4, -1, 5, 1), st.Get(4))

	st.PostfixIndex()

	assert.Equal(t, 5, len(st))
	assert.Equal(t, Tuple(0, -1, 4, 1, 0), st.Get(0))
	assert.Equal(t, Tuple(1, -1, 2, 1, 1), st.Get(1))
	assert.Equal(t, Tuple(2, -1, 1, 1, 2), st.Get(2))
	assert.Equal(t, Tuple(3, -1, 3, 1, 3), st.Get(3))
	assert.Equal(t, Tuple(4, -1, 5, 1, 4), st.Get(4))

	st.Swap(2, 3)

	assert.Equal(t, 5, len(st))
	assert.Equal(t, Tuple(0, -1, 4, 1, 0), st.Get(0))
	assert.Equal(t, Tuple(1, -1, 2, 1, 1), st.Get(1))
	assert.Equal(t, Tuple(3, -1, 3, 1, 3), st.Get(2))
	assert.Equal(t, Tuple(2, -1, 1, 1, 2), st.Get(3))
	assert.Equal(t, Tuple(4, -1, 5, 1, 4), st.Get(4))

	st.Reverse()

	assert.Equal(t, 5, len(st))
	assert.Equal(t, Tuple(4, -1, 5, 1, 4), st.Get(0))
	assert.Equal(t, Tuple(2, -1, 1, 1, 2), st.Get(1))
	assert.Equal(t, Tuple(3, -1, 3, 1, 3), st.Get(2))
	assert.Equal(t, Tuple(1, -1, 2, 1, 1), st.Get(3))
	assert.Equal(t, Tuple(0, -1, 4, 1, 0), st.Get(4))

	st.SortAsc()

	assert.Equal(t, 5, len(st))
	assert.Equal(t, Tuple(0, -1, 4, 1, 0), st.Get(0))
	assert.Equal(t, Tuple(1, -1, 2, 1, 1), st.Get(1))
	assert.Equal(t, Tuple(2, -1, 1, 1, 2), st.Get(2))
	assert.Equal(t, Tuple(3, -1, 3, 1, 3), st.Get(3))
	assert.Equal(t, Tuple(4, -1, 5, 1, 4), st.Get(4))

	st.SortAscBy(2)

	assert.Equal(t, 5, len(st))
	assert.Equal(t, Tuple(2, -1, 1, 1, 2), st.Get(0))
	assert.Equal(t, Tuple(1, -1, 2, 1, 1), st.Get(1))
	assert.Equal(t, Tuple(3, -1, 3, 1, 3), st.Get(2))
	assert.Equal(t, Tuple(0, -1, 4, 1, 0), st.Get(3))
	assert.Equal(t, Tuple(4, -1, 5, 1, 4), st.Get(4))

	st.SortDescBy(2)

	assert.Equal(t, 5, len(st))
	assert.Equal(t, Tuple(4, -1, 5, 1, 4), st.Get(0))
	assert.Equal(t, Tuple(0, -1, 4, 1, 0), st.Get(1))
	assert.Equal(t, Tuple(3, -1, 3, 1, 3), st.Get(2))
	assert.Equal(t, Tuple(1, -1, 2, 1, 1), st.Get(3))
	assert.Equal(t, Tuple(2, -1, 1, 1, 2), st.Get(4))

	st.SortDesc()

	assert.Equal(t, 5, len(st))
	assert.Equal(t, Tuple(4, -1, 5, 1, 4), st.Get(0))
	assert.Equal(t, Tuple(3, -1, 3, 1, 3), st.Get(1))
	assert.Equal(t, Tuple(2, -1, 1, 1, 2), st.Get(2))
	assert.Equal(t, Tuple(1, -1, 2, 1, 1), st.Get(3))
	assert.Equal(t, Tuple(0, -1, 4, 1, 0), st.Get(4))

	st1 := st.Copy()
	st2 := st.CopySlice()
	st.Get(0)[0] = 3

	assert.Equal(t, 5, len(st2))
	assert.Equal(t, Tuple(3, -1, 5, 1, 4), st2.Get(0))
	assert.Equal(t, Tuple(3, -1, 3, 1, 3), st2.Get(1))
	assert.Equal(t, Tuple(2, -1, 1, 1, 2), st2.Get(2))
	assert.Equal(t, Tuple(1, -1, 2, 1, 1), st2.Get(3))
	assert.Equal(t, Tuple(0, -1, 4, 1, 0), st2.Get(4))

	assert.Equal(t, 5, len(st1))
	assert.Equal(t, Tuple(4, -1, 5, 1, 4), st1.Get(0))
	assert.Equal(t, Tuple(3, -1, 3, 1, 3), st1.Get(1))
	assert.Equal(t, Tuple(2, -1, 1, 1, 2), st1.Get(2))
	assert.Equal(t, Tuple(1, -1, 2, 1, 1), st1.Get(3))
	assert.Equal(t, Tuple(0, -1, 4, 1, 0), st1.Get(4))

	assert.Equal(t, 5, len(st))
	assert.Equal(t, Tuple(3, -1, 5, 1, 4), st.Get(0))
	assert.Equal(t, Tuple(3, -1, 3, 1, 3), st.Get(1))
	assert.Equal(t, Tuple(2, -1, 1, 1, 2), st.Get(2))
	assert.Equal(t, Tuple(1, -1, 2, 1, 1), st.Get(3))
	assert.Equal(t, Tuple(0, -1, 4, 1, 0), st.Get(4))

	assert.Equal(t, Tuple(2, -1, 1, 1, 2), st.Delete(2))

	assert.Equal(t, 5, len(st2))
	assert.Equal(t, Tuple(3, -1, 5, 1, 4), st2.Get(0))
	assert.Equal(t, Tuple(3, -1, 3, 1, 3), st2.Get(1))
	assert.Equal(t, Tuple(2, -1, 1, 1, 2), st2.Get(2))
	assert.Equal(t, Tuple(1, -1, 2, 1, 1), st2.Get(3))
	assert.Equal(t, Tuple(0, -1, 4, 1, 0), st2.Get(4))

	assert.Equal(t, 5, len(st1))
	assert.Equal(t, Tuple(4, -1, 5, 1, 4), st1.Get(0))
	assert.Equal(t, Tuple(3, -1, 3, 1, 3), st1.Get(1))
	assert.Equal(t, Tuple(2, -1, 1, 1, 2), st1.Get(2))
	assert.Equal(t, Tuple(1, -1, 2, 1, 1), st1.Get(3))
	assert.Equal(t, Tuple(0, -1, 4, 1, 0), st1.Get(4))

	assert.Equal(t, 4, len(st))
	assert.Equal(t, Tuple(3, -1, 5, 1, 4), st.Get(0))
	assert.Equal(t, Tuple(3, -1, 3, 1, 3), st.Get(1))
	assert.Equal(t, Tuple(1, -1, 2, 1, 1), st.Get(2))
	assert.Equal(t, Tuple(0, -1, 4, 1, 0), st.Get(3))

	assert.Equal(t, Tuple(3, -1, 5, 1, 4), st.DeleteFirst())

	assert.Equal(t, 5, len(st2))
	assert.Equal(t, Tuple(3, -1, 5, 1, 4), st2.Get(0))
	assert.Equal(t, Tuple(3, -1, 3, 1, 3), st2.Get(1))
	assert.Equal(t, Tuple(2, -1, 1, 1, 2), st2.Get(2))
	assert.Equal(t, Tuple(1, -1, 2, 1, 1), st2.Get(3))
	assert.Equal(t, Tuple(0, -1, 4, 1, 0), st2.Get(4))

	assert.Equal(t, 5, len(st1))
	assert.Equal(t, Tuple(4, -1, 5, 1, 4), st1.Get(0))
	assert.Equal(t, Tuple(3, -1, 3, 1, 3), st1.Get(1))
	assert.Equal(t, Tuple(2, -1, 1, 1, 2), st1.Get(2))
	assert.Equal(t, Tuple(1, -1, 2, 1, 1), st1.Get(3))
	assert.Equal(t, Tuple(0, -1, 4, 1, 0), st1.Get(4))

	assert.Equal(t, 3, len(st))
	assert.Equal(t, Tuple(3, -1, 3, 1, 3), st.Get(0))
	assert.Equal(t, Tuple(1, -1, 2, 1, 1), st.Get(1))
	assert.Equal(t, Tuple(0, -1, 4, 1, 0), st.Get(2))

	assert.Equal(t, Tuple(0, -1, 4, 1, 0), st.DeleteLast())

	assert.Equal(t, 5, len(st2))
	assert.Equal(t, Tuple(3, -1, 5, 1, 4), st2.Get(0))
	assert.Equal(t, Tuple(3, -1, 3, 1, 3), st2.Get(1))
	assert.Equal(t, Tuple(2, -1, 1, 1, 2), st2.Get(2))
	assert.Equal(t, Tuple(1, -1, 2, 1, 1), st2.Get(3))
	assert.Equal(t, Tuple(0, -1, 4, 1, 0), st2.Get(4))

	assert.Equal(t, 5, len(st1))
	assert.Equal(t, Tuple(4, -1, 5, 1, 4), st1.Get(0))
	assert.Equal(t, Tuple(3, -1, 3, 1, 3), st1.Get(1))
	assert.Equal(t, Tuple(2, -1, 1, 1, 2), st1.Get(2))
	assert.Equal(t, Tuple(1, -1, 2, 1, 1), st1.Get(3))
	assert.Equal(t, Tuple(0, -1, 4, 1, 0), st1.Get(4))

	assert.Equal(t, 2, len(st))
	assert.Equal(t, Tuple(3, -1, 3, 1, 3), st.Get(0))
	assert.Equal(t, Tuple(1, -1, 2, 1, 1), st.Get(1))
}

func TestSliceTupleRelease(t *testing.T) {
	st := NewSliceTuple(2, 3)
	for i := 0; i < 10000; i++ {
		stc := st.Copy()
		stsc := st.CopySlice()
		st.Get(0)[0] = 0
		assert.Equal(t, Tuple(0), stsc.Get(0))
		assert.Equal(t, Tuple(2), stc.Get(0))
		stc.Delete(0)
		stsc.Delete(0)
		assert.Equal(t, 2, len(st))
		st.Get(0)[0] = 2
		stc.CopyRelease()
		stsc.CopySliceRelease()
	}
}
