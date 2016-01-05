package io

import (
	"bufio"
	"bytes"
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
	correctFn := strings.TrimSuffix(inputFn, ".in") + ".correct"

	outputBuffer := &bytes.Buffer{}

	input := NewBufferedInput(parser.inputData(inputFn))
	output := newOutput(outputBuffer)

	for i := 1; i <= input.T; i++ {
		parser.f(input.GetInput(i), output)
		output.flush(i)
	}

	if _, err := os.Stat(correctFn); err == nil {
		CompareOutput(correctFn, bytes.NewReader(outputBuffer.Bytes()), input)
	}

	parser.writeOutput(outputFn, outputBuffer.Bytes())
}

func (parser *Parser) inputData(inputFn string) []string {
	data := []string{}

	inputF, err := os.Open(inputFn)
	if err != nil {
		log.Fatalln("Error opening input file:", err)
	}
	defer inputF.Close()

	scanner := bufio.NewScanner(inputF)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln("Error reading input:", err)
	}

	return data
}

func (parser *Parser) writeOutput(outputFn string, data []byte) {
	outputF, err := os.Create(outputFn)
	if err != nil {
		log.Fatalln("Error creating output file:", err)
	}
	defer outputF.Close()

	outputF.Write(data)
}
