package integer

import (
	"bytes"
	"strconv"
)

type Slice []int

func NewSlice(l int) Slice {
	s := make([]int, l)
	return s
}

func (s Slice) Fill(data ...int) {
	copy(s, data)
}

func (s Slice) String() string {
	buffer := &bytes.Buffer{}
	for i, d := range s {
		if i != 0 {
			buffer.WriteByte(' ')
		}
		buffer.WriteString(strconv.Itoa(d))
	}
	return buffer.String()
}
