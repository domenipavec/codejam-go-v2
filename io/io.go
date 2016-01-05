package io

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type TestCaseFunc func(*Input, *Output)

type Parser struct {
	f TestCaseFunc
}

func TestCases(f TestCaseFunc) {
	parser := Parser{
		f: f,
	}

	if len(os.Args) < 2 {
		log.Fatalln("You need to specify at least one input file")
	}
	for _, inputFn := range os.Args[1:] {
		parser.ParseFile(inputFn)
	}
}

func (parser *Parser) ParseFile(inputFn string) {
	outputFn := strings.TrimSuffix(inputFn, ".in") + ".out"

	inputF, err := os.Open(inputFn)
	if err != nil {
		log.Fatalln("Error opening input file:", err)
	}
	defer inputF.Close()

	outputF, err := os.Create(outputFn)
	if err != nil {
		log.Fatalln("Error creating output file:", err)
	}
	defer outputF.Close()

	inputScanner := bufio.NewScanner(inputF)
	inputScanner.Split(bufio.ScanWords)

	input := newInput(inputScanner)
	output := newOutput(outputF)

	T := input.Int()
	for i := 1; i <= T; i++ {
		parser.f(input, output)
		output.flush(i)
	}
}
