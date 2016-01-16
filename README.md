# codejam-go

Some helper stuff for codejam competition in go. See example.

## Input

Reads whitespace separated stuff from input file

- **input.String()** - string input
- **input.Bytes()** - []byte input
- **input.Int()** - int input (64-bit usually)
- **input.Float()** - float64 input
- **input.Digits()** - split digits without space to []int slice
- **input.SliceInt(n)** - *n* ints to []int
- **input.GridInt(y, x)** - *y* rows and *x* cols to [][]int, first is row index
- **input.SliceFloat(n)** - *n* floats to []float64
- **input.GridFloat(y, x)** - *y* rows and *x* cols to [][]float64, first is row index


## Output

Writes output after **Case #n:**, beginning space and trailing newline are provided if necessary.

- **output.Print(...interface{})** - prints all, spaces are added between operands when neither is a string
- **output.Printf(format, ...interface{})** - prints with format string
- **output.AssertByteCount(byte, count)** - check if output has *count* number of *byte*-s
- **output.AssertCount(count)** - check if output has *count* bytes
