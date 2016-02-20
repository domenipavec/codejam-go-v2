package integer

import "sort"

type SliceTuple [][]int

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

func (st *SliceTuple) Prepend(a []int) {
	*st = append([][]int{a}, *st...)
}

func (st *SliceTuple) Append(a []int) {
	*st = append(*st, a)
}

func (st *SliceTuple) Insert(i int, a []int) {
	*st = append((*st)[:i], append([][]int{a}, (*st)[i:]...)...)
}

func (st SliceTuple) PrefixConst(a ...int) {
	for index := range st {
		st[index] = append(a, st[index]...)
	}
}

func (st SliceTuple) PrefixIndex() {
	for index := range st {
		st[index] = append([]int{index}, st[index]...)
	}
}

func (st SliceTuple) PostfixConst(a ...int) {
	for index := range st {
		st[index] = append(st[index], a...)
	}
}

func (st SliceTuple) PostfixIndex() {
	for index := range st {
		st[index] = append(st[index], index)
	}
}

func (st *SliceTuple) Delete(i int) []int {
	value := (*st)[i]
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
	cp := make([][]int, len(st))
	for index, value := range st {
		cp[index] = make([]int, len(value))
		copy(cp[index], value)
	}
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
	minlen := Min(len(sta[i]), len(sta[j]))
	for k := 0; k < minlen; k++ {
		if sta[i][k] != sta[j][k] {
			return sta[i][k] < sta[j][k]
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
	minlen := Min(len(sta[i]), len(sta[j]))
	for k := 0; k < minlen; k++ {
		if sta[i][k] != sta[j][k] {
			return sta[i][k] > sta[j][k]
		}
	}
	return false
}
func (st SliceTupleDesc) Swap(i, j int) { st[i], st[j] = st[j], st[i] }
func (st SliceTupleDesc) Len() int      { return len(st) }

func (st SliceTuple) SortDesc() {
	sort.Sort(SliceTupleDesc(st))
}
