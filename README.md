# codejam-go [![Build Status](https://travis-ci.org/matematik7/codejam-go.svg?branch=master)](https://travis-ci.org/matematik7/codejam-go)

Some helper stuff for codejam competition in go. See example.


## input

Reads whitespace separated stuff from input file

- **input.String()** - string input
- **input.Bytes()** - []byte input
- **input.Int()** - int input (64-bit usually)
- **input.Float()** - float64 input
- **input.BigInt()** - \*big.Int input
- **input.Digits()** - split digits without space to []int slice
- **input.SliceInt(n)** - *n* ints to []int
- **input.MultiSetInt(n)** - *n* ints to integer.MultiSet
- **input.GridInt(y, x)** - *y* rows and *x* cols to [][]int, first is row index
- **input.SliceFloat(n)** - *n* floats to []float64
- **input.GridFloat(y, x)** - *y* rows and *x* cols to [][]float64, first is row index
- **input.SliceString(n)** - *n* strings to []string


## output

Writes output after **Case #n:**, beginning space and trailing newline are provided if necessary.

Output to solution file:
- **output.Print(...interface{})** - prints all, spaces are added between operands when neither is a string
- **output.Printf(format, ...interface{})** - prints with format string

Output to console:
- **output.DebugCase()** - prints case number, input and output
- **output.Debug(...interface{})** - first calls *DebugCase()*, then prints all, spaces are added between operands when neither is a string
- **output.Debugf(format, ...interface{})** - first calls *DebugCase()*, then prints with format string
- **output.Fatal(...interface{})** - same as *Debug()*, but terminates
- **output.Fatalf(format, ...interface{})** - same as *Debugf*, but terminates

Testing asserts:
(all asserts have optional fatal bool parameter, that can be set to terminate if mistake is encountered)
- **output.AssertByteCount(byte, count, fatal)** - check if output has *count* number of *byte*-s
- **output.AssertCount(count, fatal)** - check if output has *count* bytes


## integer

Useful function for integers

- **integer.Min(...int)** - returns minimal from the given ints
- **integer.Max(...int)** - returns maximal from the given ints
- **integer.Abs(a)** - returns absolute value of a
- **integer.Ceil(a, b)** - divides a by b and rounds the result up
- **integer.Range(max)** - python style range, return slice of ints from *0* to *max-1*
- **integer.Range(min,max)** - python style range, return slice of ints from *min* to *max-1*
- **integer.Range(min,max,step)** - python style range, return slice of ints from *min* to *max-1* with step spacing

### integer.MultiSet

Multiset implementation for integers (using map[int]int)

- **ms := integer.NewMultiSet(...int)** - construct MultiSet from given ints
- **ms.Contains(a)** - returns true if a is in MultiSet
- **ms.ContainsAll(...int)** - returns true if all given ints are in MultiSet
- **ms.ContainsAny(...int)** - returns true if any of given ints is in MultiSet
- **ms.Len()** - returns number of elements in MultiSet (this is O(n), where n is number of different elements)
- **ms.Count(a)** - returns number of a's in MultiSet
- **ms.Insert(...int)** - insert all given ints in MultiSet
- **ms.InsertN(a, n)** - insert n a's in MultiSet
- **ms.RemoveOne(...int)** - remove one of each given int from MultiSet
- **ms.RemoveAll(...int)** - remove all of each given int from MultiSet
- **ms.Clear()** - remove all elements from MultiSet
