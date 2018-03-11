package queue

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "ValueType=BUILTINS,slice.SliceInt,slice.SliceSliceInt,slice.SliceFloat64,slice.SliceSliceFloat64"

type ValueType generic.Type

type QueueValueType struct {
	data  []ValueType
	front int
	back  int
}

func NewValueType() QueueValueType {
	return QueueValueType{
		data:  make([]ValueType, 64),
		front: 0,
		back:  0,
	}
}

func (q *QueueValueType) inc(i int) int {
	return (i + 1) & (len(q.data) - 1)
}

func (q *QueueValueType) dec(i int) int {
	return (i - 1) & (len(q.data) - 1)
}

func (q QueueValueType) Len() int {
	return (q.back - q.front) & (len(q.data) - 1)
}

func (q *QueueValueType) growIfNeeded() {
	l := q.Len()
	if l < len(q.data)-1 {
		return
	}
	newData := make([]ValueType, 4*len(q.data))
	if q.front < q.back {
		copy(newData, q.data[q.front:q.back])
	} else {
		n := copy(newData, q.data[q.front:])
		copy(newData[n:], q.data[:q.back])
	}
	q.data = newData
	q.front = 0
	q.back = l
}

func (q *QueueValueType) Push(v ValueType) {
	q.growIfNeeded()
	q.data[q.back] = v
	q.back = q.inc(q.back)
}

func (q *QueueValueType) PushFront(v ValueType) {
	q.growIfNeeded()
	q.front = q.dec(q.front)
	q.data[q.front] = v
}

func (q QueueValueType) Front() ValueType {
	return q.data[q.front]
}

func (q QueueValueType) Back() ValueType {
	return q.data[q.dec(q.back)]
}

func (q *QueueValueType) Pop() ValueType {
	v := q.Front()
	q.front = q.inc(q.front)
	return v
}

func (q *QueueValueType) PopBack() ValueType {
	q.back = q.dec(q.back)
	return q.data[q.back]
}
