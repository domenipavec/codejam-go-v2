package io

import (
	"bufio"
	"math/big"
	"strings"
	"testing"

	"github.com/matematik7/codejam-go-v2/datastructures/slice"
	"github.com/stretchr/testify/assert"
)

func initInput(s string) *Input {
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	return newInput(scanner)
}

func TestString(t *testing.T) {
	i := initInput("input1 input2\ninput3")
	assert.Equal(t, "input1", i.String())
	assert.Equal(t, "input2", i.String())
	assert.Equal(t, "input3", i.String())
}

func TestBytes(t *testing.T) {
	i := initInput("input1 input2\ninput3")
	assert.Equal(t, []byte("input1"), i.Bytes())
	assert.Equal(t, []byte("input2"), i.Bytes())
	assert.Equal(t, []byte("input3"), i.Bytes())
}

func TestInt(t *testing.T) {
	i := initInput("1 -123\n1000000")
	assert.Equal(t, 1, i.Int())
	assert.Equal(t, -123, i.Int())
	assert.Equal(t, 1000000, i.Int())
}

func TestFloat(t *testing.T) {
	i := initInput("1.0 -123.1\n1e15")
	assert.Equal(t, 1.0, i.Float64())
	assert.Equal(t, -123.1, i.Float64())
	assert.Equal(t, 1e15, i.Float64())
}

func TestBigInt(t *testing.T) {
	i := initInput("1 -123\n1000000")
	assert.Equal(t, big.NewInt(1), i.BigInt())
	assert.Equal(t, big.NewInt(-123), i.BigInt())
	assert.Equal(t, big.NewInt(1000000), i.BigInt())
}

func TestDigits(t *testing.T) {
	i := initInput("1234")
	assert.Equal(t, []int{1, 2, 3, 4}, i.Digits())
}

func TestSliceInt(t *testing.T) {
	i := initInput("1 -123\n1000000")
	assert.Equal(t, slice.SliceInt([]int{1, -123, 1000000}), i.SliceInt(3))
}

func TestSliceFloat(t *testing.T) {
	i := initInput("1.0 -123.1\n1e15")
	assert.Equal(t, slice.SliceFloat64([]float64{1.0, -123.1, 1e15}), i.SliceFloat64(3))
}

func TestSliceString(t *testing.T) {
	i := initInput("input1 input2\ninput3")
	assert.Equal(t, slice.SliceString([]string{"input1", "input2", "input3"}), i.SliceString(3))
}
