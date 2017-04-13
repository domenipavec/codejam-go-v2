package io

import (
	"log"
	"math/big"
	"strconv"

	"github.com/matematik7/codejam-go-v2/datastructures/grid"
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
	data := i.scanner.Bytes()
	data_copy := make([]byte, len(data))
	copy(data_copy, data)
	return data_copy
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

func (i *Input) SliceInt(n int) []int {
	ints := make([]int, 0, n)
	for j := 0; j < n; j++ {
		ints = append(ints, i.Int())
	}
	return ints
}

func (i *Input) SliceFloat(n int) []float64 {
	floats := make([]float64, 0, n)
	for j := 0; j < n; j++ {
		floats = append(floats, i.Float())
	}
	return floats
}

func (i *Input) SliceString(n int) []string {
	strs := make([]string, 0, n)
	for j := 0; j < n; j++ {
		strs = append(strs, i.String())
	}
	return strs
}

func (i *Input) SliceBytes(n int) [][]byte {
	sb := make([][]byte, n)
	for j := range sb {
		sb[j] = i.Bytes()
	}
	return sb
}

func (input *Input) GridInt(n, m int) grid.GridInt {
	newGrid := grid.NewInt(n, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			newGrid[i][j] = input.Int()
		}
	}
	return newGrid
}
