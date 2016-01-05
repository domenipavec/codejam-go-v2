package io

import (
	"bufio"
	"io"
	"log"
	"strconv"
)

type Input struct {
	scanner *bufio.Scanner
}

func newInput(r io.Reader) *Input {
	input := &Input{
		scanner: bufio.NewScanner(r),
	}

	input.scanner.Split(bufio.ScanWords)

	return input
}

func (i *Input) Scan() {
	if ok := i.scanner.Scan(); !ok {
		log.Fatalln("Error scanning input:", i.scanner.Err())
	}
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
