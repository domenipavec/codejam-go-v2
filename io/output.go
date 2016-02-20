package io

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"
	"unicode"
)

type Output struct {
	w io.Writer

	caseN int
	input *Input

	output *bytes.Buffer

	periodicPrint       chan struct{}
	previousPeriodicInt int
}

func newOutput(w io.Writer) *Output {
	return &Output{
		w:             w,
		output:        &bytes.Buffer{},
		periodicPrint: make(chan struct{}, 10),
	}
}

func (o *Output) resetPeriodic() {
	for {
		select {
		case <-o.periodicPrint:
		default:
			return
		}
	}
}

func (o *Output) triggerPeriodic() {
	o.resetPeriodic()
	o.periodicPrint <- struct{}{}
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
	log.Printf("Case #%d, input: %v, output: %q\n", o.caseN, o.input.currentCase(), string(o.output.Bytes()))
}

func (o *Output) Periodic(a ...interface{}) {
	select {
	case <-o.periodicPrint:
		log.Println(a...)
	default:
	}
}

func (o *Output) PeriodicInt(a int) {
	select {
	case <-o.periodicPrint:
		log.Println(a, "Rate =", a-o.previousPeriodicInt)
		o.previousPeriodicInt = a
	default:
	}
}

func (o *Output) Periodicf(format string, a ...interface{}) {
	select {
	case <-o.periodicPrint:
		log.Println(a...)
	default:
	}
}

func (o *Output) init(input *Input, caseN int) {
	o.input = input
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

func (o *Output) AssertEqual(data string, fatal ...bool) {
	if strings.TrimSpace(string(o.output.Bytes())) != strings.TrimSpace(data) {
		o.assertOutputf(fatal, "Output should be: %q", data)
	}
}
