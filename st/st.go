package st

import (
	"bytes"
	"container/heap"
	"sort"
	"strconv"
)

type Tuple struct {
	Ints    []int
	Floats  []float64
	Strings []string
}

func IntTuple(ints ...int) *Tuple {
	t := &Tuple{
		Ints: make([]int, len(ints)),
	}
	copy(t.Ints, ints)
	return t
}

func FloatTuple(floats ...float64) *Tuple {
	t := &Tuple{
		Floats: make([]float64, len(floats)),
	}
	copy(t.Floats, floats)
	return t
}

func StringTuple(strings ...string) *Tuple {
	t := &Tuple{
		Strings: make([]string, len(strings)),
	}
	copy(t.Strings, strings)
	return t
}

func (t *Tuple) Copy() *Tuple {
	tnew := &Tuple{}
	if len(t.Ints) > 0 {
		tnew.Ints = make([]int, len(t.Ints))
		copy(tnew.Ints, t.Ints)
	}
	if len(t.Floats) > 0 {
		tnew.Floats = make([]float64, len(t.Floats))
		copy(tnew.Floats, t.Floats)
	}
	if len(t.Strings) > 0 {
		tnew.Strings = make([]string, len(t.Strings))
		copy(tnew.Strings, t.Strings)
	}
	return tnew
}

func (t Tuple) String() string {
	buffer := &bytes.Buffer{}
	i := 0
	for _, d := range t.Ints {
		if i != 0 {
			buffer.WriteByte(' ')
		}
		buffer.WriteString(strconv.Itoa(d))
		i++
	}
	for _, d := range t.Floats {
		if i != 0 {
			buffer.WriteByte(' ')
		}
		buffer.WriteString(strconv.FormatFloat(d, 'f', -1, 64))
		i++
	}
	for _, d := range t.Strings {
		if i != 0 {
			buffer.WriteByte(' ')
		}
		buffer.WriteString(d)
		i++
	}
	return buffer.String()
}

type Type int

const (
	Int Type = iota
	Float
	String
)

type Sorter struct {
	t       Type
	collumn int
	reverse bool
}

type SliceTuple struct {
	Tuples  []*Tuple
	sorters []Sorter
}

func NewSliceTuple(ts ...*Tuple) *SliceTuple {
	return &SliceTuple{
		Tuples: ts,
	}
}

func FromInts(c int, data ...int) *SliceTuple {
	st := NewSliceTuple()
	for i := 0; i < len(data); i += c {
		st.Append(IntTuple(data[i : i+c]...))
	}
	return st
}

func FromFloats(c int, data ...float64) *SliceTuple {
	st := NewSliceTuple()
	for i := 0; i < len(data); i += c {
		st.Append(FloatTuple(data[i : i+c]...))
	}
	return st
}

func FromStrings(c int, data ...string) *SliceTuple {
	st := NewSliceTuple()
	for i := 0; i < len(data); i += c {
		st.Append(StringTuple(data[i : i+c]...))
	}
	return st
}

func (st *SliceTuple) Copy() *SliceTuple {
	stnew := &SliceTuple{
		Tuples: make([]*Tuple, st.Len()),
	}

	for i, tuple := range st.Tuples {
		stnew.Tuples[i] = tuple.Copy()
	}

	return stnew
}

func (st *SliceTuple) CopySlice() *SliceTuple {
	stnew := &SliceTuple{
		Tuples: make([]*Tuple, st.Len()),
	}
	copy(stnew.Tuples, st.Tuples)
	return stnew
}

func (st *SliceTuple) Append(ts ...*Tuple) {
	st.Tuples = append(st.Tuples, ts...)
}

func (st *SliceTuple) Prepend(ts ...*Tuple) {
	st.Tuples = append(ts, st.Tuples...)
}

func (st *SliceTuple) Insert(i int, ts ...*Tuple) {
	st.Tuples = append(st.Tuples[:i], append(ts, st.Tuples[i:]...)...)
}

func (st *SliceTuple) Remove(i int) *Tuple {
	t := st.Tuples[i]
	st.Tuples = append(st.Tuples[:i], st.Tuples[i+1:]...)
	return t
}

func (st *SliceTuple) RemoveFirst() *Tuple {
	return st.Remove(0)
}

func (st *SliceTuple) RemoveLast() *Tuple {
	return st.Remove(st.Len() - 1)
}

func (st *SliceTuple) PrefixIntConst(cs ...int) {
	for _, tuple := range st.Tuples {
		tuple.Ints = append(append([]int{}, cs...), tuple.Ints...)
	}
}

