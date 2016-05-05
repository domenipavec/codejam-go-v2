package stringmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringMap(t *testing.T) {
	sm := New()

	i1 := sm.Int("test1")
	i2 := sm.Int("test2")
	i3 := sm.Int("test3")

	assert.NotEqual(t, i1, i2)
	assert.NotEqual(t, i2, i3)

	assert.Equal(t, i1, sm.Int("test1"))
	assert.Equal(t, i2, sm.Int("test2"))
	assert.Equal(t, i3, sm.Int("test3"))

	assert.Equal(t, "test1", sm.Get(i1))
	assert.Equal(t, "test2", sm.Get(i2))
	assert.Equal(t, "test3", sm.Get(i3))
}
