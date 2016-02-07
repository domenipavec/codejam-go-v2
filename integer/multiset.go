package integer

import "log"

type MultiSet map[int]int

func NewMultiSet(as ...int) MultiSet {
	ms := MultiSet(make(map[int]int))
	ms.Insert(as...)
	return ms
}

func (ms MultiSet) Copy() MultiSet {
	copyMs := MultiSet(make(map[int]int))
	for k, v := range ms {
		copyMs[k] = v
	}
	return copyMs
}

func (ms MultiSet) Contains(a int) bool {
	return ms[a] > 0
}

func (ms MultiSet) ContainsAll(as ...int) bool {
	for _, a := range as {
		if ms[a] <= 0 {
			return false
		}
	}
	return true
}

func (ms MultiSet) ContainsAny(as ...int) bool {
	for _, a := range as {
		if ms[a] > 0 {
			return true
		}
	}
	return false
}

func (ms MultiSet) Len() int {
	l := 0
	for _, i := range ms {
		l += i
	}
	return l
}

func (ms MultiSet) Count(a int) int {
	return ms[a]
}

func (ms MultiSet) Insert(as ...int) {
	for _, a := range as {
		ms[a] += 1
	}
}

func (ms MultiSet) InsertN(a, n int) {
	ms[a] += n
}

func (ms MultiSet) RemoveOne(as ...int) {
	for _, a := range as {
		if ms[a] <= 0 {
			log.Fatalln("Nothing to remove when removing:", a)
		}
		ms[a] -= 1
		if ms[a] == 0 {
			delete(ms, a)
		}
	}
}

func (ms MultiSet) RemoveAll(as ...int) {
	for _, a := range as {
		if ms[a] <= 0 {
			log.Fatalln("Nothing to remove when removing:", a)
		}
		delete(ms, a)
	}
}

func (ms *MultiSet) Clear() {
	*ms = make(map[int]int)
}
