package io

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareOutput(t *testing.T) {
	input := `Case #3: test
Case #123: test1
test2`

	co := NewCompareOutput(strings.NewReader(input))

	assert.True(t, co.HasOutput(3))
	assert.True(t, co.HasOutput(123))
	assert.False(t, co.HasOutput(1))
	assert.False(t, co.HasOutput(122))

	assert.Equal(t, []byte(" test\n"), co.GetOutput(3))
	assert.Equal(t, []byte(" test1\ntest2"), co.GetOutput(123))
}
