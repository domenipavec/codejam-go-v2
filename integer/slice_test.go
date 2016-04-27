package integer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	s := NewSlice(3)

	assert.Equal(t, 0, s[0])
	assert.Equal(t, 0, s[1])
	assert.Equal(t, 0, s[2])

	s.Fill(1, 2, 3)

	assert.Equal(t, 1, s[0])
	assert.Equal(t, 2, s[1])
	assert.Equal(t, 3, s[2])

	assert.Equal(t, "1 2 3", s.String())
}
