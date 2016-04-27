package io

import (
	"log"
	"math/big"
	"strconv"

	"github.com/matematik7/codejam-go/integer"
	"github.com/matematik7/codejam-go/st"
)

type InputProvider interface {
	Scan() bool
	Text() string
	Bytes() []byte
	Err() error
}

type Input struct {
	scanner InputProvider
	current []string
}

func newInput(ip InputProvider) *Input {
	return &Input{
		scanner: ip,
	}
}

func (i *Input) init() {
	i.current = i.current[:0]
}

func (i *Input) currentCase() []string {
	return i.current
}

func (i *Input) Scan() {
	if ok := i.scanner.Scan(); !ok {
		log.Fatalln("Error scanning input:", i.scanner.Err())
	}
	i.current = append(i.current, i.scanner.Text())
}

func (i *Input) String() string {
	i.Scan()
	return i.scanner.Text()
}

func (i *Input) Bytes() []byte {
	i.Scan()
	return i.scanner.Bytes()
}

func (i *Input) Int() int {
	n, err := strconv.Atoi(i.String())
	if err != nil {
		log.Fatalln("Error scanning for int:", err)
	}
	return n
}

func (i *Input) Float() float64 {
	f, err := strconv.ParseFloat(i.String(), 64)
	if err != nil {
		log.Fatalln("Error scanning for float:", err)
	}
	return f
}

func (i *Input) BigInt() *big.Int {
	n := &big.Int{}
	str := i.String()

	n, ok := n.SetString(str, 10)
	if !ok {
		log.Fatalln("Error scanning for big int:", str)
	}
	return n
}

func (i *Input) Digits() []int {
	str := i.String()
	ints := make([]int, 0, len(str))
	for _, chr := range str {
		if chr < 48 || chr > 57 {
			log.Fatalln("String element not a digit:", chr)
		}
		ints = append(ints, int(chr-48))
	}
	return ints
}

func (i *Input) SliceInt(n int) integer.Slice {
	ints := make([]int, 0, n)
	for j := 0; j < n; j++ {
		ints = append(ints, i.Int())
	}
	return ints
}

func (i *Input) SetInt(n int) *integer.Set {
	return integer.NewSet(i.SliceInt(n)...)
}

func (i *Input) MultiSetInt(n int) *integer.MultiSet {
	return integer.NewMultiSet(i.SliceInt(n)...)
}

func (i *Input) SliceTupleFromInts(n, m int) *st.SliceTuple {
	return st.FromInts(m, i.SliceInt(n*m)...)
}

func (i *Input) SliceTupleFromFloats(n, m int) *st.SliceTuple {
	return st.FromFloats(m, i.SliceFloat(n*m)...)
}

func (i *Input) SliceTupleFromStrings(n, m int) *st.SliceTuple {
	return st.FromStrings(m, i.SliceString(n*m)...)
}

func (i *Input) GridInt(y, x int) integer.Grid {
	grid := make([][]int, 0, y)
	for j := 0; j < y; j++ {
		grid = append(grid, i.SliceInt(x))
	}
	return grid
}

func (i *Input) SliceFloat(n int) []float64 {
	floats := make([]float64, 0, n)
	for j := 0; j < n; j++ {
		floats = append(floats, i.Float())
	}
	return floats
}

func (i *Input) GridFloat(y, x int) [][]float64 {
	grid := make([][]float64, 0, y)
	for j := 0; j < y; j++ {
		grid = append(grid, i.SliceFloat(x))
	}
	return grid
}

func (i *Input) SliceString(n int) []string {
	strs := make([]string, 0, n)
	for j := 0; j < n; j++ {
		strs = append(strs, i.String())
	}
	return strs
}
