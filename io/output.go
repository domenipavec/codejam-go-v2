package io

import (
	"bytes"
	"fmt"
	"io"
)

type Output struct {
	w io.Writer

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

func (o *Output) flush(caseN int) {
	fmt.Fprintf(o.w, "Case #%d: ", caseN)
	o.w.Write(o.output.Bytes())
	if o.output.Bytes()[o.output.Len()-1] != '\n' {
		o.w.Write([]byte{'\n'})
	}
	o.output.Reset()
}
