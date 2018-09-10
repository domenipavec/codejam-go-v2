package io

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutput(t *testing.T) {
	b := &bytes.Buffer{}
	o := newOutput(b, false)
	o.init(nil, 3)

	o.Print("test")
	o.flush()
	assert.Equal(t, "Case #3: test\n", string(b.Bytes()))
	b.Reset()

	o.Print("\ntest")
	o.flush()
	assert.Equal(t, "Case #3:\ntest\n", string(b.Bytes()))
	b.Reset()

	o.Print("test\n")
	o.flush()
	assert.Equal(t, "Case #3: test\n", string(b.Bytes()))
	b.Reset()

	o.Print(" test")
	o.flush()
	assert.Equal(t, "Case #3: test\n", string(b.Bytes()))
	b.Reset()

	o.Println("test", 1)
	o.Println("test1")
	o.flush()
	assert.Equal(t, "Case #3: test 1\ntest1\n", string(b.Bytes()))
	b.Reset()
}
