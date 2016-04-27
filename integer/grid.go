package integer

import (
	"bytes"
	"strconv"
	"text/tabwriter"

	"github.com/matematik7/codejam-go/st"
)

type Grid [][]int

func NewGrid(y, x int) Grid {
	grid := make([][]int, y)
	for i := range grid {
		grid[i] = make([]int, x)
	}
	return grid
}

func (g Grid) FillRow(i int, data ...int) {
	copy(g[i], data)
}

func (g Grid) FillCol(i int, data ...int) {
	for j := range g {
		g[j][i] = data[j]
	}
}

func (g Grid) FillRowTuple(i int, tuple *st.Tuple) {
	g.FillRow(i, tuple.Ints...)
}

func (g Grid) FillColTuple(i int, tuple *st.Tuple) {
	g.FillCol(i, tuple.Ints...)
}

func (g Grid) GetRow(i int) Slice {
	s := NewSlice(len(g[i]))
	s.Fill(g[i]...)
	return s
}

func (g Grid) GetCol(i int) Slice {
	s := NewSlice(len(g))
	for j := range g {
		s[j] = g[j][i]
	}
	return s
}

func (g Grid) String() string {
	buffer := &bytes.Buffer{}
	for y := range g {
		if y != 0 {
			buffer.WriteByte('\n')
		}
		for x, d := range g[y] {
			if x != 0 {
				buffer.WriteByte(' ')
			}
			buffer.WriteString(strconv.Itoa(d))
		}
	}
	return buffer.String()
}

func (g Grid) GoString() string {
	buffer := &bytes.Buffer{}
	tw := tabwriter.NewWriter(buffer, 0, 0, 0, ' ', 0)
	for y := range g {
		if y != 0 {
			tw.Write([]byte{'\n'})
		}
		for x, d := range g[y] {
			if x != 0 {
				tw.Write([]byte{'\t', '|', '\t'})
			}
			tw.Write([]byte(strconv.Itoa(d)))
		}
	}
	tw.Flush()
	return buffer.String()
}
