package mod

import "strconv"

const PRIME = 1000000007

type Mod struct {
	value int
}

func New(v int) Mod {
	m := Mod{
		value: v,
	}
	m.fix()
	return m
}

func (m *Mod) fix() {
	if m.value >= PRIME || m.value < 0 {
		m.value = m.value % PRIME
		if m.value < 0 {
			m.value += PRIME
		}
	}
}

func (m *Mod) Inc() {
	m.value++
	if m.value >= PRIME {
		m.value -= PRIME
	}
}

func (m *Mod) Dec() {
	m.value--
	if m.value < 0 {
		m.value += PRIME
	}
}

func (m *Mod) Add(values ...Mod) {
	for _, value := range values {
		m.value += value.value
		if m.value >= PRIME {
			m.value -= PRIME
		}
	}
}

func (m *Mod) Sub(values ...Mod) {
	for _, value := range values {
		m.value -= value.value
		if m.value < 0 {
			m.value += PRIME
		}
	}
}

func (m *Mod) Mul(values ...Mod) {
	for _, value := range values {
		m.value *= value.value
		m.fix()
	}
}

func (m *Mod) Div(value Mod) {
	value.Exp(PRIME - 2)
	m.Mul(value)
}

func (m *Mod) Exp(value int) {
	if value < 0 {
		m.Exp(-1 * value)
		m.Exp(PRIME - 2)
	} else if value == 0 {
		m.value = 1
	} else if value == 1 {

	} else if value%2 == 0 {
		m.Mul(*m)
		m.Exp(value / 2)
	} else {
		tmp := *m
		m.Mul(tmp)
		m.Exp((value - 1) / 2)
		m.Mul(tmp)
	}
}

func (m Mod) Value() int {
	return m.value
}

func (m Mod) String() string {
	return strconv.Itoa(m.value)
}
