package st

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceTupleInt(t *testing.T) {
	st := FromInts(1, 2, 3)

	assert.Equal(t, 2, st.Len())
	assert.Equal(t, IntTuple(2), st.Tuples[0])
	assert.Equal(t, IntTuple(3), st.Tuples[1])

	st.Prepend(IntTuple(4))

	assert.Equal(t, 3, st.Len())
	assert.Equal(t, IntTuple(4), st.Tuples[0])
	assert.Equal(t, IntTuple(2), st.Tuples[1])
	assert.Equal(t, IntTuple(3), st.Tuples[2])

	st.Append(IntTuple(5))

	assert.Equal(t, 4, st.Len())
	assert.Equal(t, IntTuple(4), st.Tuples[0])
	assert.Equal(t, IntTuple(2), st.Tuples[1])
	assert.Equal(t, IntTuple(3), st.Tuples[2])
	assert.Equal(t, IntTuple(5), st.Tuples[3])

	st.Insert(2, IntTuple(1))

	assert.Equal(t, 5, st.Len())
	assert.Equal(t, IntTuple(4), st.Tuples[0])
	assert.Equal(t, IntTuple(2), st.Tuples[1])
	assert.Equal(t, IntTuple(1), st.Tuples[2])
	assert.Equal(t, IntTuple(3), st.Tuples[3])
	assert.Equal(t, IntTuple(5), st.Tuples[4])

	st.PrefixIntConst(-1)

	assert.Equal(t, 5, st.Len())
	assert.Equal(t, IntTuple(-1, 4), st.Tuples[0])
	assert.Equal(t, IntTuple(-1, 2), st.Tuples[1])
	assert.Equal(t, IntTuple(-1, 1), st.Tuples[2])
	assert.Equal(t, IntTuple(-1, 3), st.Tuples[3])
	assert.Equal(t, IntTuple(-1, 5), st.Tuples[4])

	st.PrefixIntIndex()

	assert.Equal(t, 5, st.Len())
	assert.Equal(t, IntTuple(0, -1, 4), st.Tuples[0])
	assert.Equal(t, IntTuple(1, -1, 2), st.Tuples[1])
	assert.Equal(t, IntTuple(2, -1, 1), st.Tuples[2])
	assert.Equal(t, IntTuple(3, -1, 3), st.Tuples[3])
	assert.Equal(t, IntTuple(4, -1, 5), st.Tuples[4])

	st.PostfixIntConst(1)

	assert.Equal(t, 5, st.Len())
	assert.Equal(t, IntTuple(0, -1, 4, 1), st.Tuples[0])
	assert.Equal(t, IntTuple(1, -1, 2, 1), st.Tuples[1])
	assert.Equal(t, IntTuple(2, -1, 1, 1), st.Tuples[2])
	assert.Equal(t, IntTuple(3, -1, 3, 1), st.Tuples[3])
	assert.Equal(t, IntTuple(4, -1, 5, 1), st.Tuples[4])

	st.PostfixIntIndex()

	assert.Equal(t, 5, st.Len())
	assert.Equal(t, IntTuple(0, -1, 4, 1, 0), st.Tuples[0])
	assert.Equal(t, IntTuple(1, -1, 2, 1, 1), st.Tuples[1])
	assert.Equal(t, IntTuple(2, -1, 1, 1, 2), st.Tuples[2])
	assert.Equal(t, IntTuple(3, -1, 3, 1, 3), st.Tuples[3])
	assert.Equal(t, IntTuple(4, -1, 5, 1, 4), st.Tuples[4])

	assert.Equal(t, "0 -1 4 1 0", st.Get(0).String())

	st.Swap(2, 3)

	assert.Equal(t, 5, st.Len())
	assert.Equal(t, IntTuple(0, -1, 4, 1, 0), st.Tuples[0])
	assert.Equal(t, IntTuple(1, -1, 2, 1, 1), st.Tuples[1])
	assert.Equal(t, IntTuple(3, -1, 3, 1, 3), st.Tuples[2])
	assert.Equal(t, IntTuple(2, -1, 1, 1, 2), st.Tuples[3])
	assert.Equal(t, IntTuple(4, -1, 5, 1, 4), st.Tuples[4])

	st.Reverse()

	assert.Equal(t, 5, st.Len())
	assert.Equal(t, IntTuple(4, -1, 5, 1, 4), st.Tuples[0])
	assert.Equal(t, IntTuple(2, -1, 1, 1, 2), st.Tuples[1])
	assert.Equal(t, IntTuple(3, -1, 3, 1, 3), st.Tuples[2])
	assert.Equal(t, IntTuple(1, -1, 2, 1, 1), st.Tuples[3])
	assert.Equal(t, IntTuple(0, -1, 4, 1, 0), st.Tuples[4])

	st.SortOrder(IntAsc(0))
	st.Sort()

	assert.Equal(t, 5, st.Len())
	assert.Equal(t, IntTuple(0, -1, 4, 1, 0), st.Tuples[0])
	assert.Equal(t, IntTuple(1, -1, 2, 1, 1), st.Tuples[1])
	assert.Equal(t, IntTuple(2, -1, 1, 1, 2), st.Tuples[2])
	assert.Equal(t, IntTuple(3, -1, 3, 1, 3), st.Tuples[3])
	assert.Equal(t, IntTuple(4, -1, 5, 1, 4), st.Tuples[4])

	st.SortOrder(IntAsc(2))
	st.Sort()

	assert.Equal(t, 5, st.Len())
	assert.Equal(t, IntTuple(2, -1, 1, 1, 2), st.Tuples[0])
	assert.Equal(t, IntTuple(1, -1, 2, 1, 1), st.Tuples[1])
	assert.Equal(t, IntTuple(3, -1, 3, 1, 3), st.Tuples[2])
	assert.Equal(t, IntTuple(0, -1, 4, 1, 0), st.Tuples[3])
	assert.Equal(t, IntTuple(4, -1, 5, 1, 4), st.Tuples[4])

	st.SortOrder(IntDesc(2))
	st.Sort()

	assert.Equal(t, 5, st.Len())
	assert.Equal(t, IntTuple(4, -1, 5, 1, 4), st.Tuples[0])
	assert.Equal(t, IntTuple(0, -1, 4, 1, 0), st.Tuples[1])
	assert.Equal(t, IntTuple(3, -1, 3, 1, 3), st.Tuples[2])
	assert.Equal(t, IntTuple(1, -1, 2, 1, 1), st.Tuples[3])
	assert.Equal(t, IntTuple(2, -1, 1, 1, 2), st.Tuples[4])

	st.SortOrder(IntDesc(0))
	st.Sort()

	assert.Equal(t, 5, st.Len())
	assert.Equal(t, IntTuple(4, -1, 5, 1, 4), st.Tuples[0])
	assert.Equal(t, IntTuple(3, -1, 3, 1, 3), st.Tuples[1])
	assert.Equal(t, IntTuple(2, -1, 1, 1, 2), st.Tuples[2])
	assert.Equal(t, IntTuple(1, -1, 2, 1, 1), st.Tuples[3])
	assert.Equal(t, IntTuple(0, -1, 4, 1, 0), st.Tuples[4])

	st1 := st.Copy()
	st2 := st.CopySlice()
	st.Tuples[0].Ints[0] = 3

	assert.Equal(t, 5, st2.Len())
	assert.Equal(t, IntTuple(3, -1, 5, 1, 4), st2.Tuples[0])
	assert.Equal(t, IntTuple(3, -1, 3, 1, 3), st2.Tuples[1])
	assert.Equal(t, IntTuple(2, -1, 1, 1, 2), st2.Tuples[2])
	assert.Equal(t, IntTuple(1, -1, 2, 1, 1), st2.Tuples[3])
	assert.Equal(t, IntTuple(0, -1, 4, 1, 0), st2.Tuples[4])

	assert.Equal(t, 5, st1.Len())
	assert.Equal(t, IntTuple(4, -1, 5, 1, 4), st1.Tuples[0])
	assert.Equal(t, IntTuple(3, -1, 3, 1, 3), st1.Tuples[1])
	assert.Equal(t, IntTuple(2, -1, 1, 1, 2), st1.Tuples[2])
	assert.Equal(t, IntTuple(1, -1, 2, 1, 1), st1.Tuples[3])
	assert.Equal(t, IntTuple(0, -1, 4, 1, 0), st1.Tuples[4])

	assert.Equal(t, 5, st.Len())
	assert.Equal(t, IntTuple(3, -1, 5, 1, 4), st.Tuples[0])
	assert.Equal(t, IntTuple(3, -1, 3, 1, 3), st.Tuples[1])
	assert.Equal(t, IntTuple(2, -1, 1, 1, 2), st.Tuples[2])
	assert.Equal(t, IntTuple(1, -1, 2, 1, 1), st.Tuples[3])
	assert.Equal(t, IntTuple(0, -1, 4, 1, 0), st.Tuples[4])

	assert.Equal(t, IntTuple(2, -1, 1, 1, 2), st.Remove(2))

	assert.Equal(t, 5, st2.Len())
	assert.Equal(t, IntTuple(3, -1, 5, 1, 4), st2.Tuples[0])
	assert.Equal(t, IntTuple(3, -1, 3, 1, 3), st2.Tuples[1])
	assert.Equal(t, IntTuple(2, -1, 1, 1, 2), st2.Tuples[2])
	assert.Equal(t, IntTuple(1, -1, 2, 1, 1), st2.Tuples[3])
	assert.Equal(t, IntTuple(0, -1, 4, 1, 0), st2.Tuples[4])

	assert.Equal(t, 5, st1.Len())
	assert.Equal(t, IntTuple(4, -1, 5, 1, 4), st1.Tuples[0])
	assert.Equal(t, IntTuple(3, -1, 3, 1, 3), st1.Tuples[1])
	assert.Equal(t, IntTuple(2, -1, 1, 1, 2), st1.Tuples[2])
	assert.Equal(t, IntTuple(1, -1, 2, 1, 1), st1.Tuples[3])
	assert.Equal(t, IntTuple(0, -1, 4, 1, 0), st1.Tuples[4])

	assert.Equal(t, 4, st.Len())
	assert.Equal(t, IntTuple(3, -1, 5, 1, 4), st.Tuples[0])
	assert.Equal(t, IntTuple(3, -1, 3, 1, 3), st.Tuples[1])
	assert.Equal(t, IntTuple(1, -1, 2, 1, 1), st.Tuples[2])
	assert.Equal(t, IntTuple(0, -1, 4, 1, 0), st.Tuples[3])

	assert.Equal(t, IntTuple(3, -1, 5, 1, 4), st.RemoveFirst())

	assert.Equal(t, 5, st2.Len())
	assert.Equal(t, IntTuple(3, -1, 5, 1, 4), st2.Tuples[0])
	assert.Equal(t, IntTuple(3, -1, 3, 1, 3), st2.Tuples[1])
	assert.Equal(t, IntTuple(2, -1, 1, 1, 2), st2.Tuples[2])
	assert.Equal(t, IntTuple(1, -1, 2, 1, 1), st2.Tuples[3])
	assert.Equal(t, IntTuple(0, -1, 4, 1, 0), st2.Tuples[4])

	assert.Equal(t, 5, st1.Len())
	assert.Equal(t, IntTuple(4, -1, 5, 1, 4), st1.Tuples[0])
	assert.Equal(t, IntTuple(3, -1, 3, 1, 3), st1.Tuples[1])
	assert.Equal(t, IntTuple(2, -1, 1, 1, 2), st1.Tuples[2])
	assert.Equal(t, IntTuple(1, -1, 2, 1, 1), st1.Tuples[3])
	assert.Equal(t, IntTuple(0, -1, 4, 1, 0), st1.Tuples[4])

	assert.Equal(t, 3, st.Len())
	assert.Equal(t, IntTuple(3, -1, 3, 1, 3), st.Tuples[0])
	assert.Equal(t, IntTuple(1, -1, 2, 1, 1), st.Tuples[1])
	assert.Equal(t, IntTuple(0, -1, 4, 1, 0), st.Tuples[2])

	assert.Equal(t, IntTuple(0, -1, 4, 1, 0), st.RemoveLast())

	assert.Equal(t, 5, st2.Len())
	assert.Equal(t, IntTuple(3, -1, 5, 1, 4), st2.Tuples[0])
	assert.Equal(t, IntTuple(3, -1, 3, 1, 3), st2.Tuples[1])
	assert.Equal(t, IntTuple(2, -1, 1, 1, 2), st2.Tuples[2])
	assert.Equal(t, IntTuple(1, -1, 2, 1, 1), st2.Tuples[3])
	assert.Equal(t, IntTuple(0, -1, 4, 1, 0), st2.Tuples[4])

	assert.Equal(t, 5, st1.Len())
	assert.Equal(t, IntTuple(4, -1, 5, 1, 4), st1.Tuples[0])
	assert.Equal(t, IntTuple(3, -1, 3, 1, 3), st1.Tuples[1])
	assert.Equal(t, IntTuple(2, -1, 1, 1, 2), st1.Tuples[2])
	assert.Equal(t, IntTuple(1, -1, 2, 1, 1), st1.Tuples[3])
	assert.Equal(t, IntTuple(0, -1, 4, 1, 0), st1.Tuples[4])

	assert.Equal(t, 2, st.Len())
	assert.Equal(t, IntTuple(3, -1, 3, 1, 3), st.Tuples[0])
	assert.Equal(t, IntTuple(1, -1, 2, 1, 1), st.Tuples[1])
}
