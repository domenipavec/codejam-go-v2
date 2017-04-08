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
