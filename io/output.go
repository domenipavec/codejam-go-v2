package io

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"unicode"
)

type Output struct {
	w io.Writer

	caseN int
	bip   *bufferedInputProvider

	output *bytes.Buffer
}

func newOutput(w io.Writer) *Output {
	return &Output{
		w:      w,
		output: &bytes.Buffer{},
	}
}

func (o *Output) Print(a ...interface{}) {
	fmt.Fprint(o.output, a...)
}

func (o *Output) Printf(format string, a ...interface{}) {
	fmt.Fprintf(o.output, format, a...)
}

func (o *Output) Fatal(a ...interface{}) {
	o.DebugCase()
	log.Fatalln(a...)
}

func (o *Output) Fatalf(format string, a ...interface{}) {
	o.DebugCase()
	log.Fatalf(format, a...)
}

func (o *Output) Debug(a ...interface{}) {
	o.DebugCase()
	log.Println(a...)
}

func (o *Output) Debugf(format string, a ...interface{}) {
	o.DebugCase()
	log.Printf(format, a...)
}

func (o *Output) DebugCase() {
	log.Printf("Case #%d, input: %v, output: %s\n", o.caseN, o.bip.Data, string(o.output.Bytes()))
}

func (o *Output) init(bip *bufferedInputProvider, caseN int) {
	o.bip = bip
	o.caseN = caseN
}

func (o *Output) flush() {
	if o.output.Len() <= 0 {
		o.Fatal("No output")
	}
	fmt.Fprintf(o.w, "Case #%d:", o.caseN)
	if !unicode.In(rune(o.output.Bytes()[0]), unicode.White_Space) {
		o.w.Write([]byte{' '})
	}
	o.w.Write(o.output.Bytes())
	if o.output.Bytes()[o.output.Len()-1] != '\n' {
		o.w.Write([]byte{'\n'})
	}
	o.output.Reset()
}

func (o *Output) assertOutput(fatal []bool, a ...interface{}) {
	output := o.Debug
	if len(fatal) > 0 && fatal[0] {
		output = o.Fatal
	}
	output(a...)
}

func (o *Output) assertOutputf(fatal []bool, format string, a ...interface{}) {
	outputf := o.Debugf
	if len(fatal) > 0 && fatal[0] {
		outputf = o.Fatalf
	}
	outputf(format, a...)
}

func (o *Output) AssertByteCount(a byte, count int, fatal ...bool) {
	for _, b := range o.output.Bytes() {
		if a == b {
			count--
		}
	}
	if count != 0 {
		o.assertOutputf(fatal, "AssertByteCount: byte: %q difference: %d", string(a), count*-1)
	}
}

func (o *Output) AssertCount(count int, fatal ...bool) {
	if o.output.Len() != count {
		o.assertOutput(fatal, "AssertCount: ", count)
	}
}
