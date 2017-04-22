# codejam-go [![Build Status](https://travis-ci.org/matematik7/codejam-go-v2.svg?branch=master)](https://travis-ci.org/matematik7/codejam-go-v2)

Some helper stuff for codejam competition in go version 2. See example.

To enable pprof images install:
```
sudo apt-get install graphviz
```


## input

Reads whitespace separated stuff from input file

- **input.String()** - string input
- **input.Bytes()** - []byte input
- **input.Int()** - int input (64-bit usually)
- **input.Float()** - float64 input
- **input.BigInt()** - \*big.Int input
- **input.Digits()** - split digits without space to []int slice
- **input.SliceInt(n)** - *n* ints to []int
- **input.SliceFloat(n)** - *n* floats to []float64
- **input.SliceString(n)** - *n* strings to []string
- **input.SliceBytes(n)** - *n* []byte words to [][]byte


## output

Writes output after **Case #n:**, beginning space and trailing newline are provided if necessary.

Output to solution file:
- **output.Print(...interface{})** - prints all, spaces are added between operands when neither is a string
- **output.Println(...interface{})** - prints all, spaces are always added between operands and a newline is appended
- **output.Printf(format, ...interface{})** - prints with format string

Output to console:
- **output.DebugCase()** - prints case number, input and output
- **output.Debug(...interface{})** - first calls *DebugCase()*, then prints all, spaces are added between operands when neither is a string
- **output.Debugf(format, ...interface{})** - first calls *DebugCase()*, then prints with format string
- **output.Fatal(...interface{})** - same as *Debug()*, but terminates
- **output.Fatalf(format, ...interface{})** - same as *Debugf*, but terminates
- **output.Periodic(...interface{})** - prints all only every second, for fast loops, spaces are added between operands when neither is a string
- **output.Periodicf(...interface{})** - prints with format string only every second, for fast loops
- **output.PeriodicInt(a)** - prints integer *a* only every second, for very fast loops, avoids memory allocation for interface{}
- **ouptut.PeriodicCount()** - increases internal count every call, prints only every second

Output to chart (draws a png chart per testcase):
- **output.Point(x, y)** - chart point with float64 coordinates
- **output.PointInt(x,y)** - chart point with int coordinates

Testing asserts:
(all asserts have optional fatal bool parameter, that can be set to terminate if mistake is encountered)
- **output.AssertByteCount(byte, count, fatal)** - check if output has *count* number of *byte*-s
- **output.AssertCount(count, fatal)** - check if output has *count* bytes
- **output.AssertEqual(data, fatal)** - check if output is equal to data
- **output.AssertIntEqual(a, b, fatal)** - check if a and b are equal ints
- **output.AssertTrue(a, fatal)** - check if a is true

Timers:
- **output.TimerStart(key)** - start timer *key*
- **output.TimerStop(key)** - stop timer *key*

## integer

Useful function for integers

- **integer.MAX** - maximum int can hold
- **integer.MIN** - minimum int can hold
- **integer.Min(...int)** - returns minimal from the given ints
- **integer.Max(...int)** - returns maximal from the given ints
- **integer.Abs(a)** - returns absolute value of a
- **integer.Ceil(a, b)** - divides a by b and rounds the result up
- **integer.Range(max)** - python style range, return slice of ints from *0* to *max-1*
- **integer.Range(min,max)** - python style range, return slice of ints from *min* to *max-1*
- **integer.Range(min,max,step)** - python style range, return slice of ints from *min* to *max-1* with step spacing
- **integer.Gcd(...int)** - return greatest common divider of given ints
- **integer.Lcm(...int)** - return least common multiple of given ints
- **integer.Pow(a,b)** - returns a to the power b
- **integer.Log10(a)** - returns log base 10 of a
- **integer.Log2(a)** - returns log base 2 of a
- **integer.Round(f)** - round float to int
- **integer.Floor(f)** - floor float to int
- **integer.Ceil(f)** - ceil float to int

## Set

- **set.NewValueType()** - create new set for ValueType
- **s.Add(...v)** - add v(s) to set
- **s.Remove(...v)** - remove v(s) from set
- **s.Len()** - number of elements in set
- **s.Contains(element)** - is element in set

## MultiSet

- **multiset.NewValueType()** - create new multi set for ValueType
- **s.Add(...v)** - add v(s) to set
- **s.AddN(n, ...v)** - add n instances of v(s) to set
- **s.Remove(...v)** - remove v(s) from set
- **s.RemoveN(n, ...v)** - remove n instances of v(s) from set
- **s.Len()** - number of elements in set
- **s.Count(element)** - count of *element* in set

## Slice

- **slice.NewSliceValueType()** - create new slice of ValueType
- **s.Copy()** - new independent copy of slice
- **s.SpiralIterator()** - returns coordinates in spiral for 2d slice
- **s.Set(v)** - set all elements to v
- **s.SortAsc()** - sort ascending
- **s.SortDesc()** - sort descending
- **s.Len()** - length of slice
- **s.Get(i)** - get i-th element with negative index support
- **s.Swap(i, j)** - swap i-th and j-th elements
- **s.Push(interface)** - cast interface to value and append
- **s.Pop()** - get last element as interface and remove it
- **s.Append(...v)** - append v-s
- **s.Prepend(...v)** - prepend v-s
- **s.Print(sep)** - print using separator
- **s.MinHeap()** - returns minheap
- **s.MaxHeap()** - returns maxheap

### MaxHeap
(based on slice)

- **hp.Slice** - underlying slice
- **hp.Min()** - get min value
- **hp.Fix(i)** - fix i-th after it was changed
- **hp.Pop()** - remove and return minimum element
- **hp.Push(e)** - add e to heap
- **hp.Remove(i)** - remove i-th element from heap

### MinHeap
(based on slice)

- **hp.Slice** - underlying slice
- **hp.Max()** - get max value
- **hp.Fix(i)** - fix i-th after it was changed
- **hp.Pop()** - remove and return minimum element
- **hp.Push(e)** - add e to heap
- **hp.Remove(i)** - remove i-th element from heap

## Left leaning red black binary search tree

- **llrb.NewValueType()** - get llrb of ValueType
- **l.Len()** - number of elements in tree
- **l.Insert(v)** - insert v to tree
- **l.Delete(v)** - delete v from tree
- **l.Min()** - get min v from tree
- **l.Max()** - get max v from tree
- **l.Search(v)** - is v in tree
- **l.Get(i)** - get i-th smallest v
- **l.Rank(v)** - get rank of v in tree
- **l.IterateAsc(cb)** - call cb callback with vs in ascending order
- **l.IterateDesc(cb)** - call cb callback with vs in descending order
