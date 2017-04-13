package grid

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "ValueType=BUILTINS"

// ValueType genny generic type for grid
type ValueType generic.Type

// GridValueType type
type GridValueType [][]ValueType

// NewValueType creates grid n by m
func NewValueType(n, m int) GridValueType {
	grid := make([][]ValueType, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]ValueType, m)
	}
	return grid
}

// Copy makes a new independent copy of grid
func (grid GridValueType) Copy() GridValueType {
	newGrid := NewValueType(len(grid), len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			newGrid[i][j] = grid[i][j]
		}
	}
	return newGrid
}

// Set sets all element to c
func (grid GridValueType) Set(c ValueType) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			grid[i][j] = c
		}
	}
}

// SpiralIterator returns []Coordinate in spiral order
func (grid GridValueType) SpiralIterator() []Coordinate {
	data := make([]Coordinate, len(grid)*len(grid[0]))
	return grid.spiralTopRight(data, 0, 0, len(grid)-1, len(grid[0])-1)
}

func (grid GridValueType) spiralTopRight(data []Coordinate, x1, y1, x2, y2 int) []Coordinate {
	for j := x1; j <= x2; j++ {
		data = append(data, Coordinate{I: y1, J: j})
	}

	for i := y1 + 1; i <= y2; i++ {
		data = append(data, Coordinate{I: i, J: x2})
	}

	if x2-x1 > 0 {
		data = grid.spiralBottomLeft(data, x1, y1+1, x2-1, y2)
	}

	return data
}

func (grid GridValueType) spiralBottomLeft(data []Coordinate, x1, y1, x2, y2 int) []Coordinate {
	for j := x2; j >= x1; j-- {
		data = append(data, Coordinate{I: y2, J: j})
	}

	for i := y2 - 1; i >= y1; i-- {
		data = append(data, Coordinate{I: i, J: x1})
	}

	if x2-x1 > 0 {
		data = grid.spiralTopRight(data, x1+1, y1, x2, y2-1)
	}

	return data
}
