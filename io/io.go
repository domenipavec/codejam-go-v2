package io

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type TestCaseFunc func(*Input, *Output)

type Parser struct {
	f             TestCaseFunc
	input         *Input
	output        *Output
	compareOutput *CompareOutput

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

func (parser *Parser) formatDuration(d int64) string {
	var i int
	df := float64(d)
	units := []string{"ns", "us", "ms", "s"}
	for i = 0; df >= 1000; i++ {
		df /= 1000
		if i >= 3 {
			break
		}
	}
	res := fmt.Sprintf("%.2f%s", df, units[i])
	return res
}

func (parser *Parser) ParseFile() {
	inputF, err := os.Open(parser.inputFn)
	if err != nil {
		log.Fatalln("Error opening input file:", err)
	}
	defer inputF.Close()

	outputF, err := os.Create(parser.outputFn)
	if err != nil {
		log.Fatalln("Error creating output file:", err)
	}
	defer outputF.Close()

	scanner := bufio.NewScanner(inputF)
	scanner.Split(bufio.ScanWords)

	parser.output = newOutput(outputF)
	parser.input = newInput(scanner)

	parser.compareOutput = nil
	if _, err := os.Stat(parser.correctFn); err == nil {
		parser.compareOutput = NewCompareOutput(parser.correctFn)
	}

	T := parser.input.Int()

	startTime := time.Now().UnixNano()
	for i := 1; i <= T; i++ {
		parser.runTestCase(i)
	}
	log.Println("Total time:", parser.formatDuration(time.Now().UnixNano()-startTime))
}

func (parser *Parser) runTestCase(i int) {
	warningTimer := time.NewTimer(500 * time.Millisecond)
	doneChan := make(chan bool)

	go func() {
		parser.output.init(parser.input, i)
		parser.input.init()

		parser.f(parser.input, parser.output)

		if parser.compareOutput != nil && parser.compareOutput.HasOutput(i) {
			parser.output.AssertEqual(string(parser.compareOutput.GetOutput(i)))
		}

		parser.output.flush()
		doneChan <- true
	}()

	select {
	case <-warningTimer.C:
		parser.output.Debug("Long calculation")
		<-doneChan
	case <-doneChan:
	}
}
