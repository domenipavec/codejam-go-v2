package integer

import "log"

type MultiSet struct {
	Data   map[int]int
	length int
}

func NewMultiSet(as ...int) *MultiSet {
	ms := &MultiSet{
		Data:   make(map[int]int),
		length: 0,
	}
	ms.Insert(as...)
	return ms
}

func (ms *MultiSet) Copy() *MultiSet {
	copyMs := &MultiSet{
		Data:   make(map[int]int),
		length: ms.length,
	}
	for k, v := range ms.Data {
		copyMs.Data[k] = v
	}
	return copyMs
}

func (ms *MultiSet) Contains(a int) bool {
	return ms.Data[a] > 0
}

func (ms *MultiSet) ContainsAll(as ...int) bool {
	for _, a := range as {
		if ms.Data[a] <= 0 {
			return false
		}
	}
	return true
}

func (ms *MultiSet) ContainsAny(as ...int) bool {
	for _, a := range as {
		if ms.Data[a] > 0 {
			return true
		}
	}
	return false
}

func (ms *MultiSet) Len() int {
	return ms.length
}

func (ms *MultiSet) Count(a int) int {
	return ms.Data[a]
}

func (ms *MultiSet) Insert(as ...int) {
	ms.InsertN(1, as...)
}

func (ms *MultiSet) InsertN(n int, as ...int) {
	for _, a := range as {
		ms.Data[a] += n
	}
	ms.length += n * len(as)
}

func (ms *MultiSet) RemoveOne(as ...int) {
	ms.RemoveN(1, as...)
}

func (ms *MultiSet) RemoveN(n int, as ...int) {
	for _, a := range as {
		if ms.Data[a] < n {
			log.Fatalf("Not enough %d in multiset to remove %d.", a, n)
		}
		ms.Data[a] -= n
		if ms.Data[a] == 0 {
			delete(ms.Data, a)
		}
	}
	ms.length -= n * len(as)
}

func (ms *MultiSet) RemoveAll(as ...int) {
	for _, a := range as {
		if ms.Data[a] <= 0 {
			log.Fatalln("Nothing to remove when removing:", a)
		}
		ms.length -= ms.Data[a]
		delete(ms.Data, a)
	}
}

func (ms *MultiSet) Clear() {
	ms.Data = make(map[int]int)
	ms.length = 0
}