func (st *SliceTuple) PrefixFloatConst(cs ...float64) {
	for _, tuple := range st.Tuples {
		tuple.Floats = append(append([]float64{}, cs...), tuple.Floats...)
	}
}

func (st *SliceTuple) PrefixStringConst(cs ...string) {
	for _, tuple := range st.Tuples {
		tuple.Strings = append(append([]string{}, cs...), tuple.Strings...)
	}
}

func (st *SliceTuple) PostfixIntConst(cs ...int) {
	for _, tuple := range st.Tuples {
		tuple.Ints = append(tuple.Ints, cs...)
	}
}

func (st *SliceTuple) PostfixFloatConst(cs ...float64) {
	for _, tuple := range st.Tuples {
		tuple.Floats = append(tuple.Floats, cs...)
	}
}

func (st *SliceTuple) PostfixStringConst(cs ...string) {
	for _, tuple := range st.Tuples {
		tuple.Strings = append(tuple.Strings, cs...)
	}
}

func (st *SliceTuple) PrefixIntIndex() {
	for i, tuple := range st.Tuples {
		tuple.Ints = append([]int{i}, tuple.Ints...)
	}
}

func (st *SliceTuple) PostfixIntIndex() {
	for i, tuple := range st.Tuples {
		tuple.Ints = append(tuple.Ints, i)
	}
}

func (st *SliceTuple) Reverse() {
	for i := st.Len()/2 - 1; i >= 0; i-- {
		opp := st.Len() - 1 - i
		st.Swap(i, opp)
	}
}

func (st *SliceTuple) Get(i int) *Tuple {
	return st.Tuples[i]
}

func (st *SliceTuple) First() *Tuple {
	return st.Tuples[0]
}

func (st *SliceTuple) Last() *Tuple {
	return st.Tuples[len(st.Tuples)-1]
}

func (st *SliceTuple) Less(i, j int) bool {
	it := st.Tuples[i]
	jt := st.Tuples[j]
	for _, sorter := range st.sorters {
		switch sorter.t {
		case Int:
			iint := it.Ints[sorter.collumn]
			jint := jt.Ints[sorter.collumn]
			if iint != jint {
				if sorter.reverse {
					return iint > jint
				} else {
					return iint < jint
				}
			}
		case Float:
			ifloat := it.Floats[sorter.collumn]
			jfloat := jt.Floats[sorter.collumn]
			if ifloat != jfloat {
				if sorter.reverse {
					return ifloat > jfloat
				} else {
					return ifloat < jfloat
				}
			}
		case String:
			istring := it.Strings[sorter.collumn]
			jstring := jt.Strings[sorter.collumn]
			if istring != jstring {
				if sorter.reverse {
					return istring > jstring
				} else {
					return istring < jstring
				}
			}
		}
	}
	return false
}
func (st *SliceTuple) Swap(i, j int) { st.Tuples[i], st.Tuples[j] = st.Tuples[j], st.Tuples[i] }
func (st *SliceTuple) Len() int      { return len(st.Tuples) }

func (st *SliceTuple) Sort() {
	sort.Sort(st)
}

func (st *SliceTuple) SortOrder(ss ...Sorter) {
	st.sorters = ss
}

func IntAsc(c int) Sorter {
	return Sorter{
		t:       Int,
		collumn: c,
		reverse: false,
	}
}

func IntDesc(c int) Sorter {
	return Sorter{
		t:       Int,
		collumn: c,
		reverse: true,
	}
}

func FloatAsc(c int) Sorter {
	return Sorter{
		t:       Float,
		collumn: c,
		reverse: false,
	}
}

func FloatDesc(c int) Sorter {
	return Sorter{
		t:       Float,
		collumn: c,
		reverse: true,
	}
}

func StringAsc(c int) Sorter {
	return Sorter{
		t:       String,
		collumn: c,
		reverse: false,
	}
}

func StringDesc(c int) Sorter {
	return Sorter{
		t:       String,
		collumn: c,
		reverse: true,
	}
}

func (st *SliceTuple) Push(x interface{}) {
	st.Append(x.(*Tuple))
}

func (st *SliceTuple) Pop() interface{} {
	return st.RemoveLast()
}

func (st *SliceTuple) HeapInit() {
	heap.Init(st)
}

func (st *SliceTuple) HeapPop() *Tuple {
	return heap.Pop(st).(*Tuple)
}

func (st *SliceTuple) HeapPush(t *Tuple) {
	heap.Push(st, t)
}

func (st *SliceTuple) HeapFix(i int) {
	heap.Fix(st, i)
}
