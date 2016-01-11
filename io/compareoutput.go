package io

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/matematik7/codejam-go/integer"
)

func CompareOutput(correctFn string, output io.Reader, input *BufferedInput) {
	msgs := []string{}

	correctF, err := os.Open(correctFn)
	if err != nil {
		log.Fatalln("Error opening correct file:", err)
	}
	defer correctF.Close()

	correctScanner := bufio.NewScanner(correctF)
	outputScanner := bufio.NewScanner(output)

	for i := 1; i <= input.T; i++ {
		if !correctScanner.Scan() {
			log.Fatalln("Error scanning correct:", correctScanner.Err())
		}
		if !outputScanner.Scan() {
			log.Fatalln("Error scanning output:", outputScanner.Err())
		}

		if correctScanner.Text() != outputScanner.Text() {
			caseString := fmt.Sprintf("Case #%d: ", i)
			correctString := strings.TrimPrefix(correctScanner.Text(), caseString)
			outputString := strings.TrimPrefix(outputScanner.Text(), caseString)
			msg := fmt.Sprint(caseString, "correct: ", correctString, ", output: ", outputString, ", input: ", input.InputProviders[i-1].Data)

			msgs = append(msgs, msg)
		}
	}

	if len(msgs) > 0 {
		log.Printf("%d differences comapred with %s:\n", len(msgs), correctFn)
		for _, msg := range msgs {
			log.Println(msg[:integer.Min(len(msg), 100)])
		}
	}
}
