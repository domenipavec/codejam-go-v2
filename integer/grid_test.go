package integer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrid(t *testing.T) {
	g := NewGrid(2, 3)

	assert.Equal(t, 0, g[0][0])
	assert.Equal(t, 0, g[0][1])
	assert.Equal(t, 0, g[0][2])
	assert.Equal(t, 0, g[1][0])
	assert.Equal(t, 0, g[1][1])
	assert.Equal(t, 0, g[1][2])

	g.FillCol(1, 1, 2)

	assert.Equal(t, 0, g[0][0])
	assert.Equal(t, 1, g[0][1])
	assert.Equal(t, 0, g[0][2])
	assert.Equal(t, 0, g[1][0])
	assert.Equal(t, 2, g[1][1])
	assert.Equal(t, 0, g[1][2])

	g.FillRow(1, 4, 5, 6)

	assert.Equal(t, 0, g[0][0])
	assert.Equal(t, 1, g[0][1])
	assert.Equal(t, 0, g[0][2])
	assert.Equal(t, 4, g[1][0])
	assert.Equal(t, 5, g[1][1])
	assert.Equal(t, 6, g[1][2])

	assert.Equal(t, Slice([]int{1, 5}), g.GetCol(1))
	assert.Equal(t, Slice([]int{4, 5, 6}), g.GetRow(1))

	assert.Equal(t, "0 1 0\n4 5 6", g.String())
}
