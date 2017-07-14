package mod

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert.Equal(t, 1000, New(1000).Value())
	assert.Equal(t, 200, New(200).Value())
	assert.Equal(t, 0, New(1000000007).Value())
	assert.Equal(t, 1, New(1000000008).Value())
	assert.Equal(t, 1000000006, New(-1).Value())
}

func TestInc(t *testing.T) {
	m := New(3)
	m.Inc()
	assert.Equal(t, 4, m.Value())
	m = New(1000000006)
	m.Inc()
	assert.Equal(t, 0, m.Value())
}

func TestDec(t *testing.T) {
	m := New(3)
	m.Dec()
	assert.Equal(t, 2, m.Value())
	m = New(0)
	m.Dec()
	assert.Equal(t, 1000000006, m.Value())
}

func TestAdd(t *testing.T) {
	m := New(1000000005)
	m.Add(New(5))
	assert.Equal(t, 3, m.Value())

	m = New(0)
	a := New(1000000000)
	b := New(1000000000)
	c := New(1000000000)
	m.Add(a, b, c)
	assert.Equal(t, 999999986, m.Value())
	assert.Equal(t, 1000000000, a.Value())
	assert.Equal(t, 1000000000, b.Value())
	assert.Equal(t, 1000000000, c.Value())

	m = New(3)
	m.Add(New(2000000017))
	assert.Equal(t, 6, m.Value())
}

func TestSub(t *testing.T) {
	m := New(3)
	m.Sub(New(5))
	assert.Equal(t, 1000000005, m.Value())

	m = New(0)
	a := New(1000000000)
	b := New(1000000000)
	c := New(1000000000)
	m.Sub(a, b, c)
	assert.Equal(t, 21, m.Value())
	assert.Equal(t, 1000000000, a.Value())
	assert.Equal(t, 1000000000, b.Value())
	assert.Equal(t, 1000000000, c.Value())

	m = New(3)
	m.Sub(New(2000000017))
	assert.Equal(t, 0, m.Value())
}

func TestMul(t *testing.T) {
	m := New(100000)
	m.Mul(New(100000))
	assert.Equal(t, 999999937, m.Value())

	m = New(100000)
	m.Mul(New(-1))
	assert.Equal(t, 999900007, m.Value())

	m = New(100000)
	m.Mul(New(9223372036854775807))
	assert.Equal(t, 200096181, m.Value())

	m = New(100000)
	m.Mul(New(-9223372036854775807))
	assert.Equal(t, 799903826, m.Value())
}

func TestExp(t *testing.T) {
	m := New(100000)
	m.Exp(0)
	assert.Equal(t, 1, m.Value())

	m = New(100000)
	m.Exp(1)
	assert.Equal(t, 100000, m.Value())

	m = New(100000)
	m.Exp(2)
	assert.Equal(t, 999999937, m.Value())

	m = New(100000)
	m.Exp(3)
	assert.Equal(t, 993000007, m.Value())

	m = New(100000)
	m.Exp(4)
	assert.Equal(t, 4900, m.Value())

	m = New(100000)
	m.Exp(9223372036854775807)
	assert.Equal(t, 841443036, m.Value())

	n := New(4)
	m = New(2)
	m.Exp(-1)
	n.Mul(m)
	assert.Equal(t, 2, n.Value())
	assert.Equal(t, 500000004, m.Value())

	n = New(4)
	m = New(2)
	m.Exp(-2)
	n.Mul(m)
	assert.Equal(t, 1, n.Value())
	assert.Equal(t, 250000002, m.Value())
}

func TestDiv(t *testing.T) {
	m := New(6)
	m.Div(New(3))
	assert.Equal(t, 2, m.Value())

	m = New(73741817)
	m.Div(New(2))
	assert.Equal(t, 536870912, m.Value())
}

func TestOutput(t *testing.T) {
	m := New(1000000119)

	buf := &bytes.Buffer{}
	fmt.Fprint(buf, m)
	assert.Equal(t, "112", buf.String())

	buf.Reset()
	fmt.Fprintln(buf, m)
	assert.Equal(t, "112\n", buf.String())

	buf.Reset()
	fmt.Fprintf(buf, "%v", m)
	assert.Equal(t, "112", buf.String())
}
