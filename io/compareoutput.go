package io

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"strconv"
)

type CompareOutput struct {
	outputs map[int][]byte
}

func NewCompareOutput(correctF io.Reader) *CompareOutput {
	co := &CompareOutput{
		outputs: make(map[int][]byte),
	}

	correctData, err := ioutil.ReadAll(correctF)
	if err != nil {
		log.Fatalln("Error opening correct file:", err)
	}

	casesData := bytes.Split(correctData, []byte("Case #"))
	for _, caseData := range casesData {
		if len(caseData) == 0 {
			continue
		}

		colonIndex := bytes.Index(caseData, []byte(":"))
		if colonIndex == -1 {
			log.Fatalln("No colon in correct case:", caseData)
		}

		caseNumber, err := strconv.Atoi(string(caseData[:colonIndex]))
		if err != nil {
			log.Fatalln("Invalid case number in correct file:", string(caseData[:colonIndex]))
		}

		co.outputs[caseNumber] = caseData[colonIndex+1:]
	}

	return co
}

func (co *CompareOutput) HasOutput(i int) bool {
	_, ok := co.outputs[i]
	return ok
}

func (co *CompareOutput) GetOutput(i int) []byte {
	return co.outputs[i]
}
