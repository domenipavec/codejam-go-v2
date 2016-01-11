package io

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strings"
	"time"
)

type TestCaseFunc func(*Input, *Output)

type Parser struct {
	f      TestCaseFunc
	input  *BufferedInput
	output *Output

	inputFn   string
	outputFn  string
	correctFn string
}

func TestCases(f TestCaseFunc) {
	log.SetFlags(0)

	parser := Parser{
		f: f,
	}

	if len(os.Args) < 2 {
		log.Fatalln("You need to specify at least one input file")
	}
	for _, inputFn := range os.Args[1:] {
		parser.SetFn(inputFn)
		parser.ParseFile()
	}
}

func (parser *Parser) SetFn(inputFn string) {
	parser.inputFn = inputFn
	baseFn := strings.TrimSuffix(inputFn, ".in")
	parser.outputFn = baseFn + ".out"
	parser.correctFn = baseFn + ".correct"
}

func (parser *Parser) ParseFile() {
	outputBuffer := &bytes.Buffer{}

	parser.inputData()
	parser.output = newOutput(outputBuffer)

	for i := 1; i <= parser.input.T; i++ {
		parser.runTestCase(i)
	}

	if _, err := os.Stat(parser.correctFn); err == nil {
		CompareOutput(parser.correctFn, bytes.NewReader(outputBuffer.Bytes()), parser.input)
	}

	parser.writeOutput(outputBuffer.Bytes())
}

func (parser *Parser) runTestCase(i int) {
	warningTimer := time.NewTimer(500 * time.Millisecond)
	doneChan := make(chan bool)

	go func() {
		parser.f(parser.input.GetInput(i), parser.output)
		parser.output.flush(i)
		doneChan <- true
	}()

	select {
	case <-warningTimer.C:
		log.Printf("Long calculation #%d, input: %v\n", i, parser.input.InputProviders[i-1].Data)
		<-doneChan
	case <-doneChan:
	}
}

func (parser *Parser) inputData() {
	data := []string{}

	inputF, err := os.Open(parser.inputFn)
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

	parser.input = NewBufferedInput(data)
}

func (parser *Parser) writeOutput(data []byte) {
	outputF, err := os.Create(parser.outputFn)
	if err != nil {
		log.Fatalln("Error creating output file:", err)
	}
	defer outputF.Close()

	outputF.Write(data)
}
