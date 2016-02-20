package integer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxHeapTuple(t *testing.T) {
	mht := NewMaxHeapTuple(NewSliceTuple(3, 2, 5))

	assert.Equal(t, Tuple(5), mht.Max())

	mht.Push(Tuple(6))

	assert.Equal(t, Tuple(6), mht.Max())

	assert.Equal(t, Tuple(6), mht.Pop())

	assert.Equal(t, Tuple(5), mht.Max())

	mht.Max()[0] = 1
	mht.FixMax()

	assert.Equal(t, Tuple(3), mht.Pop())
	assert.Equal(t, Tuple(2), mht.Pop())
	assert.Equal(t, Tuple(1), mht.Pop())
}
