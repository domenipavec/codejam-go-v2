package io

import (
	"log"
	"os"
	"strings"
)

type TestCaseFunc func(*Input, *Output)

func TestCases(f TestCaseFunc) {
	if len(os.Args) < 2 {
		log.Fatalln("You need to specify at least one input file")
	}
	for _, inputFn := range os.Args[1:] {
		outputFn := strings.TrimSuffix(inputFn, ".in") + ".out"

		inputF, err := os.Open(inputFn)
		if err != nil {
			log.Fatalln("Error opening input file:", err)
		}
		outputF, err := os.Create(outputFn)
		if err != nil {
			log.Fatalln("Error creating output file:", err)
		}

		input := newInput(inputF)
		output := newOutput(outputF)

		T := input.Int()
		for i := 1; i <= T; i++ {
			f(input, output)
			output.flush(i)
		}

		if err = inputF.Close(); err != nil {
			log.Fatalln("Error closing input file:", err)
		}
		if err = outputF.Close(); err != nil {
			log.Fatalln("Error closing output file:", err)
		}
	}
}
