package integer

import "container/heap"

type maxHeapTupleHelper MaxHeapTuple

func (mht maxHeapTupleHelper) Less(i, j int) bool {
	minlen := Min(len(mht[i]), len(mht[j]))
	for k := 0; k < minlen; k++ {
		if mht[i][k] != mht[j][k] {
			return mht[i][k] > mht[j][k]
		}
	}
	return false
}
func (mht maxHeapTupleHelper) Swap(i, j int) { mht[i], mht[j] = mht[j], mht[i] }
func (mht maxHeapTupleHelper) Len() int      { return len(mht) }

func (mht *maxHeapTupleHelper) Push(x interface{}) {
	*mht = append(*mht, x.([]int))
}

func (mht *maxHeapTupleHelper) Pop() interface{} {
	i := len(*mht) - 1
	value := (*mht)[i]
	*mht = (*mht)[:i]
	return value
}

type MaxHeapTuple SliceTuple

func NewMaxHeapTuple(st SliceTuple) MaxHeapTuple {
	mht := maxHeapTupleHelper(st)
	heap.Init(&mht)
	return MaxHeapTuple(mht)
}

func (mht MaxHeapTuple) Max() []int {
	return mht[0]
}

func (mht *MaxHeapTuple) FixMax() {
	heap.Fix((*maxHeapTupleHelper)(mht), 0)
}

func (mht *MaxHeapTuple) Push(as ...[]int) {
	for _, a := range as {
		heap.Push((*maxHeapTupleHelper)(mht), a)
	}
}

func (mht *MaxHeapTuple) Pop() []int {
	return heap.Pop((*maxHeapTupleHelper)(mht)).([]int)
}

func (mht MaxHeapTuple) Copy() MaxHeapTuple {
	return MaxHeapTuple(SliceTuple(mht).Copy())
}
