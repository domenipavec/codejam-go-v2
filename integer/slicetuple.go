package integer

import "sort"

type SliceTuple []*[]int

func Tuple(a ...int) []int {
	return a
}

func NewSliceTuple(as ...int) SliceTuple {
	var st SliceTuple
	for _, a := range as {
		st.Append(Tuple(a))
	}
	return st
}

func (st SliceTuple) Get(i int) []int {
	return *st[i]
}

func (st *SliceTuple) Prepend(as ...[]int) {
	for _, a := range as {
		*st = append([]*[]int{&a}, *st...)
	}
}

func (st *SliceTuple) Append(as ...[]int) {
	for _, a := range as {
		*st = append(*st, &a)
	}
}

func (st *SliceTuple) Insert(i int, as ...[]int) {
	for _, a := range as {
		*st = append((*st)[:i], append([]*[]int{&a}, (*st)[i:]...)...)
		i++
	}
}

func (st SliceTuple) PrefixConst(a ...int) {
	for index := range st {
		*st[index] = append(a, *st[index]...)
	}
}

func (st SliceTuple) PrefixIndex() {
	for index := range st {
		*st[index] = append([]int{index}, *st[index]...)
	}
}

func (st SliceTuple) PostfixConst(a ...int) {
	for index := range st {
		*st[index] = append(*st[index], a...)
	}
}

func (st SliceTuple) PostfixIndex() {
	for index := range st {
		*st[index] = append(*st[index], index)
	}
}

func (st *SliceTuple) Delete(i int) []int {
	value := *(*st)[i]
	*st = append((*st)[:i], (*st)[i+1:]...)
	return value
}

func (st *SliceTuple) DeleteFirst() []int {
	return st.Delete(0)
}

func (st *SliceTuple) DeleteLast() []int {
	return st.Delete(len(*st) - 1)
}

func (st SliceTuple) Copy() SliceTuple {
	cp := make([]*[]int, len(st))
	for index, value := range st {
		newValue := make([]int, len(*value))
		copy(newValue, *value)
		cp[index] = &newValue
	}
	return cp
}

func (st SliceTuple) CopySlice() SliceTuple {
	cp := make([]*[]int, len(st))
	copy(cp, st)
	return cp
}

func (st SliceTuple) Swap(i, j int) {
	st[i], st[j] = st[j], st[i]
}

func (st SliceTuple) Reverse() {
	for i := len(st)/2 - 1; i >= 0; i-- {
		opp := len(st) - 1 - i
		st.Swap(i, opp)
	}
}

func (st SliceTuple) Len() int {
	return len(st)
}

type SliceTupleAsc SliceTuple

func (sta SliceTupleAsc) Less(i, j int) bool {
	is := *sta[i]
	js := *sta[j]
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
func (st SliceTupleAsc) Swap(i, j int) { st[i], st[j] = st[j], st[i] }
func (st SliceTupleAsc) Len() int      { return len(st) }

func (st SliceTuple) SortAsc() {
	sort.Sort(SliceTupleAsc(st))
}

type SliceTupleDesc SliceTuple

func (sta SliceTupleDesc) Less(i, j int) bool {
	is := *sta[i]
	js := *sta[j]
	var minlen int
	if len(is) < len(js) {
		minlen = len(is)
	} else {
		minlen = len(js)
	}
	for k := 0; k < minlen; k++ {
		if is[k] != js[k] {
			return is[k] > js[k]
		}
	}
	return false
}
func (st SliceTupleDesc) Swap(i, j int) { st[i], st[j] = st[j], st[i] }
func (st SliceTupleDesc) Len() int      { return len(st) }

func (st SliceTuple) SortDesc() {
	sort.Sort(SliceTupleDesc(st))
}

type SliceTupleAscBy struct {
	st SliceTuple
	n  int
}

func (sta SliceTupleAscBy) Less(i, j int) bool {
	return (*sta.st[i])[sta.n] < (*sta.st[j])[sta.n]
}
func (sta SliceTupleAscBy) Swap(i, j int) { sta.st[i], sta.st[j] = sta.st[j], sta.st[i] }
func (sta SliceTupleAscBy) Len() int      { return len(sta.st) }

func (st SliceTuple) SortAscBy(n int) {
	sta := SliceTupleAscBy{
		st: st,
		n:  n,
	}
	sort.Sort(sta)
}

type SliceTupleDescBy struct {
	st SliceTuple
	n  int
}

func (sta SliceTupleDescBy) Less(i, j int) bool {
	return (*sta.st[i])[sta.n] > (*sta.st[j])[sta.n]
}
func (sta SliceTupleDescBy) Swap(i, j int) { sta.st[i], sta.st[j] = sta.st[j], sta.st[i] }
func (sta SliceTupleDescBy) Len() int      { return len(sta.st) }

func (st SliceTuple) SortDescBy(n int) {
	sta := SliceTupleDescBy{
		st: st,
		n:  n,
	}
	sort.Sort(sta)
}
