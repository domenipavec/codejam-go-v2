package integer

import "container/heap"

type minHeapTupleHelper MinHeapTuple

func (mht minHeapTupleHelper) Less(i, j int) bool {
	is := mht[i]
	js := mht[j]
	var minlen int
	if len(is) < len(js) {
		minlen = len(is)
	} else {
		minlen = len(js)
	}
	for k := 0; k < minlen; k++ {
		if is[k] != js[k] {
			return is[k] < js[k]
		}
	}
	return false
}
func (mht minHeapTupleHelper) Swap(i, j int) { mht[i], mht[j] = mht[j], mht[i] }
func (mht minHeapTupleHelper) Len() int      { return len(mht) }

func (mht *minHeapTupleHelper) Push(x interface{}) {
	*mht = append(*mht, x.([]int))
}

func (mht *minHeapTupleHelper) Pop() interface{} {
	i := len(*mht) - 1
	value := (*mht)[i]
	*mht = (*mht)[:i]
	return value
}

type MinHeapTuple SliceTuple

func NewMinHeapTuple(st SliceTuple) MinHeapTuple {
	mht := minHeapTupleHelper(st)
	heap.Init(&mht)
	return MinHeapTuple(mht)
}

func (mht MinHeapTuple) Min() []int {
	return mht[0]
}

func (mht *MinHeapTuple) FixMin() {
	heap.Fix((*minHeapTupleHelper)(mht), 0)
}

func (mht *MinHeapTuple) Push(as ...[]int) {
	for _, a := range as {
		heap.Push((*minHeapTupleHelper)(mht), a)
	}
}

func (mht *MinHeapTuple) Pop() []int {
	return heap.Pop((*minHeapTupleHelper)(mht)).([]int)
}

func (mht MinHeapTuple) Copy() MinHeapTuple {
	return MinHeapTuple(SliceTuple(mht).Copy())
}
