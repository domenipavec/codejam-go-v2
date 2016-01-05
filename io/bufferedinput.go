package io

import (
	"log"
	"strconv"
)

type BufferedInput struct {
	T              int
	InputProviders []*bufferedInputProvider
}

func NewBufferedInput(data []string) *BufferedInput {
	T, err := strconv.Atoi(data[0])
	if err != nil {
		log.Fatalln("Error parsing number of cases:", err)
	}

	data = data[1:]

	bi := &BufferedInput{
		T:              T,
		InputProviders: make([]*bufferedInputProvider, T),
	}

	N := len(data) / bi.T
	for i := range bi.InputProviders {
		bi.InputProviders[i] = &bufferedInputProvider{
			i:    -1,
			Data: data[N*i : N*(i+1)],
		}
	}

	return bi
}

func (bi *BufferedInput) GetInput(i int) *Input {
	return newInput(bi.InputProviders[i-1])
}

type bufferedInputProvider struct {
	i    int
	Data []string
}

func (b *bufferedInputProvider) Scan() bool {
	b.i++
	return b.i < len(b.Data)
}

func (b *bufferedInputProvider) Text() string {
	if b.i == -1 {
		log.Fatalln("Need to call Scan first.")
	}
	return b.Data[b.i]
}

func (b *bufferedInputProvider) Bytes() []byte {
	return []byte(b.Text())
}

func (b *bufferedInputProvider) Err() error {
	return nil
}
