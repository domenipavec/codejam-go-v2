package io

import (
	"bufio"
	"math/big"
	"strings"
	"testing"

	"github.com/matematik7/codejam-go/integer"
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
	assert.Equal(t, 1.0, i.Float())
	assert.Equal(t, -123.1, i.Float())
	assert.Equal(t, 1e15, i.Float())
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
	assert.Equal(t, []int{1, -123, 1000000}, i.SliceInt(3))
}

func TestSetInt(t *testing.T) {
	i := initInput("1 -123\n1000000")
	ms := i.SetInt(3)
	assert.True(t, ms.ContainsAll(1, -123, 1000000))
}

func TestMultiSetInt(t *testing.T) {
	i := initInput("1 -123\n1000000")
	ms := i.MultiSetInt(3)
	assert.True(t, ms.ContainsAll(1, -123, 1000000))
}

func TestSliceTuple(t *testing.T) {
	i := initInput("1 -123\n1000000")
	st := i.SliceTuple(3)
	assert.Equal(t, 3, len(st))
	assert.Equal(t, integer.Tuple(1), st.Get(0))
	assert.Equal(t, integer.Tuple(-123), st.Get(1))
	assert.Equal(t, integer.Tuple(1000000), st.Get(2))
}

func TestGridInt(t *testing.T) {
	i := initInput("1 2 3\n4 5 6")
	assert.Equal(t, [][]int{[]int{1, 2, 3}, []int{4, 5, 6}}, i.GridInt(2, 3))
}

func TestSliceFloat(t *testing.T) {
	i := initInput("1.0 -123.1\n1e15")
	assert.Equal(t, []float64{1.0, -123.1, 1e15}, i.SliceFloat(3))
}

func TestGridFloat(t *testing.T) {
	i := initInput("1.1 2.2 3.3\n4.4 5.5 6.6")
	assert.Equal(t, [][]float64{[]float64{1.1, 2.2, 3.3}, []float64{4.4, 5.5, 6.6}}, i.GridFloat(2, 3))
}

func TestSliceString(t *testing.T) {
	i := initInput("input1 input2\ninput3")
	assert.Equal(t, []string{"input1", "input2", "input3"}, i.SliceString(3))
}
